package protocol

import "fmt"

const XMLDocument = "<person><name>John Doe</name><email>jdoe@example.com</email></person>"

func HTTP() {
	documentByteString := []byte(XMLDocument)

	fmt.Printf("Data: \n%s\nBytes on the network: %d\n", documentByteString, len(documentByteString))
}
