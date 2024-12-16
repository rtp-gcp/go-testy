
# this code uses the env variable GOOGLE_APPLICATION_CREDENTIALS

from google.cloud import firestore

db=firestore.Client(project="shawdavis")


# add data
doc_ref = db.collection("users").document("alovelace")
doc_ref.set({"first": "Ada", "last": "Lovelace", "born": 1815})




