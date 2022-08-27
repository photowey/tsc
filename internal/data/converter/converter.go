package converter

import (
	"github.com/photowey/tsc/internal/data/types"
)

type Converter interface {
	Convert(modelName string, body types.AnyMap) string
}
