cd Services/Protos/

protoc \
--plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
--plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
--proto_path=${GOPATH}/src:. \
--micro_out=../ --go_out=../ ./*.proto

protoc --go_out=plugins=grpc:../../ServiceGateWay test.proto
protoc --grpc-gateway_out=logtostderr=true:../../ServiceGateWay test.proto

protoc-go-inject-tag -input=../test.pb.go

cd .. && cd ..