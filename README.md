This project demonstrates how the 2FA (2 Factor Authentication) is done using Google Authenticator.

# Requirement
Depends on Postgres to save the user details and shared secret. The shared secret is not encrypted for simplicity. In production environment, this value should be encrypted.
Start the DB on your local using the following command from the root dir:

>docker-compose up

# How it works
Start the app using the following command from the root dir:

>go run main.go

When the app is started, DB migration takes care of creating necessary tables and inserting a test user with id as **1**. 

### Signup for 2FA
Signup endpoint is called to sign-up for the 2FA. This will generate and save the shared secret into the database. 
The endpoint also returns a QR code as a png file back to the browser. The user can then can the QR code in the Google Authenticator App to add. From this point, the Google Authenticator will start showing a 6 digit code which changes every 30 seconds or so.
>http://localhost:8080/mfa/signup?userId=1

### Verify the code
Verify endpoint is used to verify the code. The endpoint will return HTTP 200 if the code is valid. Otherwise, HTTP 403 is returned.

>http://localhost:8080/mfa/verify?userId=1&code=685014
