/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stringz

import (
	"sort"
	"strings"
)

const (
	single = 1
)

var (
	DefaultEmptySeparator = ""
	DefaultSeparator      = ","
	DefaultJoiner         = "="
	DefaultSignSeparator  = "&"
)

// StringBuffer string 的 buffer 包装
type StringBuffer struct {
	buffer []string
}

// NewStringBuffer 构建 StringBuffer 实例, 并获取指针
//
// 支持 可变参数列表 0个 或者 1个
//
// 否则: 默认: 容量为: 0
func NewStringBuffer(size ...int) *StringBuffer {
	length := 0
	switch len(size) {
	case single:
		length = size[0]
	}

	return &StringBuffer{
		buffer: make([]string, length),
	}
}

// Append 追加
func (buffer *StringBuffer) Append(needle string) *StringBuffer {
	buf := append(buffer.buffer, needle)
	buffer.buffer = buf

	return buffer
}

// Join 连接字符串
//
// k joiner v
//
// e.g.:
//
// k=v
func (buffer *StringBuffer) Join(key, value, joiner string) string {
	return key + joiner + value
}

func (buffer *StringBuffer) Print() string {
	return buffer.ToString(DefaultEmptySeparator)
}

func (buffer *StringBuffer) String(separators ...string) string {
	return buffer.ToString(separators...)
}

// ToString 采用默认的分隔符 DefaultSignSeparator 转换为 string 字符串
func (buffer *StringBuffer) ToString(separators ...string) string {
	separator := DefaultSignSeparator
	switch len(separators) {
	case 1:
		separator = separators[0]
	}

	return implode(buffer.buffer, separator)
}

// ToStrings 采用指定的分隔符 separator 转换为 string 字符串
func (buffer *StringBuffer) ToStrings(separator string) string {
	return buffer.ToString(separator)
}

// ToSortString 采用默认的分隔符 DefaultSignSeparator ,自然排序之后,转换为 string 字符串
func (buffer *StringBuffer) ToSortString() string {
	cloneSlice := buffer.cloneSlice()
	sort.Strings(cloneSlice)

	return implode(cloneSlice, DefaultSignSeparator)
}

// ToSortStrings 采用默认的分隔符 DefaultSignSeparator ,自然排序之后,转换为 string 字符串
func (buffer *StringBuffer) ToSortStrings(separator string) string {
	cloneSlice := buffer.cloneSlice()
	sort.Strings(cloneSlice)

	return implode(cloneSlice, separator)
}

// Length 计算 buffer 的长度
func (buffer *StringBuffer) Length() int {
	return len(buffer.buffer)
}

func (buffer *StringBuffer) cloneSlice() []string {
	cloneSlice := make([]string, len(buffer.buffer))
	copy(cloneSlice, buffer.buffer)

	return cloneSlice
}

func implode(haystack []string, separator string) string {
	var buf strings.Builder
	for _, str := range haystack {
		if DefaultEmptySeparator == str {
			continue
		}
		buf.WriteString(str)
		buf.WriteString(separator)
	}

	return strings.TrimRight(buf.String(), separator)
}
