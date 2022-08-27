package handler

type TypeHandler interface {
	Supports(v any) bool
	Handle(k string, v any) string
}
