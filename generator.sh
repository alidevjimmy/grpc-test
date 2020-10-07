#protoc -I greeter/greeterpb/ greeter/greeterpb/greeter.greeterpb --go-grpc_out=greeter/greeterpb
#protoc greeter/greeterpb/greeter.greeterpb --go_out=plugins=grpc:.
#protoc --go_out=

protoc --go_out=plugins=grpc:. greeter/greeterpb/*.proto
