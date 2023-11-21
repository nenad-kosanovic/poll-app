# Polling Application


## Requirements

The Backend service must be written using Golang. Postgres SDK should be used to integrate your backend service with your database. Poll CRUD and List APIs should be defined and implemented inside this service. An API to vote on a poll should also be implemented.
Later if time permits JWT auth can also be implemented.


- Storing poll and votes related data inside a DB
- Poll CRUD via API and UI
- Option Addition/Deletion (optional)
- Get and List Polls
- Vote on a Poll (via UI and API)
- Implement JWT based simple/basic authentication (Optional)
- Use postgres as your database. Come up with a database design to model Polls and Votes. Create those table(s) in your Database

## Database

Database connection string is set in ./app.env file. For local dev env, there od ./docker-compose.yml available in order to run db.
Before starting application, ./migrate/migrate.go should run, in order to init tables.


## Missing functionality

- Poll update
- JWT and user session
- Number of votes on poll is not limited and there is no option to remove vote.


