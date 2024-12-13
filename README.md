

### grpcurl
```bash
 grpcurl -protoset <(buf build -o -) -plaintext -d '{"someStr": "test"}' localhost:10010 test.v1.TestService/Test
```

### curl
```bash
curl --header "Content-Type: application/json" --data '{"someStr": "test"}' http://localhost:10010/test.v1.TestService/Test
```