package protocol

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/kamalashreen/network-latency/proto/network-latency/grpc"
)

func GRPC() {
	person := &grpc.Person{
		Name:  "John Doe",
		Email: "jdoe@example.com",
	}

	fmt.Printf("Data: \n%v\nBytes on the network: %d\n", person, proto.Size(person))
}
