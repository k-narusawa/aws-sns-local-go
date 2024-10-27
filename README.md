# aws-sns-local-go

![coverage](https://raw.githubusercontent.com/k-narusawa/aws-sns-local-go/badges/.badges/main/coverage.svg)

## cli

```shell
aws sns create-topic --name my-topic --endpoint-url http://localhost:8080 --tags '[{"Key":"key1","Value":"value1"},{"Key":"key2","Value":"value2"}]'
aws sns create-topic --name my-topic --endpoint-url http://localhost:8080
```
