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
```

View metrics:
```
$ curl http://localhost:36701/debug/vars
{
"access_GetNodeEntities": 0,
"access_GetNodeEntitiesCount": 0,
"access_GetNodeEntity": 2,
"cmdline": ["./storage"],
"duration_ms_GetEntity_p50": 1,
"duration_ms_GetEntity_p90": 2,
"duration_ms_GetEntity_p95": 2,
"duration_ms_GetEntity_p99": 2,
"duration_ms_GetNodeEntitiesCount_p50": 0,
"duration_ms_GetNodeEntitiesCount_p90": 0,
"duration_ms_GetNodeEntitiesCount_p95": 0,
"duration_ms_GetNodeEntitiesCount_p99": 0,
"duration_ms_GetNodeEntities_p50": 0,
"duration_ms_GetNodeEntities_p90": 0,
"duration_ms_GetNodeEntities_p95": 0,
"duration_ms_GetNodeEntities_p99": 0,
"memstats": {"Alloc":
```
