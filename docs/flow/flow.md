### User auth flow.

##### Expected input
```JSON
{
  "login" : "",
  "password" : ""
}
```

- User passes login and password hash.
- Server checks if login exists.
  - Server verifies password.
    - Server generates JWT token.
    - Returns generated token.

##### **Token** has to contain such claims
1. user id
2. user role
3. user name
4. user role id
    
##### Expected output
```JSON
{
  "id" : "",
  "name" : "",
  "role" : "",
  "role_id" : "",
  "created_at" : "",
  "token" : ""
}
```

### User get requests and responders flow.

- Validate user JWT token. (Roles = **Dispatcher**, **Admin**)
  - Make calls to get all the active responders.
  - Make call to get all the requests for the last 7 days.
  - Map requests to its responders.
  - Return result.

#### Expected response object
``` JSON
[
  { // Responder
    "id" : "",
    "name" : "",
    "role" : "",
    "role_id" : "",
    "created_at" : "",
    "requests" : [
      {
        "id" : "",
        "title" : "",
        "description" : "",
        "created_at" : "",
        "updated_at" : "",
        "closed_at" : "",
        "dispatcher_id" : "",
        "responder_id" : "",
        "request_type_id" : "",
        "request_type" : ""
      }
    ]
  }
]
```

### User get request types

- Validate user JWT token.
  - Get all the **relevant** request types.

### Create new request

##### Expected input
```JSON
{
  "title" : "",
  "description" : "",
  "dispatcher_id" : "",
  "responder_id" : "",
  "request_type_id" : "",
}
```
- Validate user JWT token. (Roles = **Dispatcher**)
  - Check if provided **dispatcher_id**, **responder_id** and **request_type_id** are relevant and active.
  - Create new **request**.
  - Send new **request** to all the clients.

###### Expected output to all the clients
``` JSON
{
  "id" : "",
  "title" : "",
  "created_at" : "",
  "updated_at" : "",
  "closed_at" : "",
  "description" : "",
  "dispatcher_id" : "",
  "responder_id" : "",
  "request_type_id" : "",
  "request_type" : ""
}
```
