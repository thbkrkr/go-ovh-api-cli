# OVH API Go CLI

A go CLI for the OVH API.

## Example

```
echo '{
  "ID":"thbkrkr.stream",
  "partitions":1,
  "replicationFactor":1
}' | docker run --rm --env-file thb.env -i krkr/ovhapi POST /dbaas/queue/v2aaaqibavahf/topic
{"id":"thbkrkr.stream","partitions":1,"replicationFactor":1}
```

```
docker run --rm --env-file thb.env -i krkr/ovhapi GET /dbaas/queue/v2aaaqibavahf/topic
["thbkrkr.stream","thbkrkr.biiim","thbkrkr.tchiktchak"]
```