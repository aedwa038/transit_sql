FROM postgres

## postgres defaults
ENV POSTGRES_DB transit 
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD mysecretpassword

ADD scripts/transit.sql /docker-entrypoint-initdb.d/
