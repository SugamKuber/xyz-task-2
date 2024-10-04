### xyz-task-2 Docs

## Deployed link
https://stimuler.centralindia.cloudapp.azure.com/health

Will be shut down after 72 hours from the current commit time


## Server setup

```
make fmt


make docker-run
docker-compose down
```

## scylla DB setup

docker pull scylladb/scylla
docker run --name scyllatest -d scylladb/scylla
docker exec -it scyllatest nodetool status
docker exec -it scyllatest cqlsh
sudo docker stop $(sudo docker ps -aq)

## Database config
```
docker-compose exec scylla cqlsh

CREATE KEYSPACE xyz
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': '1'
};
USE xyz;
DESCRIBE TABLES;
```



## API

```
curl localhost:8080/api/generate-exercise?user_id=<user ID>

curl localhost:8080/api/users

curl localhost:8080/health
```


## Improvments

- Optimise queries and data handlling
- Improvement cachin

## Note

- Make sure you add the keyspace in scylla
- Rerun the command if the setup doesnt start after running make docker-run