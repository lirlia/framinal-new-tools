# grpc-gateway

## client

```bash
❯ grpcurl -protoset <(buf build -o -) -d '{"name": "Jane2"}' -plaintext 0.0.0.0:8081 helloworld.Greeter/SayHello
{
  "message": "Jane2 world"
}

❯ curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " Jane"}'
```
