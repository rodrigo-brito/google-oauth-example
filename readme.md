## Golang OAuth Example - Google

## Setup

- Obtain OAuth 2.0 credentials in [Google Cloud Console](https://console.cloud.google.com/)
- Save the credentials in project root: `client_secrets.json`
- CAUTION! Don't include this file in a public repository, save it safely. 

### OAuth APIs
This example uses calendar API, but it works with any oAuth User API. To set up another api, change the scope in `login.go` and `app.go` files.
  - Example of scopes for Google Calendar: https://developers.google.com/calendar/auth  

### Login for backend applications

You need to login once, it will save a file named `request.token` with the credentials, and a refresh token, save it safely and use as a cache for your app.

CAUTION! Don't include this file in a public repository. 

```
go run cmd/login/login.go
```

## Example of usage

Listing all users calendar

```
 go run cmd/app/app.go 
```
