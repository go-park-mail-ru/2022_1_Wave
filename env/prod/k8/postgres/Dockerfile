FROM postgres:13

RUN ls . -a
COPY db/migrations/1.0.0-initial_schema/ /docker-entrypoint-initdb.d/
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD music
ENV POSTGRES_DB wave
