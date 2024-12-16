# GCP Firestore Notes

Native mode Firestore

us-east1

```
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read, write: if false;
    }
  }
}
```



## Service Account

```
gcloud init
gcloud components update
gcloud iam service-accounts create rs-firestore-sa
```

At this point, I went to gun console and create a key for this service account.

When I tried to run the code, it said this

```
gcloud auth application-default login --impersonate-service-account SERVICE_ACCT_EMAIL
```

Format: 
```
gcloud projects add-iam-policy-binding PROJECT_ID --member="serviceAccount:SERVICE_ACCOUNT_NAME@PROJECT_ID.iam.gserviceaccount.com" --role=ROLE
```

This requires a role setting, thus created a custom role, visible in IAM/Admin/Roles

```
projects/shawdavis/roles/CustomRole4Firestore
```


Command:
```
gcloud projects add-iam-policy-binding shawdavis --member="serviceAccount:rs-firestore-sa@shawdavis.iam.gserviceaccount.com" --role="projects/shawdavis/roles/CustomRole4Firestore"

```

Then I needed to associate this with my email

Format:
```
gcloud aim service-accounts add-iam-policy-binding SERVICE_ACCOUNT_NAME@PROJECT_ID.iam.gserviceaccount.com --member="user:USER_EMAIL" --role=roles/iam.serviceAccountUser

```
Command:
```
gcloud iam service-accounts add-iam-policy-binding rs-firestore-sa@shawdavis.iam.gserviceaccount.com --member="user:davisjf@gmail.com" --role=roles/iam.serviceAccountUser

```
This still failed until I added a permission for the service account to have the cloud datastore user role.

I circled back and modified the custom role to have cloud datastore user role permissions, deleted the cloud datastore role at IAM level and it still worked.



