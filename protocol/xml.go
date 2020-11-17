package protocol

import "fmt"

const document = "<person><name>John Doe</name><email>jdoe@example.com</email></person>"

func XML() {
	documentByteString := []byte(document)

	fmt.Printf("xml bytes size %d", len(documentByteString))
}
