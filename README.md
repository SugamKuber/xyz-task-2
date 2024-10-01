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
- 