# grpc-gateway

## client

```bash
❯ grpcurl -protoset <(buf build -o -) -d '{"name": "Jane2"}' -plaintext 0.0.0.0:8081 helloworld.Greeter/SayHello
{
  "message": "Jane2 world"
}
```
