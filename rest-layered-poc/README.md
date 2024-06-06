# go layered REST backend PoC

The main idea is to have something like a layered architecture to start building on.

## links

Started with:

https://go.dev/doc/tutorial/web-service-gin

And read:

https://dev.to/codypotter/layered-architectures-in-go-3cg8

## Generate mocks with Mockery

Configuration in `.mockery.yaml`

```yaml
# we want generated files to be in a specific dir and package (mocks)
inpackage: False
# we want also mocks to be exported, and _test.go prevent that from happening
# even with first letter Upercases 
testonly: False
with-expecter: True
```

```bash
mockery --all --output=./mocks
```

## Testing

Runing all test in subdirectories, it is the standard but the syntax
is special so I document it to remember it one day:

```bash
go test ./...
```