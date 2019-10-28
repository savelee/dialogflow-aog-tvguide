#!/bin/bash

bold() {
  echo ". $(tput bold)" "$*" "$(tput sgr0)";
}

err() {
  echo "$*" >&2;
}

bold "Set all vars..."
PROJECT_ID=$(gcloud info --format='value(config.project)')
SERVICE_ACCOUNT_NAME="tvguide-app"
REGION=europe-west1
GCLOUD_STORAGE_BUCKET_NAME=tvguidebucket

bold "Starting the setup process in project $PROJECT_ID..."
bold "Enable APIs..."
gcloud services enable \
  dialogflow.googleapis.com \
  logging.googleapis.com \
  monitoring.googleapis.com \
  sourcerepo.googleapis.com \
  run.googleapis.com \
  cloudbuild.googleapis.com

bold "Creating a service account $SERVICE_ACCOUNT_NAME..."

gcloud iam service-accounts create \
  $SERVICE_ACCOUNT_NAME \
  --display-name $SERVICE_ACCOUNT_NAME

SA_EMAIL=$(gcloud iam service-accounts list \
  --filter="displayName:$SERVICE_ACCOUNT_NAME" \
  --format='value(email)')
  
if [ -z "$SA_EMAIL" ]; then
  err "Service Account email is empty. Exiting."
fi

bold "Adding policy binding to $SERVICE_ACCOUNT_NAME email: $SA_EMAIL..."
loud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/clouddebugger.agent
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/dialogflow.admin
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/dialogflow.reader
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/errorreporting.admin
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/logging.logWriter
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/storage.objectCreator
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/storage.objectViewer
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:$SA_EMAIL \
  --role roles/iam.serviceAccountKeyAdmin

bold "Saving the key..."
gcloud iam service-accounts keys create ../master.json \
  --iam-account $SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com

GOOGLE_APPLICATION_CREDENTIALS=../master.json
ACCESS_TOKEN="$(gcloud auth application-default print-access-token)"

bold "Creating Storage bucket..."
gsutil mb gs://$GCLOUD_STORAGE_BUCKET_NAME/

bold "Zipping Intents..."
zip -r dialogflow/agent/agent.zip dialogflow/agent
bold "Uploading Intents to $GCLOUD_STORAGE_BUCKET_NAME..."
gsutil cp dialogflow/agent/agent.zip gs://$GCLOUD_STORAGE_BUCKET_NAME/

bold "Create a Dialogflow Agent..."
echo $ACCESS_TOKEN

JSONPROD="{\"defaultLanguageCode\":\"en\",\"displayName\":\"tvguide\",\"parent\":\"projects/$PROJECT_ID\",\"timeZone\":\"Europe/Madrid\"}"
curl -H "Content-Type: application/json; charset=utf-8"  \
-H "Authorization: Bearer $ACCESS_TOKEN" \
-d $JSONPROD "https://dialogflow.googleapis.com/v2/projects/$PROJECT_ID/agent"

IMPORTFILES="{\"agentUri\":\"gs://$GCLOUD_STORAGE_BUCKET_NAME/agent.zip\"}"
bold "Import Intents to Prod"
curl -X POST \
-H "Authorization: Bearer $ACCESS_TOKEN" \
-H "Content-Type: application/json; charset=utf-8" \
-d $IMPORTFILES \
https://dialogflow.googleapis.com/v2/projects/$PROJECT_ID/agent:import

bold "Deploy Webserver To Cloud Run..."
docker build -t gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide cloudrun/go-tvguide-json-api
docker push gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide
gcloud beta run deploy --image gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide \
--platform managed --region $REGION --allow-unauthenticated

BACKEND_URL=curl GET -H "Authorization: Bearer $ACCESS_TOKEN" https://europe-west1-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/${GOOGLE_CLOUD_PROJECT}/services \
| jq -r '.items[0].status.url'

bold "Eval the templates & deploy CF..."
envsubst < cloudfunction/tvguide/index.js | gcloud functions deploy tvguide --region=$REGION \
--memory=512MB \
--runtime=nodejs10 \
--source=cloudfunction/tvguide \
--stage-bucket=$GCLOUD_STORAGE_BUCKET_NAME \
--timeout=60s \
--entry-point=tvguide

bold "Setup & Deployment complete!"