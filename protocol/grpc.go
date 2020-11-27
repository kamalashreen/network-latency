package protocol

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/kamalashreen/network-latency/proto/network-latency/grpc"
)

// GRPC demonstrates the gRPC protocol by using a protocol buffer message
func GRPC() {
	person := &grpc.Person{
		Name:  "John Doe",
		Email: "jdoe@example.com",
	}

	personBytes, _ := proto.Marshal(person)
	fmt.Printf("Request data: \n%v\nData on wire: \n%b\nBytes on the network: %d\n", person, personBytes, proto.Size(person))
}
