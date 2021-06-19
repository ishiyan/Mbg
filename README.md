# Mbg

## Linting

For linting, install the [golangci-lint](https://golangci-lint.run/) by runnig
`go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`.

All linters are enabled with default settings.
[Here](https://github.com/golangci/golangci-lint/blob/master/.golangci.example.yml) you can see the full configuration for all linters.

To run the `golangci-lint`, execute `golangci-lint run` in the root directory.

Sometimes you get warnings about one or another file being not 'gofumpt-ed'.
To fix it, run `gofumpt -l -w filename.go` on a troubled file..

Sometimes you get warnings about one or another file being not 'gci-ed'.
To fix it, install `go get github.com/daixiang0/gci` and run `gci -w filename.go` on a troubled file.

## Testing

Run `go test -cover ./...` in the root directory.

To benchmark, run `go test -bench=.` in a package directory

Useful links: [1](https://github.com/end0/sic4-list) [2](https://www.isindb.com/fix-cusip-calculate-cusip-check-digit/)
