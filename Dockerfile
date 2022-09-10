FROM postgres

## postgres defaults
ENV POSTGRES_DB transit 
ENV POSTGRES_USER postgres
COPY scripts/transit.sql /docker-entrypoint-initdb.d/
