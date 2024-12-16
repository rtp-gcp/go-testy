This is just code to verify the service account and google cloud platform firestore microservice works.

# workflow to recreate venv and run test code

## setup python venv

### macOS

1. python3 -m venv venv
2. . ./venv/bin/activate
3. pip install --upgrade pip
4. pip install -r requirements.txt

## Setup env variable for google sdk

1. . ../../bin/setenv.sh

## Verify setup

If this code runs without error, ie. no error messages, then
the service account and gcp datastore work.


