package Common

import "encoding/json"

func MarshalJson(structure any) []byte {
	payload, err := json.Marshal(structure)
	if err != nil {
		panic(err)
	}
	return payload
}
