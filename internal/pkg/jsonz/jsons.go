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

package jsonz

import (
	"encoding/json"
	"io"
)

func String[T any](body T) string {
	content, _ := StringE[T](body)

	return content
}

func StringE[T any](body T) (string, error) {
	bytez, err := json.Marshal(body)

	return string(bytez), err
}

func Pretty[T any](body T) string {
	content, _ := PrettyE[T](body)

	return content
}

func PrettyE[T any](body T) (string, error) {
	bytez, err := json.MarshalIndent(body, "", "\t")

	return string(bytez), err
}

func DecodeStruct[T any](reader io.Reader, structy T) error {
	if err := json.NewDecoder(reader).Decode(structy); err != nil {
		return err
	}

	return nil
}

func UnmarshalStruct[T any](data []byte, structy T) error {
	if err := json.Unmarshal(data, structy); err != nil {
		return err
	}

	return nil
}

func UnmarshalMap[T any](data []byte) map[string]T {
	maps, _ := UnmarshalMapE[T](data)

	return maps
}

func UnmarshalMapE[T any](data []byte) (map[string]T, error) {
	maps := make(map[string]T)
	if err := json.Unmarshal(data, &maps); err != nil {
		return maps, err
	}

	return maps, nil
}
