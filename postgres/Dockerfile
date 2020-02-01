FROM postgres:12-alpine

ARG db_pass_env
ENV POSTGRES_USER pratik
ENV POSTGRES_PASSWORD $db_pass_env
ENV POSTGRES_DB test

COPY schema.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

