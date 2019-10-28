# Auto Install:

This script will install the following resources:

* Dialogflow Agent
* Google Cloud Storage Bucket
* Cloud Function
* Cloud Run
* Service Account

1. Make sure `$PROJECT_ID` is set: `gcloud config set project $PROJECT_ID`

2. To start installation: `. setup.sh`

## Enable the fulfillment webhook in Dialogflow

1. In the Dialogflow Console, select **Fulfillment** from the menu.
2.  **Enable** Webhooks
3.  Enter the URL of your Cloud Function, you can find the
    url, in the Google Cloud Functions overview. It's likely
    https://europe-west1-[YOUR-PROJECT-ID].cloudfunctions.net/tvguide
4. Click **Save**

## Configure Actions on Google

1. Navigate to http://console.dialogflow.com and login with
google account. (the same, used for GCP).
2. Click on **See how it works in Google Assistant**.  in the right hand pane. 
3. It will open http://console.actions.google.com
4. Try to open your action in the simulator, by clicking on the project name.
3. Select **Test** in the menu bar 
4. Make sure, the simulator is set to **English** and Click on **Talk to my test-app**
The action will greet you, with the basic Dialogflow default intent.

## Test your action.

*Ask: "What's on MTV at noon?"*

# Remove all

This script will remove the following resources:

* Dialogflow Agent
* Google Cloud Storage Bucket
* Cloud Function
* Cloud Run
* Service Account

To uninstall run: `. teardown.sh`