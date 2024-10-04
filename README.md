### xyz-task-2 Docs

## Deployed link
https://stimuler.centralindia.cloudapp.azure.com/health

Above link & VM vill be shut down after 72 hours or earlier (from the current commit time)

Please read the below setup and docs to setup on local (API given below)

- /health to check if server is running
- /api/users to get sample users
- /api/generate-exercise?user_id=<user ID> to get users frequent errors

## Why ScyllaDB
- **High Throughput & Low Latency**: Handles millions of operations per second with sub-millisecond latencies.
- **Horizontal Scalability**: Scales linearly by adding nodes, ideal for billions of records.
- **Efficient Resource Usage**: Optimized for modern hardware (multi-core, large RAM).
- **Tunable Consistency**: Offers both strong and eventual consistency.

## Why Golang
- **Concurrency Support**: BuiltIn support for goroutines, ideal for high concurrency like the current application.
- **Performance**: Compiled language, faster execution than interpreted languages.
- **Simplicity**: Clean syntax & easy to maintain and deploy.
- **Efficient Resource Utilization**: Low memory footprint, suitable for distributed systems.=

## Schema Design (Detailed)

### 1. **`users` Table**
- **Purpose**: Stores basic user information.
- **Columns**:
  - `id (UUID)`: Unique identifier for each user, **Primary Key**.
  - `username (TEXT)`: Stores the username.
- **Design Explanation**:
  - Simple, lightweight table with unique user IDs. Uses **UUID** for distributed uniqueness, and data is easily spread across nodes.

### 2. **`user_errors` Table**
- **Purpose**: Logs user errors with detailed categorization and timestamps.
- **Columns**:
  - `user_id (UUID)`: Reference to the user, **Partition Key**.
  - `conversation_id (UUID)`: Identifies the conversation in which the error occurred.
  - `timestamp (TIMESTAMP)`: When the error happened.
  - `error_category (TEXT)`: General error type (e.g., authentication).
  - `error_subcategory (TEXT)`: More specific error type (e.g., password reset).
  - `error_details (TEXT)`: Full description of the error.
- **Primary Key**: 
  - **Partition Key**: `user_id`, groups errors by user.
  - **Clustering Keys**: `error_category`, `error_subcategory`, `conversation_id`, `timestamp`.
- **Design Explanation**:
  - Allows efficient logging and querying of errors by category and time. Recent errors are prioritized via descending clustering on `conversation_id` and `timestamp`, optimizing lookups for latest issues per user.

### 3. **`error_frequencies` Table**
- **Purpose**: Tracks how frequently each error occurs for each user.
- **Columns**:
  - `user_id (UUID)`: Reference to the user, **Partition Key**.
  - `error_category (TEXT)`: General error category.
  - `error_subcategory (TEXT)`: Specific error type.
  - `frequency (counter)`: Tracks the occurrence count of each error type.
- **Primary Key**: 
  - **Partition Key**: `user_id`, stores error frequencies by user.
  - **Clustering Keys**: `error_category`, `error_subcategory`.
- **Design Explanation**:
  - The **counter** type efficiently tracks error counts without complex locking mechanisms. The schema is designed for fast, scalable updates, ensuring real-time frequency tracking per user and error type.

---

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

- Optimise queries and data handlling current approach is very bad for high traffic
- Improvement caching

## Note

- Make sure you add the keyspace in scylla
- Rerun the command if the setup doesnt start after running make docker-run
