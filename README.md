# Dataimpact

Welcome to the Dataimpact API documentation!

This API provides access to various resources that allow you to retrieve, create, update, and delete data in a database.

## Getting Started

Make sure you have docker installed on your machine, then clone the project and use :

```shell
$ docker compose up
```

Once the app is launched you can either use the api through postman or with the "requests" folder that contains some usecases.  

## API Endpoints

This API provides the following endpoints:

### Users

    GET http://localhost:8080/api/v1/users/         // retrive all users
    GET http://localhost:8080/api/v1/users/{id}     // retrive one user by id 
    POST http://localhost:8080/api/v1/users/{id}    // seed the database users from a dataset
    PATCH http://localhost:8080/api/v1/users/{id}   // partially update a user
    DELETE http://localhost:8080/api/v1/users/{id}  // delete a user and its data

### Session

    POST http://localhost:8080/api/v1/login     // provides a JWT token to authenticate 

Authetication request body:
```json
    {
        "id": "your_user_id",
        "password": "your_password"
    }
```
## Authentication

To access this API, you need to authenticate yourself with a valid id and password. You can do this by sending a POST request to the /login endpoint with your credentials.

The API will then generate a JSON Web Token (JWT) that you need to include in the Authorization header of all subsequent requests.

## Tests

To run the tests locally, you need to create a mongodb container :

```shell
$  make mongo
```

Then execute:

```shell
$  make test
```
