package adapter

import "fmt"

type XmlAdapter struct {
}

func NewXmlAdapter() *XmlAdapter {
	return &XmlAdapter{}
}

func (xmlAdapter *XmlAdapter) Adapt(content []byte) (result []byte, err error) {
	fmt.Println("Xml Adapt")
	result=content
	return
}
