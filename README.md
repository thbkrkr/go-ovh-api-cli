# OVH API Go CLI

A go CLI for the OVH API.

## Install

```
ovhapi() {
  OVH_API_CREDS=~/.ovhapi-xyz.env
  case $1 in
    GET)
      docker run --rm \
        --env-file $OVH_API_CREDS krkr/ovhapi-go-cli $@
    ;;
    *)
      cat /dev/stdin |
        docker run --rm -ti \
          --env-file $OVH_API_CREDS -i krkr/ovhapi-go-cli $@
    ;;
  esac
}

```

## Example

```
ovhapi GET /me | jq -r .name
you
```

```
echo '{
  "ID":"thbkrkr.stream",
  "partitions":1,
  "replicationFactor":1
}' | ovhapi POST /dbaas/queue/v2aaaqibavahf/topic

{"id":"thbkrkr.stream","partitions":1,"replicationFactor":1}
```
