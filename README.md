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

## Generate gRPC code using Google tools

Install the Go protocol buffers plugin.

```bash
# go install google.golang.org/protobuf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc-gen-go --version
protoc-gen-go-grpc --version
```

Install the `protoc` compiler from the [release page](https://github.com/protocolbuffers/protobuf/releases/latest).
Use `apt` for Linux. For Windows, download binary zip, then copy executable and the `include` folder to the `bin` folder where `protoc-gen-go.exe` and `protoc-gen-go-grpc.exe` are located. 

```bash
apt install -y protobuf-compiler
protoc --version
```

In the root of repository, execute  te following to generate the code.

```bash
protoc --go_out=. --go-grpc_out=. --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./trading/protobuf/indicators/*.proto ./trading/protobuf/data/entities/*.proto ./trading/protobuf/*.proto
```

## Generate gRPC code using Buf

Follow [installation instructions](https://docs.buf.build/installation).
For Windows, download `buf-Windows-x86_64.exe`, `protoc-gen-buf-lint-Windows-x86_64.exe` and `protoc-gen-buf-lint-Windows-x86_64.exe` from [release page](https://github.com/bufbuild/buf/releases) and move them to the `bin` folder where `protoc-gen-go.exe` and `protoc-gen-go-grpc.exe` are located. Rename `buf-Windows-x86_64.exe` to `buf.exe`.

For Linux, use `apt install buf`.

```bash
buf --version
```

In the root of repository, execute  te following to generate the code.

```bash
buf generate
buf lint
```

## Upload CSV to client
https://blog.angular-university.io/angular-file-upload/
https://stackoverflow.com/questions/47936183/angular-file-upload
https://stackoverflow.com/questions/54971238/upload-json-file-using-angular-6
https://stackoverflow.com/questions/51070418/angular-2-upload-parse-csv
https://gyawaliamit.medium.com/reading-csv-file-on-angular-5694b64faaa5
