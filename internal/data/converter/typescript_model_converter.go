package converter

import (
	"fmt"

	"github.com/photowey/tsc/internal/data/handler"
	"github.com/photowey/tsc/internal/data/types"
	"github.com/photowey/tsc/internal/funcs"
	"github.com/photowey/tsc/pkg/stringz"
)

/*
typescript data-model converter
*/

const (
	DefaultModelName = "DataModel"
)

var _ Converter = (*TypeScriptModelConverter)(nil)

type AnyMapWrapper struct {
	name   string
	AnyMap types.AnyMap
}

// TypeScriptModelConverter a converter that convert the map[string]T to typescript data-model
type TypeScriptModelConverter struct{}

// Convert convert the map[string]T to typescript data-model
//
/*
Notes:
Use the "required" keyword to decorate the map's key to tell the converter that this property is required

- input:

{
  "name|required": "photowey",
  "age|readonly": 18,
  "balance": 10.24,
  "boy": true,
  "address": "chongqing",
  "hobby": [
    {
      "name": "badminton",
      "description": "badminton"
    }
  ],
  "university": {
    "name": "cqjtu",
    "address": "ertang"
  }
}

- output:

export interface DataModel {
	name: string;
	readonly age: number;
	balance?: number;
	boy?: boolean;
	address?: string;
	hobby?: HobbyMode[];
}

export interface HobbyModel {
  name: string;
  description?: string;
}

export interface University {
  name?: string;
  address?: string;
}
*/
func (cvt TypeScriptModelConverter) Convert(modelName string, body types.AnyMap) string {
	buffer := stringz.NewStringBuffer()
	buffer.Append(fmt.Sprintf("export interface %s {", funcs.Pascal(modelName)))
	buffer.Append("\n")
	typeHandlers := handler.TypeHandlers()

	anyMaps := make([]AnyMapWrapper, 0)

	for k, v := range body {
		for _, th := range typeHandlers {
			if th.Supports(v) {
				property := th.Handle(k, v)
				buffer.Append("    " + property)
			}
		}
		buffer.Append("\n")
		if mv, ok := v.(types.AnyMap); ok {
			anyMaps = append(anyMaps, AnyMapWrapper{
				name:   k,
				AnyMap: mv,
			})
		}
		if mv, ok := v.([]any); ok {
			if len(mv) > 0 {
				v0 := mv[0]
				if mv0, ok0 := v0.(types.AnyMap); ok0 {
					anyMaps = append(anyMaps, AnyMapWrapper{
						name:   k,
						AnyMap: mv0,
					})
				}
			}
		}
	}

	buffer.Append("}")
	buffer.Append("\n")

	for _, anyMap := range anyMaps {
		sub := cvt.Convert(anyMap.name, anyMap.AnyMap)
		buffer.Append("\n")
		buffer.Append(sub)
	}

	return buffer.Print()
}

func Convert(body map[string]any) string {
	converter := TypeScriptModelConverter{}

	return converter.Convert(DefaultModelName, body)
}
