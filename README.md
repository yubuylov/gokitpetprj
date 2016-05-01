# gokitpetprj

public-api - public json-api with business logic.

storage - internal json-api to access data.

##Usage:
1) You need **run both** services.
```
cd ./storage; go build && ./storage
```
```
cd ./public-api; go build && ./public-api
```

Send POST request to **public-api**:
```
curl -X POST -d '{"nid":2, "uid":1, "cvc":"26e163ffa259c3e1ece3c39d21e3d246"}' http://localhost:8001/entities
```
> Will be sended some requests to storage in parallel..


##Try storage:
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
