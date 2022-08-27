package app

import (
	"fmt"

	"github.com/photowey/tsc/internal/data/converter"
	"github.com/photowey/tsc/internal/data/jsonparser"
)

func onData() {
	if data != "" {
		body := jsonparser.ParseMap[any](data)

		model := converter.Convert(body)

		fmt.Printf(model)
	}
}
