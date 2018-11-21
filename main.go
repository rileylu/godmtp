package main

import "godmtp/adapter"

func main(){
	xmlAdapter:=adapter.NewXmlAdapter()
	xmlAdapter.Adapt([]byte{1,2,3})
}
