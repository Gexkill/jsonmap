package jsonmap

const (
	JSON_FORMAT  = "j"
	VALUE_FORMAT = "v"
)

type Maps map[string]interface{}

type Map struct {
	m Maps
}
