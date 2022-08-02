FROM postgres

ENV POSTGRES_PASSWORD upb
ENV POSTGRES_DB upb

COPY schema.sql /docker-entrypoint-initdb.d/
