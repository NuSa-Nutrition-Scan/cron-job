# Cron Job

This project implements a cron job using Google App Engine and Firestore. The cron job is scheduled to run every night at 00:00 GMT+7 (Asia/Jakarta) and resets some fields in user data stored in Firestore.


## Team:
Cloud Computing

## Functionality

The cron job performs the following tasks:

- Connects to Firestore using the Google Cloud Firestore Go library.
- Queries the Firestore collection to identify user documents that need to be reset.
- Updates specific fields in the identified user documents to their default/reset values.
- Logs the execution of the cron job.

## How To

1. Make sure you have the following prerequisites installed:
   - Go (version 1.20+)
   - Google Cloud SDK (`gcloud`) command-line tool

2. Clone this repository and navigate to the project directory.

3. Set up your Google Cloud project:
   - Create a new project in the [Google Cloud Console](https://console.cloud.google.com/).
   - Enable the Firestore API for your project.
   - Set up authentication and generate a service account key file for your project. Download the JSON key file and save it securely as serviceAccountKey.json in cloned project.

4. Replace configuration in `app.yaml` and `cron.yaml` (if you want to)

5. Update the code in `main.go` with the logic to reset the specific fields in user data according to your requirements (if you want to).

6. Deploy the application to Google App Engine by running the following command:
   ```bash
   gcloud app deploy

7. Wait until 00:00, and see the result

## LICENSE

This project is licensed under the [MIT License](https://github.com/NuSa-Nutrition-Scan/cron-job/blob/main/LICENSE).
