package client

import (
	"fmt"
	"regexp"
	"strings"
)

func TransformMap(aMap map[string]interface{}) interface{} {
	newMap := map[string]interface{}{}
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			newMap[key] = TransformMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			newMap[key] = TransformArray(val.([]interface{}))
			fmt.Printf("broken thing: %T\n", newMap[key])
		default:
			newKey := ToSnakeCase(key)
			newMap[newKey] = aMap[key]
			fmt.Println(ToSnakeCase(key), ":", concreteVal)
		}
	}
	return newMap
}

func TransformArray(anArray []interface{}) interface{} {
	newArray := make([]interface{}, 0)

	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			newArray = append(newArray, TransformMap(val.(map[string]interface{})))
		case []interface{}:
			fmt.Println("Index:", i)
			newArray = append(newArray, TransformArray(val.([]interface{})))
		default:
			fmt.Println("Index", i, ":", concreteVal)
		}
	}
	return newArray
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
