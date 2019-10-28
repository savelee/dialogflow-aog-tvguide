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

SA_EMAIL=$(gcloud iam service-accounts list \
  --filter="displayName:$SERVICE_ACCOUNT_NAME" \
  --format='value(email)')

bold "Removing Storage"
gsutil rm -r gs://$GCLOUD_STORAGE_BUCKET_NAME/

bold "Deleting Cloud Functions"
gcloud functions delete tvguide \
--region=$REGION

bold "Deleting Cloud Run"
gcloud beta run services delete $PROJECT_ID \
--region=$REGION --platform managed

bold "Removing All Dialogflow Agents..."
ACCESS_TOKEN="$(gcloud auth application-default print-access-token)"
curl -H "Content-Type: application/json; charset=utf-8"  \
-H "Authorization: Bearer $ACCESS_TOKEN" \
-X "DELETE" "https://dialogflow.googleapis.com/v2/projects/$PROJECT_ID/agent"

bold "Deleting service account $SERVICE_ACCOUNT_NAME..."
gcloud iam service-accounts delete $SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com --quiet

bold "Uninstallation complete!"