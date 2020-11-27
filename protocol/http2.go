package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// HTTP2 represents the HTTP 2 protocol using a binary encoded XML document
func HTTP2() {
	buf := new(bytes.Buffer)
	for _, b := range []byte(xmlDocument) {
		err := binary.Write(buf, binary.LittleEndian, b)
		if err != nil {
			return
		}
	}

	fmt.Printf("Request data: \n%v\nData after passing through binary framing layer: \n%b\nBytes on the network: %d\n", xmlDocument, buf.Bytes(), buf.Len())
}
