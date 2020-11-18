package protocol

import "fmt"

const XMLDocument = "<person><name>John Doe</name><email>jdoe@example.com</email></person>"

func HTTP() {
	documentByteString := []byte(XMLDocument)

	fmt.Printf("Given data: %s\nThis corresponds to %d bytes on the network\n", documentByteString, len(documentByteString))
}
