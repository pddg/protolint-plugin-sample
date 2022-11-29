# simple deny rule

```bash
go build -o ./protolint-plugin-simple .
protolint lint -plugin ./protolint-plugin-simple testdata/*.proto
```

Result:

```bash
[testdata/test.proto:3:1] sample visitor: given=simple
```
