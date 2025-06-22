# grpc-practise-2

# RESOURCE :: https://youtu.be/mPESsBfUKkc

1. Install protoc
2. Install go install grpc and proto .exe files (Check grpc - go documentation)
3. go get google.golang.org/grpc


<!-- Command to create .go files from proto3 -->
#  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative coffee_shop.proto

<!-- If you want them to be output to a different file then change the output directory -->

# protoc --go_out=./coffeeshop_protos --go_opt=paths=source_relative --go-grpc_out=./coffeeshop_protos --go-grpc_opt=paths=source_relative coffee_shop.proto

<!--  -->


protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
  	--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative



<!-- To run server file -->
go run server.go
<!-- To run client file -->
another terminal - cd to client directory

go run client.go