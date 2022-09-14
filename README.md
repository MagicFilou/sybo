# SYBO SBE Assignment

## Setup

The program is written in Go.
Before you get the program running there are a few steps to be done:

1. Run the migration to create the necessary table in your postgres db. The SQL is in `migrations/000001_users.up`
2. Set your environment variables, here are the following ones required:

- ENV should have value _local_ or _prod_ according to the moe you are running
- PORT should be the port in user for example: _1930_
- DB_NAME name of the DB
- DB_HOST host of the db
- DB_PORT port of the db
- DB_USER user for that db
- DB_PASSWORD password for the user above

## Run the code

### Locally

1. Setup the envs for example with `source .env`
2. Run `go run main.go`

### Docker

1. Setup the envs, with a configmap or `ENV FOO=bar` in the dockerfile
2. Build the image with the command `docker build -t sybo .`
3. Run the image with the command `docker run -d -p 1930:1930 sybo`

## Endpoints documentation

### Shared info

All endpoints are JSON based.

Only the method specified here will give a result otherwise the Status code associated is _404_

If everything runs as expected the response will be a parsable JSON (see the documentation for the formats). Status code associated is _200_

If no results are found, there are no response body and the Status code associated is _204_

If there is an issue with the data provided e.g. a faulty ID or faulty request body the error will be given with the following body format:
`{"error": "error as string"}` the Status code associated is 400

Any other issue with result of a status code 500 and the error (following the same format as above) returned.

The project has 6 enpoints, corresponding to the requirements given by my "friend" aka the assignment :

- GET `/user`
  This is the only endpoint that handles parameters for now. You can put any param you want but _id_ and _name_ will be the most used ones. For now this is a slight improve from the specification and it uses a "like" clause to fetch and filter the data so it will only work with string params. For now this is case sensitive but it could be improve for future uses.

  If you would like to switch to a pure match (using "=") it is easily changable in the code

  Request example:

  - calling `/user` will return all the users
  - calling `/user?name=Jo` will return all the users with a name containing "Jo"

  Response example:
  ` { "users": [ { "id": "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1", "name": "John" }, { "id": "f9a9af78-6681-4d7d-8ae7-fc41e7a24d08", "name": "Bob" }, { "id": "2d18862b-b9c3-40f5-803e-5e100a520249", "name": "Alice" } ]`

- POST `/user`
  Creates a new user with the name provided.

  Mandatory body parameters: **name**

  Request example:
  `{ "name": "John" }`

  Reponse example:
  `{ "id": "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1", "name": "John" }`

- GET `/user/<userid>/state`
  Get the state for the user provided.

  Mandatory URL parameters: **userid**

  Reponse example:
  `{ "gamesPlayed": 42, "score": 358 }`

- PUT `/user/<userid>/state`
  Update the state for the user provided. Note that score is not mandatory because you could have a played not scoring anything but incrementing _gamesPlayed_.

  Mandatory URL parameters: **userid**
  Mandatory body parameters: **gamesPlayed**

  Request example:
  `{ "gamesPlayed": 42, "score": 358 }`

  - GET `/user/<userid>/friends`
    Get the friends for a user, containing their _id_, _name_ and _highscore_.

  Mandatory URL parameters: **userid**

  Reponse example:
  `{ "friends": [ { "id": "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1", "name": "John", "highscore": 322 }, { "id": "f9a9af78-6681-4d7d-8ae7-fc41e7a24d08", "name": "Bob", "highscore": 21 }, { "id": "2d18862b-b9c3-40f5-803e-5e100a520249", "name": "Alice", "highscore": 99332 } ] }`

- PUT `/user/<userid>/friends`
  Update the list of friends for a user.

  Mandatory URL parameters: **userid**
  Mandatory body parameters: **friends**

  Request example:
  `{ "friends": [ "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1", "f9a9af78-6681-4d7d-8ae7-fc41e7a24d08", "2d18862b-b9c3-40f5-803e-5e100a520249" ] }`

## Comments and Room for improvements

I have made the repo so there is room to be quite extended. Routes, Handlers and Models can easily be extended with either shared items or dedicated data models. The same for clients and utils which are also quite flexible.

I have also started on a migration pattern but it could be called on the startup of the service to apply migrations or manualy for any changes.

Also i have added a _mw_ folder to put middlewares there. Currently i have not added anything (except for a dummy auth mw). But this would be a good place to check tokens or use other kind of authentication. This could also be the place to make a CORS middleware to handle any issues related to CORS.

I used gorm to handle the calls to the database, it is a powerful tool and it requires less maintenance to the queries if the data layer changes. Also several drivers are available for other sql techs. I could also have used direct call to the db and queries, it has its advantages for small projects or significant advance sql requirements.

Regarding the _Friends_ list of the user I, decided to go for a comma separated list. It works, however it has advantages and inconvenient. Pros: easy to replace all, quick and simple solution. No need to check for the current list to delete the non existing ones. No need to add another table in the DB.
Cons: does **NOT** scale, cannot check as a foreign key.
In the long term I would recommend to do the friends as a separated table but it was so much easier to start with a comma separated list
I think a separate table could be apply to the scores for future uses

Still related to the friends, it would be nice if my friend harmonize the json and use _highscore_ or _score_ both places (friends and user state). it would be more consistent.

Maybe an update user endpoint (generic) could be an easier way to use rather than several smaller update endpoints.

I would recommend to use the method PATCH rather than PUT. PUT should be used to replace the entire object and PATCH only to do partial updates.

I am probably missing a few const here and there for the errors or specific strings

In the case of the Get Users, i have started to work on a filtering solution in the query. Right now it will only use the first value for each param but i am sure it could be improved in the future. I could also reflect the value given and use either _=_ or _like_ in the `where` clause depending of it. It would be fun to also be able to filter the users by ranks in the high score or number of games played.
