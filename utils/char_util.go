package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// GetIntFromString 获取字符串中的数字字符串并转换为int
func GetIntFromString(str string) int {
	var result int
	for _, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		}
	}
	return result
}

// GetIntsFromString 获取字符串中的数字，多组数字返回int数组
func GetIntsFromString(str string) []int {
	var result []int
	var tmp int
	for _, v := range str {
		if v >= '0' && v <= '9' {
			tmp = tmp*10 + int(v-'0')
		} else {
			if tmp != 0 {
				result = append(result, tmp)
				tmp = 0
			}
		}
	}
	if tmp != 0 {
		result = append(result, tmp)
	}
	return result
}

// GetFloat64sFromString 获取字符串中的数字，多组数字返回float64数组
func GetFloat64sFromString(str string) []float64 {
	var result []float64
	var tmp string
	for _, v := range str {
		if (v >= '0' && v <= '9') || v == '.' {
			tmp += string(v)
		} else {
			if tmp != "" {
				tmp = strings.TrimRight(tmp, ".")
				fmt.Println("tmp:", tmp)

				result = append(result, StrToFloat64(tmp))
				tmp = ""
			}
		}
	}
	if tmp != "" {
		tmp = strings.TrimRight(tmp, ".")
		result = append(result, StrToFloat64(tmp))
		tmp = ""
	}
	return result
}

// GetCharsFromString 获取字符串中的字母字符，允许指定排除指定的字符不输出，多个单词返回字符串数组
func GetCharsFromString(str string) []string {
	var result []string
	var tmp string
	for _, v := range str {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			tmp += string(v)
		} else {
			if tmp != "" {
				result = append(result, tmp)
				tmp = ""
			}
		}
	}
	if tmp != "" {
		result = append(result, tmp)
	}
	fmt.Println(result)
	return result
}

// IsInStringSlice 判断字符串是否在字符串数组中
func IsInStringSlice(str []string, substr string) bool {
	for _, v := range str {
		if v == substr {
			return true
		}
	}
	return false
}

// IsInString 判断字符串中是否包含指定字符串中的任意一个字符串
func IsInString(str string, substr []string) bool {
	for _, v := range substr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

// StrToInt 字符串转int
func StrToInt(str string) int {
	var result int
	for _, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		}
	}
	return result
}

// StrToFloat64 字符串转float64
func StrToFloat64(str string) float64 {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}
	rel, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("StrToFloat64 err:", err)
		return 0
	}
	return rel
}
