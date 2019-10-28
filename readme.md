[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Hands-on:  Create a TV guide action for the Google Assistant with Google Cloud, Dialogflow and Actions on Google.

**By Lee Boonstra, Developer Advocate @ Google Cloud.**

[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.svg)](https://console.cloud.google.com/cloudshell/editor?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fsavelee%2Fdialogflow-tvguide&cloudshell_tutorial=tutorial.md)

This demo, showcases a tvguide application for the
Google Assistant. It makes use of the following resources:

* Dialogflow
* Actions on Google
* Cloud Functions
* Cloud Run

## Automatic Setup on Google Cloud Platform:

Guided one click installation from Google Cloud Shell. No client tooling required.

### Create a GCP Project

1. Navigate to: http://console.cloud.google.com
2. Once you are logged into GCP console, you can create a project. Click on the dropdown, next to the Google Cloud Platform logo. It will open a pop-up.
3. Click New Project
4. Give your project the name: <yourname>-tvguide.  Click Create.  
(NOTE: In case you already have a GCP account for your organization, you might see organization settings on this page as well.)

4. Make sure the project is enabled, by selecting it in the top bar.

### Enable Billing

Open the console left side menu and select Billing.
Set or Create a Billing Account.
The use of Dialogflow Enterprise and Cloud Functions, requires a billing account. Make sure it's setup, and linked to this project.
For more information on Billing Accounts see: https://cloud.google.com/billing/docs/how-to/modify-project


[![Open in Cloud Shell](http://gstatic.com/cloudssh/images/open-btn.svg)](https://console.cloud.google.com/cloudshell/editor?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fsavelee%2Fdialogflow-tvguide&cloudshell_tutorial=tutorial.md)

**Disclaimer: This example is made by Lee Boonstra. Written code can be used as a baseline, it's not meant for production usage.**

**Copyright 2018 Google LLC. This software is provided as-is, without warranty or representation for any use or purpose. Your use of it is subject to your agreements with Google.** 


