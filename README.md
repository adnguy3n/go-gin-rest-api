# go-gin-rest-api

Just a basic RESTful CRUD API written as a learning excercise with Golang and Docker.

## Running

To run, after pulling the repository, run Docker compose build followed by Docker compose up.

## Usage

Postman, or a similar tool, is recommended as there is currently no frontend.

![POPSTMAN Get Home](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/34a3681f-2fc3-4bbc-9143-1fc843d573ad)

### User Management

A number of Requests require JWT validation. To start, it is recommended to register a user. This can be done by adding users to the request URL and sending a POST request with a JSON body consisting of:

```
{
    "Name": "name",
    "Username": "username",
    "Email": "example@example.com",
    "Password": "password"
}
```

Fill it out as you wish, and I highly recommend that you write the information down somewhere as there is no easy way to look up the information and the password is hashed before being stored.

![POPSTMAN Register User](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/acade117-481b-4a29-8c2c-110e07b682e9)

Once a user is registered, send a POST request to http://localhost:8080/users/token with a body consisting of the email and password.

![POPSTMAN Generate Token](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/2dd0d2e9-b949-4c22-9008-1e1087c233ae)

You will receive a token. Copy the token, specifically the section highlighted in the above image.

Then head to the Headers tab in Postman. Add a header with the key Authorization and put in the token for the value.

![POPSTMAN Authorization](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/f324e5c1-f288-4649-a522-0a4fac72e9fc)

The token will last for one hour, afterwhich you will have to generate a new token and update the value of the authorization header.

#### Updating/Deleting User information

Users can be updated by sending a PATCH request with the username added to the Request URL and a JSON body akin to the one used to register the user. Users can also be deleted by sending a DELETE request with the username added to the Request URL; this is a hard delete, not a soft delete. All of these requests will require an active authenthication token.

### D&D Characters
Adding characters to the request URL will let you use the character routes to manage some simple structs of D&D characters. These are stored in a local database in the directories located at /database/characters.db.

![POPSTMAN Get Characters](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/d068ceaf-a94d-4789-8bb6-bba055ceeb20)

Each character has a unique ID that is automatically generated when they are added to the database. You can get a specific character by adding the ID to the Request URL; for example, http://localhost:8080/characters/1 would get you the character with the ID of 1.

![POSTMAN Get Character 1](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/ae770b96-1a2c-40cc-b5f4-f8811d4705a6)

Any request beyond getting characters requires an active authenthication token.

#### Creating a character

Making a POST Request a JSON body will allow you to create a character that is stored in a database within the API's directories.

```
{
    "Name": "Elminster",
    "Race": "Human",
    "Class": "Wizard",
    "Level": 20
}
```

The example above would create a 20th level Human Wizard named Elminster. Do note that Level cannot exceed 20 and an error noting such will be returned if attempted.

![POSTMAN create character](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/7727d496-fd84-4524-9021-b9706bb7997e)

The character's Unique ID is automatically generated and should not be provided in the JSON body.

#### Updating a character

The PATCH and PUT requests can be used to update a character. They require a JSON Body and for the character ID to be added to the request URL. The entire JSON does not need to be provided for PATCH requests, but does, including the unique ID, for PUT requests. Also, PATCH requests cannot change the ID, only PUT requests can.

![POSTMAN update character](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/7c3d182e-1dc8-4fd0-87a2-b3c776a6799a)

PATCH and PUT requests require an active authenthication token.

#### Deleting a character
Characters can be deleted; to do so, add the character ID to the request URL and send a DELETE request.

![POSTMAN delete character](https://github.com/adnguy3n/go-gin-rest-api/assets/32573771/ba2adfe4-e794-4eda-925c-1a2f2e339e6b)
