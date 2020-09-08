cd /home/radial/Documents/CHANGE/GoProject/go_micro/gomicro/micro/Models/protos

protoc \
--plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
--plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
--proto_path=${GOPATH}/src:. \
--micro_out=../services/ --go_out=../services/  ./Prods.proto

protoc-go-inject-tag -input=../services/Prods.pb.go

cd .. && cd ..