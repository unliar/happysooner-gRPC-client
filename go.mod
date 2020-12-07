module happysooner-gRPC-client

go 1.14

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
)
require github.com/unliar/happysooner-proto/v2 v2.0.9 // indirect
