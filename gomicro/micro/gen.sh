cd ./Models/protos
protoc -I . --micro_out = ../services/ --go_out=../  Prods.proto
protoc -I .  --go_out=plugins=grpc:../services ./*.proto
protoc -I .  --micro_out = ../services ./*.proto
cd .. && cd ..



protoc --proto_path=.  --micro_out=../services/ --go_out=.  ./Prods.proto


# https://blog.csdn.net/zoeou/article/details/86739528
# github : https://github.com/micro/micro/tree/master/cmd/protoc-gen-micro
protoc \
--plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
--plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro \
--proto_path=${GOPATH}/src:. \
--micro_out=../services/ --go_out=.  ./Prods.proto


#cd pbfiles &&  protoc -I .  --go_out=plugins=grpc:../services ./*.proto
#proc -I . --grpc-gateway_out=logtostderr=true:../services ./*.proto
#
##
#protoc -I .  --go_out=plugins=grpc:../services --validate_out=lang=go:../services  ./Models.proto
#
#cd .. && cd ..
#cp ./server/services/prod.pb.go  ./client/services/prod.pb.go
#cp ./server/services/Models.pb.go  ./client/services/Models.pb.go
#cp ./server/services/Orders.pb.go  ./client/services/Orders.pb.go
#cp ./server/services/User.pb.go  ./client/services/User.pb.go
#cd ./server/