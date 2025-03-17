package common

import (
	"errors"
	"reflect"
)

func GetStructFieldValue(obj interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(obj)

	// 포인터인 경우 역참조
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 필드 가져오기
	field := v.FieldByName(fieldName)

	// 필드가 존재하지 않으면 에러 반환
	if !field.IsValid() {
		return nil, errors.New("존재하지 않는 필드: " + fieldName)
	}

	return field.Interface(), nil // 필드 값 반환
}

func GetStructFieldValues(list interface{}, fieldName string) ([]interface{}, error) {
	v := reflect.ValueOf(list)

	// 입력이 슬라이스 또는 배열이 아닌 경우 에러 반환
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, errors.New("입력 값이 슬라이스 또는 배열이 아닙니다")
	}

	// 결과 저장용 리스트
	result := make([]interface{}, 0, v.Len())

	// 슬라이스의 각 요소에서 필드 값 추출
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		value, err := GetStructFieldValue(item, fieldName)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}

	return result, nil
}

func ConvertToIntSlice(input []interface{}) []int {
	var output []int
	for _, v := range input {
		if val, ok := v.(int); ok {
			output = append(output, val)
		}
	}
	return output
}