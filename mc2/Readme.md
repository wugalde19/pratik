### COMMANDS TO RUN THE APP USING DOCKER (LOCAL ENV)
```
- docker build . -t goapp2
- docker run --rm --name goapp2-docker -p 8001:8001 goapp2:latest
```

If already have the DB container running you can use the following command
to link both containers.
```
- docker run --rm --name goapp2-docker --link pg-docker -p 8001:8001 goapp2:latest
```
