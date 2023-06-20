# Mock API Server

It's a mock API server

## Start Server
### Install docker

Install docker on your PC. You can download and install Docker Destop here https://docs.docker.com/docker-for-windows/install/


### Run below commands
```sh
dockder-compose build #only execute this line when you pull new source code
docker-compose up -d
```

## API Authentication

EXCEPT `LOGIN` AND `SIGNUP` APIs, YOU NEED TO SET HEADER `Authorization: your_token` TO ACCESS API.

## API Endpoints

### Signup API

- Endpoint: `http://localhost:5151/signup`
- Method: POST
- Request Header Content-Type: JSON
- Request JSON Body:
    - `email`: email string, max 255 characters
    - `fullname`: string, max 255 characters
    - `password`: string, max 255 characters

Request JSON Body Example
```json
{
    "email": "youremail@domain.com",
    "fullname": "Your Name",
    "password": "123456"
}
```

JSON Response Example
```json
{
    "data": {
        "id": "6422bc2c-70a2-45d5-90fe-4e60173e2aee",
        "email": "youremail@domain.com",
        "fullname": "Your Name",
        "password": "123456"
    }
}
```

### Login API

#### Login using POST method

- Endpoint: `http://localhost:5151/login`
- Method: POST
- Request JSON Body:
    - `email`: email string, max 255 characters
    - `password`: string, max 255 characters

Request JSON Body Example
```json
{
    "email": "youremail@domain.com",
    "password": "123456"
}
```

Response:
```json
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAwNzQwOTAsInVzZXJfaWQiOiIzZjExM2YzZC0zZGNhLTRlZDMtYTU2ZC00NjBjYTk2NTg3YjgifQ.y6chJiJjqG_NTXDssJzZRk_xnE7B_y9dKF6kG7TW5Z0"
}
```

#### Login using GET method

- Endpoint: `http://localhost:5151/login?email=youremail@domain.com&password=123456`
- Method: GET
- Reponse:
```json
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAwNzQwOTAsInVzZXJfaWQiOiIzZjExM2YzZC0zZGNhLTRlZDMtYTU2ZC00NjBjYTk2NTg3YjgifQ.y6chJiJjqG_NTXDssJzZRk_xnE7B_y9dKF6kG7TW5Z0"
}
```

### User Profile API

- Endpoint: `http://localhost:5151/profile`
- Method: GET
- Reponse: 
```json
{
    "data": {
        "id": "924a5ac3-2f67-4f27-92f0-2dafe129c7e2",
        "email": "youremail@gmail.com",
        "fullname": "Your Name"
    }
}
```

### Change Password API

- Endpoint: `http://localhost:5151/change_password`
- Method: POST
- Request JSON Body:
    - `new_password`: string
- Reponse: 
```json
{
    "data": true
}
```

### Logout API

- Endpoint: `http://localhost:5151/logout`
- Method: POST
- Request Header Content-Type: JSON
- Request JSON Body: No

Success Response
```json
{
    "data": true
}
```

### Create Todo API

- Endpoint: `http://localhost:5151/todos`
- Method: POST
- Request Header Content-Type: JSON
- Request JSON Body:
    - `title`: string, max 255 characters
    - `description`: text

Request JSON Body Example:
```json
{
    "title": "Eu vix putent ceteros",
    "description": "Eu vix putent ceteros. Et usu tempor perpetua, sea ludus labitur eu. Te has etiam tempor expetenda. Ex possit detracto nominati ius, vix te dicat dicam habemus, ei omnes primis omnesque ius."
}
```
Response example:
```json
"data": {
    "id": "0b8f0354-1852-499e-8f0d-ecff3c693aee",
    "user_id": "3f113f3d-3dca-4ed3-a56d-460ca96587b8",
    "title": "Eu vix putent ceteros",
    "description": "Eu vix putent ceteros. Et usu tempor perpetua, sea ludus labitur eu. Te has etiam tempor expetenda. Ex possit detracto nominati ius, vix te dicat dicam habemus, ei omnes primis omnesque ius.",
}
```

### List Todos API

- Endpoint: `http://localhost:5151/todos`
- Method: GET

Response example:

```json
{
    "data": [
        {
            "id": "0b8f0354-1852-499e-8f0d-ecff3c693aee",
            "user_id": "3f113f3d-3dca-4ed3-a56d-460ca96587b8",
            "title": "Eu vix putent ceteros",
            "description": "Eu vix putent ceteros. Et usu tempor perpetua, sea ludus labitur eu. Te has etiam tempor expetenda.",
            "created_at": ""
        },
        {
            "id": "a1405d55-7306-4f2a-92f1-c38a99b70c31",
            "user_id": "3f113f3d-3dca-4ed3-a56d-460ca96587b8",
            "title": "Lorem ipsum dolor sit amet",
            "description": "Lorem ipsum dolor sit amet, pri eu dico illum causae, ex qui legere tacimates.",
        }
    ]
}
```

## Request Sequence Example

![](./assets/mock_server.png)

