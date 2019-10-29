#!/bin/bash

gcloud auth login

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
GCLOUD_STORAGE_BUCKET_NAME=lee-tvguide

bold "Starting the setup process in project $PROJECT_ID..."
bold "Enable APIs..."
gcloud services enable \
  dialogflow.googleapis.com \
  cloudfunctions.googleapis.com \
  iam.googleapis.com \
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
gcloud iam service-accounts keys create master.json \
  --iam-account $SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com
GOOGLE_APPLICATION_CREDENTIALS=master.json
ACCESS_TOKEN="$(gcloud auth application-default print-access-token)"

bold "Creating Storage bucket..."
gsutil mb gs://$GCLOUD_STORAGE_BUCKET_NAME/

bold "Create a Dialogflow Agent..."
echo $ACCESS_TOKEN

bold "Deploy Webserver To Cloud Run..."
docker build -t gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide cloudrun/go-tvguide-json-api
docker push gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide
gcloud beta run deploy --image gcr.io/${GOOGLE_CLOUD_PROJECT}/tvguide \
--platform managed --region $REGION --allow-unauthenticated

bold "Prepare Cloud Function to connect to the right URL"

BACKEND_URL="$(curl GET -H "Authorization: Bearer $ACCESS_TOKEN" https://europe-west1-run.googleapis.com/apis/serving.knative.dev/v1/namespaces/${GOOGLE_CLOUD_PROJECT}/services \
| jq -r '.items[0].status.url')"

echo $BACKEND_URL

while IFS= read -r line
do
    case "$line" in
       *BACKEND_URL*) echo "${line/BACKEND_URL/$BACKEND_URL}" >> cloudfunction/tvguide/index.js ;;
       *) echo "$line" >> cloudfunction/tvguide/index.js ;;
    esac
done < cloudfunction/tvguide/index-old.js
echo "}" >> cloudfunction/tvguide/index.js

bold "Deploy Cloud Function..."
gcloud functions deploy tvguide --region=$REGION \
--memory=512MB \
--runtime=nodejs10 \
--source=cloudfunction/tvguide \
--stage-bucket=$GCLOUD_STORAGE_BUCKET_NAME \
--timeout=60s \
--trigger-http
--entry-point=tvguide

bold "Prepare Dialogflow to connect to the right URL"

CF_URL="https://${REGION}-${GOOGLE_CLOUD_PROJECT}.cloudfunctions.net/tvguide"
echo $CF_URL

while IFS= read -r line
do
    case "$line" in
      *CF_URL*) echo "${line/CF_URL/$CF_URL}" >> dialogflow/agent/agent.json ;;
      *) echo "$line" >> dialogflow/agent/agent.json ;;
    esac
done < dialogflow/agent/agent-old.json
echo "}" >> dialogflow/agent/agent.json

bold "Zipping Intents..."
zip -r dialogflow/agent/agent.zip dialogflow/agent
bold "Uploading Intents to $GCLOUD_STORAGE_BUCKET_NAME..."
gsutil cp dialogflow/agent/agent.zip gs://$GCLOUD_STORAGE_BUCKET_NAME/

gcloud auth activate-service-account --key-file master.json
TOKEN="$(gcloud auth print-access-token)"

JSONPROD="{\"defaultLanguageCode\":\"en\",\"displayName\":\"tvguide\",\"parent\":\"projects/$PROJECT_ID\",\"timeZone\":\"Africa/Casablanca\"}"
curl -H "Content-Type: application/json; charset=utf-8"  \
-H "Authorization: Bearer $TOKEN" \
-d $JSONPROD "https://dialogflow.googleapis.com/v2/projects/$PROJECT_ID/agent"

IMPORTFILES="{\"agentUri\":\"gs://$GCLOUD_STORAGE_BUCKET_NAME/agent.zip\"}"
bold "Import Intents to Prod"
curl -X POST \
-H "Authorization: Bearer $TOKEN" \
-H "Content-Type: application/json; charset=utf-8" \
-d $IMPORTFILES \
https://dialogflow.googleapis.com/v2/projects/$PROJECT_ID/agent:import

bold "Setup & Deployment complete!"