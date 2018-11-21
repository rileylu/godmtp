package mainflow

import (
	"godmtp/adapter"
	"godmtp/channel"
)

type MainFlow struct {
	Adapter adapter.Adapter
	Channel channel.Channel
}

func (mainFlow *MainFlow) Start() error {
	content:=[]byte{1,2,3}
	adaptedResult,_=mainFlow.Adapter.Adapt(content)


	return nil
}
