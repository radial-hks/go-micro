module gomicro_grpc

go 1.14

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/micro/go-micro/v3 v3.0.0-beta.2
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/nats-io/nats-streaming-server v0.18.0 // indirect
	github.com/prometheus/client_golang v1.7.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/xanzy/go-gitlab v0.35.1 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200908134130-d2e65c121b96 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/tools v0.0.0-20200117065230-39095c1d176c // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

// https://blog.csdn.net/qq_43442524/article/details/104997539
replace (
	// github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
	google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0
)
