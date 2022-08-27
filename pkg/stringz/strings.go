package stringz

import (
	"fmt"
	"strings"
)

const (
	DefaultEmptyString string = "" // 空串
)

// String 将 any 转为 string
func String(source any) string {
	return fmt.Sprintf("%v", source)
}

// ReplaceTemplate 模板替换
func ReplaceTemplate(template string, args ...any) string {
	return fmt.Sprintf(template, args...)
}

// ArrayContains 判定一个目标对象是否在指定的字符串数组对象中
//
// 返回: bool
//
// @param haystack 指定的数据对象
//
// @param needle 目标对象
//
// @return bool true 表示: 目标对象在数组中的索引值; false 表示: 不在数据组
func ArrayContains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}

	return false
}

// CloneSlice clone slice
func CloneSlice(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)

	return dst
}

// Implode 根据指定的分割符号聚合指定分切片
//
// @param haystack 指定切片
//
// @param separator 指定的分隔符
//
// @return 根据指定的分隔符-聚合-的字符串
func Implode(haystack []string, separator string) string {
	if len(haystack) == 0 {
		return DefaultEmptyString
	}

	var buf strings.Builder
	for _, str := range haystack {
		if DefaultEmptyString == str {
			continue
		}
		buf.WriteString(str)
		buf.WriteString(separator)
	}

	return strings.TrimRight(buf.String(), separator)
}

// Explode strings.Split 分割字符串
//
// @param haystack 指定字符串
//
// @param separator 指定的分隔符
//
// @return 根据指定分隔符-分割-的切片
func Explode(haystack string, separator string) []string {
	if DefaultEmptyString == haystack {
		return MakeStringSlice(0)
	}

	return strings.Split(haystack, separator)
}

// MakeStringSlice 初始化一个指定容量的 string slice
func MakeStringSlice(length int) []string {
	return make([]string, length)
}

// InitStringSlice 根据执行的参数列表 初始化一个 string slice
func InitStringSlice(opts ...string) []string {
	slice := MakeStringSlice(len(opts))
	for i, opt := range opts {
		slice[i] = opt
	}

	return slice
}

// MakeStringMap Make string map
//
// cap 传 0 个 或者 1 个参数
//
// 0.不指定容量 - 0 个参数
//
// 1.指定容量 - 1 个参数
//
// 3.指定容量 - 多 个参数
//
// 4.容错
//
// 4.1.如果: (⚠)错误的指定容量参数个数(多个), 将按照第一个参数作为指定的容量
//
// 4.2.如果: (⚠) 指定的容量参数值 <= 0 则, 按照不指定参数计算
func MakeStringMap(cap ...uint) map[string]string {
	switch len(cap) {
	case 0:
		return make(map[string]string)
	default:
		length := cap[0]
		if length <= 0 {
			return make(map[string]string)
		}
		return make(map[string]string, cap[0])
	}
}

// StringArrayIntersect 求交集
func StringArrayIntersect(haystacks ...[]string) []string {
	var inter []string
	mp := make(map[string]int)
	length := len(haystacks)

	if length == 0 {
		return make([]string, 0)
	}
	if length == 1 {
		for _, needle := range haystacks[0] {
			if _, ok := mp[needle]; !ok {
				mp[needle] = 1
				inter = append(inter, needle)
			}
		}
		return inter
	}

	for _, needle := range haystacks[0] {
		if _, ok := mp[needle]; !ok {
			mp[needle] = 1
		}
	}

	for _, haystack := range haystacks[1 : length-1] {
		for _, needle := range haystack {
			if _, ok := mp[needle]; ok {
				mp[needle]++
			}
		}
	}

	for _, needle := range haystacks[length-1] {
		if _, ok := mp[needle]; ok {
			if mp[needle] == length-1 {
				inter = append(inter, needle)
			}
		}
	}

	return inter
}

// ---------------------------------------------------------------- return bool

// IsBlankString 空串
func IsBlankString(str string) bool {
	return "" == str
}

// IsNotBlankString 非空串
func IsNotBlankString(str string) bool {
	return !IsBlankString(str)
}

// IsEmptyStringSlice 空切片
func IsEmptyStringSlice(target []string) bool {
	return len(target) == 0
}

// IsNotEmptyStringSlice 非空空切片
func IsNotEmptyStringSlice(target []string) bool {
	return !IsEmptyStringSlice(target)
}
