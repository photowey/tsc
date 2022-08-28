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

package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/photowey/tsc/internal/data/converter"
	"github.com/photowey/tsc/internal/data/jsonparser"
	"github.com/spf13/cobra"
)

const (
	JsonSuffix = "json"
)

var redirect = &cobra.Command{
	Use: "in",
	Run: func(cmd *cobra.Command, args []string) {
		handle(args)
	},
}

func handle(args []string) {
	pwd, _ := os.Getwd()
	fmt.Println("$ pwd")
	fmt.Println(pwd)
	for _, arg := range args {
		if ok := checkSuffix(arg); !ok {
			fmt.Printf("Only supports json files, now.")
			return
		}
		jsonFile := filepath.Join(pwd, arg)
		jsonData, err := ioutil.ReadFile(jsonFile)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		fmt.Printf("\n// ---------------------------------------------------------------- %s\n", arg)
		convert(jsonData)
	}
}

func convert(jsonByte []byte) {
	body := jsonparser.ParseMapB[any](jsonByte)

	model := converter.Convert(body)

	fmt.Printf(model)
}

func checkSuffix(input string) bool {
	return strings.HasSuffix(input, JsonSuffix)
}
