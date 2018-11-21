package adapter

import (
	"fmt"
	"testing"
)

func TestXmlAdapter_Adapt(t *testing.T) {
	xmlAdapter := NewXmlAdapter()
	if xmlAdapter != nil {
		fmt.Println(xmlAdapter.Adapt([]byte{1, 2, 3}))
	}
}
