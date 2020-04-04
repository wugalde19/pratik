### COMMANDS TO RUN THE APP USING DOCKER (LOCAL ENV)
```
- docker build . -t goapp
- docker run --rm --name goapp-docker -p 8000:8000 goapp:latest
```

If already have the DB container running you can use the following command
to link both containers.
```
- docker run --rm --name goapp-docker --link pg-docker -p 8000:8000 goapp:latest
```
