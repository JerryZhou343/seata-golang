package protocal

import "fmt"

type RpcMessage struct {
	Id          int32
	MessageType byte
	Codec       byte
	Compressor  byte
	HeadMap     map[string]string
	Body        interface{}
}


func (r RpcMessage)String()string{
	return fmt.Sprintf("id:%d messageType:%s codec: %d compressor: %d head:%v body:%v",
		r.Id, MessageType(r.MessageType), r.Codec, r.Compressor, r.HeadMap, r.Body)
}
