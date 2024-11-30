module github.com/sdfwds4/test_nrpc_qps

go 1.22

toolchain go1.22.9

require (
	github.com/nats-io/nats.go v1.31.0
	github.com/nats-rpc/nrpc v0.0.0-20240925085410-8f47f6a864d1
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace go-micro.dev/v5/api v1.18.0 => github.com/micro/go-micro/v5/api v1.18.0
