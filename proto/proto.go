package proto

import (
	"encoding/json"

	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/utils"
	"github.com/nggenius/nggame/proto/c2s"
	"github.com/nggenius/nggame/proto/s2c"
)

type JsonProto struct {
}

func (j *JsonProto) GetCodecInfo() string {
	return "json"
}

func (j *JsonProto) CreateRpcMessage(svr, method string, args interface{}) (data []byte, err error) {
	r := &s2c.Rpc{}
	r.Sender = svr
	r.Servicemethod = method
	if r.Data, err = json.Marshal(args); err != nil {
		return
	}
	data, err = json.Marshal(r)
	return
}

func (j *JsonProto) DecodeRpcMessage(msg *protocol.Message) (node, Servicemethod string, data []byte, err error) {
	request := &c2s.Rpc{}

	if err = json.Unmarshal(msg.Body, request); err != nil {
		return "", "", nil, err
	}

	return request.Node, request.ServiceMethod, request.Data, nil
}

func (j *JsonProto) DecodeMessage(msg *protocol.Message, out interface{}) error {
	r := utils.NewLoadArchiver(msg.Body)
	data, err := r.GetData()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

type JsProto struct {
}

func (j *JsProto) NewProto() protocol.ProtoCodec {
	return new(JsonProto)
}

func init() {
	service.RegisterProto(&JsProto{})
}
