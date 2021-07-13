package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type AnyMap map[interface{}]interface{}
type HashMap map[string]interface{}
type StringMap map[string]string

func MergeHashMap(toMap *HashMap, subMaps ...*HashMap) *HashMap {
	if toMap == nil {
		toMap = &HashMap{}
	}
	for _, subMap := range subMaps {
		if subMap != nil {
			for k, v := range *subMap {
				(*toMap)[k] = v
			}
		}
	}
	return toMap
}

func MergeStringMap(toMap *HashMap, subMaps ...*HashMap) *HashMap {
	if toMap == nil {
		toMap = &HashMap{}
	}
	for _, subMap := range subMaps {
		if subMap != nil {
			for k, v := range *subMap {
				(*toMap)[k] = v
			}
		}
	}
	return toMap
}

func ConvertStringMapToString(m *StringMap, separate string) string {
	var b bytes.Buffer
	for key, value := range *m {
		fmt.Fprintf(&b, "%s=%s%s", key, value, separate)
	}
	//fmt.Fprint(&b, "/0")
	return b.String()
}

func InHash(val interface{}, hash *HashMap) (exists bool, key string) {
	exists = false
	key = ""

	switch reflect.TypeOf(hash).Kind() {
	case reflect.Map:
		//s := reflect.ValueOf(hash)

		for k, v := range *hash {
			if reflect.DeepEqual(val, v) == true {
				key = k
				return
			}
		}
	}

	return
}

func StructToHashMap(obj interface{}) (newMap *HashMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &HashMap{}
	err = json.Unmarshal(data, newMap) // Convert to a map
	return
}

func StructToStringMap(obj interface{}) (newMap *StringMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &StringMap{}
	err = json.Unmarshal(data, newMap) // Convert to a string map
	return
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}

func StructToJson(obj interface{}) (strJson string, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return "", err
	}

	return string(data), nil
}
