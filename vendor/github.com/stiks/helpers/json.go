package helpers

import (
	"encoding/json"
	"log"
)

func ObjectToString(data interface{}) string {
	objBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(objBytes)

}

func ObjectToStringPretty(data interface{}) string {
	objBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("Unable to marshal: %s", err.Error())

		return ""
	}

	return string(objBytes)
}
