package redis

import jsoniter "github.com/json-iterator/go"

type Packer interface {
	Marshal(val interface{}) ([]byte, error)
	MarshalToString(val interface{}) (string, error)
	Unmarshal(data []byte, target interface{}) error
	UnmarshalFromString(str string, v interface{}) error
}

var JSONPacker jsonPacker

type jsonPacker struct{}

func (jsonPacker) Marshal(val interface{}) ([]byte, error) {
	return jsoniter.Marshal(val)
}

func (jsonPacker) MarshalToString(val interface{}) (string, error) {
	return jsoniter.MarshalToString(val)
}

func (jsonPacker) Unmarshal(data []byte, target interface{}) error {
	return jsoniter.Unmarshal(data, target)
}

func (jsonPacker) UnmarshalFromString(str string, v interface{}) error {
	return jsoniter.UnmarshalFromString(str, v)
}
