* COMMANDS TO RUN THE APP USING DOCKER (LOCAL ENV)
- docker build . -t goapp
- docker run --rm --name goapp-docker --link pg-docker -p 8000:8000 goapp:latest
