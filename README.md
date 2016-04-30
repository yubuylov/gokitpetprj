# gokitpetprj

## Storage - json-api to access data:

run storage:
```
cd ./storage
go build
./storage
```

try api:
```
$ curl http://localhost:36701/api/v1/1/entities
{"entities":[{"Id":1,"NodeID":1,"Payload":"asdf payload"},{"Id":2,"NodeID":1,"Payload":"asdf payload"}]}

$ curl http://localhost:36701/api/v1/1/entities/1
{"entity":{"Id":1,"NodeID":1,"Payload":"asdf payload"}}

$ curl http://localhost:36701/api/v1/1/entities/2
{"entity":{"Id":2,"NodeID":1,"Payload":"asdf payload"}}

$ curl http://localhost:36701/api/v1/1/entities/count
{"count":2}
```

View metrics:
```
$ curl http://localhost:36701/debug/vars
{
"access_GetNodeEntities": 1,
"access_GetNodeEntitiesCount": 1,
"access_GetNodeEntity": 11,
"cmdline": ["./storage"],
"duration_ms_GetEntity_p50": 1,
"duration_ms_GetEntity_p90": 1,
"duration_ms_GetEntity_p95": 1,
"duration_ms_GetEntity_p99": 1,
"duration_ms_GetNodeEntitiesCount_p50": 3,
"duration_ms_GetNodeEntitiesCount_p90": 3,
"duration_ms_GetNodeEntitiesCount_p95": 3,
"duration_ms_GetNodeEntitiesCount_p99": 3,
"duration_ms_GetNodeEntities_p50": 2,
"duration_ms_GetNodeEntities_p90": 2,
"duration_ms_GetNodeEntities_p95": 2,
"duration_ms_GetNodeEntities_p99": 2,
"memstats": {"Alloc":
```
