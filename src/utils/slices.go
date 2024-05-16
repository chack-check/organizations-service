package utils

func GetArrayFieldValues[T interface{}, K interface{}](arr []T, getter func(T) K) []K {
	var resultArr []K
	for _, item := range arr {
		resultArr = append(resultArr, getter(item))
	}

	return resultArr
}
