### xyz-task-2 Docs

- check health
```
curl localhost:8080/api/h                                             
```

# setup scylla DB

docker pull scylladb/scylla

docker run --name scyllatest -d scylladb/scylla

docker exec -it scyllatest nodetool status

docker exec -it scyllatest cqlsh

sudo docker stop $(sudo docker ps -aq)

docker-compose exec scylla cqlsh

CREATE KEYSPACE xyz
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': '1'
};

USE xyz;

DESCRIBE TABLES;