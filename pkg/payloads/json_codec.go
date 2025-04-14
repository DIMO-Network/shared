package payloads

import "encoding/json"

// JSONCodec is often used with CloudEvents to encode them into json or decode them from json to a struct
//
// Example:
// goka.Input(goka.Stream("topic.contracts"), new(shared.JSONCodec[shared.CloudEvent[processor.ContractEventData]]), proc.RegistryCallback), ... )
//
// var speedEventCodec = &shared.JSONCodec[shared.CloudEvent[segmenter.SpeedData]]{}
type JSONCodec[A any] struct{}

func (c *JSONCodec[A]) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *JSONCodec[A]) Decode(data []byte) (interface{}, error) {
	value := new(A)
	return value, json.Unmarshal(data, value)
}
