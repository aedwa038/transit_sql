
### Starting transit sql database

docker run --name transit_sql -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres


### listing docker containers
docker container ps