# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.13-alpine as builder
LABEL maintainer="William Ugalde <wugalde19@gmail.com>"
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN cd ./cmd/api && CGO_ENABLED=0 GOOS=linux go build -o ./../../mc1_build ./main.go


######## Start a new stage from scratch #######
FROM alpine:latest
COPY --from=builder /build/mc1_build /app/
COPY --from=builder /build/development.env /app/
COPY --from=builder /build/config/development.yml /app/config/
WORKDIR /app

EXPOSE 8000
ENTRYPOINT [ "./mc1_build", "-env=development"]
