### COMMANDS TO RUN THE APP USING DOCKER (LOCAL ENV)
```
- docker build . --build-arg db_pass_env=${DB_PASS_ENV} -t postgres
- docker run --rm --name pg-docker -p 5432:5432 postgres:latest
```

If you want to run the container using a network you'll need to run
the following command
```
- docker run --rm --name pg-docker -p 5432:5432 --network myNetwork postgres:latest
```

### POSTGRES RELATED

<!-- From a Postgres UI (like PSequel) -->
SELECT * FROM users;
CALL registerUser('Name', '123123123', 'secret', 'email@gmail.com');
CALL registerUser('Name2', '456456456', 'secret2', 'email2@gmail.com');
SELECT * FROM users WHERE password = crypt('secret', password);

<!-- Access to Postgres running container using the terminal -->
docker exec -it pg-docker psql -h localhost -p 5432 -U pratik pratikdb

<!-- Create user -->
SELECT * FROM users;
CALL registerUser('Name', '123123123', 'secret', 'email@gmail.com');
SELECT * FROM users WHERE password = crypt('secret', password);

<!-- LIST OF USEFUL POSTGRES COMMANDS -->

<!-- List DBs -->
\l

<!-- List schemas -->
\dn

<!-- Connect to a DB -->
\c <database_name>

<!-- List tables -->
\d

<!-- Describe tables -->
\d+ users

<!-- Exit -->
\q
