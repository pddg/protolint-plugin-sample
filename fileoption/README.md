# required file option rule

```bash
go build -o ./protolint-plugin-required-option .
protolint lint -plugin ./protolint-plugin-required-option testdata/*.proto
```

Result:

```bash
[testdata/without_go_package.proto:1:1] 'go_package' is required
```
