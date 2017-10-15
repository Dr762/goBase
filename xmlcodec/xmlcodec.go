package xmlcodec

import (
	"encoding/xml"
	"golang.org/x/net/websocket"
)

func xmlMarshal (v interface{}) (msg []byte,payloadType byte, err error){
	msg,err = xml.Marshal(v)
	return msg,websocket.TextFrame,nil
}

func xmlUnMarshal(msg []byte,payloadType byte,v interface{}) (err error){
	err = xml.Unmarshal(msg,v)
	return err
}

var XmlCodec = websocket.Codec{xmlMarshal,xmlUnMarshal}
