package json

import (
	std_json "encoding/json"

	"github.com/pigfall/gosdk/encoding"
)

func JsonUnmarshalFromFile(filepath string, entity interface{}) error {
	return encoding.UnMarshalByFile(filepath, entity, std_json.Unmarshal)
}
