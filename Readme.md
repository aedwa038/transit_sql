
### Starting transit sql database

docker run --name transit_sql -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d transit_db


### listing docker containers
docker container ps


### Build Database
sudo docker build . -t transit_db