package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type StringMap map[string]string

// ------------------------------- Merge --------------------------------------------
func MergeStringMap(toStringMap *StringMap, subStringMaps ...*StringMap) *StringMap {
	if toStringMap == nil {
		toStringMap = &StringMap{}
	}
	for _, subStringMap := range subStringMaps {
		if subStringMap != nil {
			// 迭代每个HashMap
			for k, v := range *subStringMap {
				toV := (*toStringMap)[k]
				// if the key is not exist in toMap
				if toV == "" && v != ""{
					(*toStringMap)[k] = v
				}

			}
		}
	}
	return toStringMap
}

func ConvertStringMapToString(m *StringMap, separate string) string {
	var b bytes.Buffer
	for key, value := range *m {
		fmt.Fprintf(&b, "%s=%s%s", key, value, separate)
	}
	//fmt.Fprint(&b, "/0")
	return b.String()
}

// ------------------------------- Conversion ---------------------------------------

func StructToStringMapWithTag(obj interface{}, tag string) (newMap *StringMap, err error) {

	newMap = &StringMap{}

	if obj==nil{
		return newMap, err
	}

	e := reflect.ValueOf(obj).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i).Interface()
		var strField string = ""

		strField = fmt.Sprintf("%v", field)
		key := e.Type().Field(i).Name
		if tag != "" {
			key = e.Type().Field(i).Tag.Get(tag)
		}
		(*newMap)[key] = strField

	}

	return newMap, err

}

func StructToStringMap(obj interface{}) (newMap *StringMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &StringMap{}
	err = json.Unmarshal(data, newMap) // Convert to a map
	return
}

func GetJoinedWithKSort(params *StringMap) string {

	var strJoined string

	// ksort
	var keys []string
	for k := range *params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// join
	for _, k := range keys {
		strJoined += k + "=" + (*params)[k] + "&"
	}

	strJoined = strJoined[0 : len(strJoined)-1]

	return strJoined
}

// ------------------------------- Search --------------------------------------------

func GetStringMapKV(maps StringMap) (keys []string, values []string) {
	mapLen := len(maps)
	keys = make([]string, 0, mapLen)
	values = make([]string, 0, mapLen)

	for k, v := range maps {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// ------------------------------- Filter --------------------------------------------
func FilterEmptyStringMap(mapData *StringMap) (filteredMap *StringMap) {
	filteredMap = &StringMap{}
	for k, v := range *mapData {
		if v != "" {
			(*filteredMap)[k] = v
		}
	}
	return filteredMap
}
