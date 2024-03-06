package redis

import jsoniter "github.com/json-iterator/go"

type Packer interface {
	Marshal(val interface{}) ([]byte, error)
	Unmarshal(data []byte, target interface{}) error
}

var JSONPacker jsonPacker

type jsonPacker struct{}

func (jsonPacker) Marshal(val interface{}) ([]byte, error) {
	return jsoniter.Marshal(val)
}

func (jsonPacker) Unmarshal(data []byte, target interface{}) error {
	return jsoniter.Unmarshal(data, target)
}
