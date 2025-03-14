package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// 读取标准输入的数据
	yamlData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// 解析YAML数据
	var jsonData interface{}
	err = yaml.Unmarshal(yamlData, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// 转换为JSON格式
	jsonBytes, err := convertToJSON(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// 输出JSON格式数据
	fmt.Println(string(jsonBytes))
}

func convertToJSON(data interface{}) ([]byte, error) {
	switch value := data.(type) {
	case map[interface{}]interface{}:
		jsonMap := make(map[string]interface{})
		for k, v := range value {
			key, ok := k.(string)
			if !ok {
				return nil, fmt.Errorf("unsupported key type: %T", k)
			}
			jsonValue, err := convertToJSON(v)
			if err != nil {
				return nil, err
			}
			jsonMap[key] = jsonValue
		}
		return json.Marshal(jsonMap)
	case []interface{}:
		jsonArray := make([]interface{}, len(value))
		for i, v := range value {
			jsonValue, err := convertToJSON(v)
			if err != nil {
				return nil, err
			}
			jsonArray[i] = jsonValue
		}
		return json.Marshal(jsonArray)
	default:
		return json.Marshal(data)
	}
}
