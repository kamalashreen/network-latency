package protocol

import "fmt"

const xmlDocument = "<person><name>John Doe</name><email>jdoe@example.com</email></person>"

// HTTP represents the HTTP 1.1 protocol using an XML document
func HTTP() {
	documentByteString := []byte(xmlDocument)

	fmt.Printf("Data: \n%s\nBytes on the network: %d\n", documentByteString, len(documentByteString))
}
