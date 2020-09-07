cd ./Models/protos
protoc -I . --micro_out = ../services/ --go_out=../  Prods.proto
cd .. && cd ..

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