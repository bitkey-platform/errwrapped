# errwrapped

errwrapped checks for code returning errors that are not wrapped by any wrapper library.

## Install

```
go install github.com/bitkey-platform/errwrapped/cmd/errwrapped@latest
```

## Use

```
errwrapped ./...
```

## Options

### Wrapper library

Use the `-wrapper` flag to specify a wrapper library or function.

```
errwrapped -wrapper xerrors ./...
```

If not specified, errors is specified by default.

### Ignore files

Use the `-ignore` flag to specify ignore file name patterns.

```
errwrapped -ignore vendors/,_test.go,/mock_ ./...
```
