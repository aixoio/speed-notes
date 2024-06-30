FROM postgres:16.3

ENV POSTGRES_USER=speed_notes
ENV POSTGRES_PASSWORD=speed_notes
ENV POSTGRES_DB=speed_notes

COPY init.sql /docker-entrypoint-initdb.d/init.sql
