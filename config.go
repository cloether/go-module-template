package go_module_template

import (
	"reflect"
	"strings"
)

type Empty struct{}

func Config() string {
	path := reflect.TypeOf(Empty{}).PkgPath()
	parts := strings.Split(path, "/")
	name := parts[len(parts)-1]
	title := strings.Title(strings.Join(strings.Split(name, "-"), " "))
	return title + " Config"
}
