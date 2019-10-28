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
ACCESS_TOKEN="$(gcloud auth application-default print-access-token)"

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

bold "Agent Setup Complete!"