package app

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/photowey/tsc/internal/data/converter"
	"github.com/photowey/tsc/internal/data/jsonparser"
	"github.com/spf13/cobra"
)

const (
	JsonSuffix = "json"
)

var redirect = &cobra.Command{
	Use: "<",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			if ok := checkSuffix(arg); !ok {
				fmt.Printf("Only supports json files, now.")
				return
			}

			jsonData, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}

			fmt.Printf("\n// ---------------------------------------------------------------- %s\n", arg)
			convert(jsonData)
		}
	},
}

func convert(jsonByte []byte) {
	body := jsonparser.ParseMapB[any](jsonByte)

	model := converter.Convert(body)

	fmt.Printf(model)
}

func checkSuffix(input string) bool {
	return strings.HasSuffix(input, JsonSuffix)
}
