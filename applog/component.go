package applog

import "fmt"
import "github.com/cloether/go-module-template/applog"

// componentData records the settings for a component.
type componentData struct {
	name string
}

var cache = map[string]*componentData{}

func (c *componentData) InfoDepth(depth int, args ...interface{}) {
	args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
	applog.InfoDepth(depth+1, args...)
}

func (c *componentData) WarningDepth(depth int, args ...interface{}) {
	args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
	grpclog.WarningDepth(depth+1, args...)
}

func (c *componentData) ErrorDepth(depth int, args ...interface{}) {
	args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
	grpclog.ErrorDepth(depth+1, args...)
}

func (c *componentData) FatalDepth(depth int, args ...interface{}) {
	args = append([]interface{}{"[" + string(c.name) + "]"}, args...)
	grpclog.FatalDepth(depth+1, args...)
}

func (c *componentData) Info(args ...interface{}) {
	c.InfoDepth(1, args...)
}

func (c *componentData) Warning(args ...interface{}) {
	c.WarningDepth(1, args...)
}

func (c *componentData) Error(args ...interface{}) {
	c.ErrorDepth(1, args...)
}

func (c *componentData) Fatal(args ...interface{}) {
	c.FatalDepth(1, args...)
}

func (c *componentData) Infof(format string, args ...interface{}) {
	c.InfoDepth(1, fmt.Sprintf(format, args...))
}

func (c *componentData) Warningf(format string, args ...interface{}) {
	c.WarningDepth(1, fmt.Sprintf(format, args...))
}

func (c *componentData) Errorf(format string, args ...interface{}) {
	c.ErrorDepth(1, fmt.Sprintf(format, args...))
}

func (c *componentData) Fatalf(format string, args ...interface{}) {
	c.FatalDepth(1, fmt.Sprintf(format, args...))
}

func (c *componentData) Infoln(args ...interface{}) {
	c.InfoDepth(1, args...)
}

func (c *componentData) Warningln(args ...interface{}) {
	c.WarningDepth(1, args...)
}

func (c *componentData) Errorln(args ...interface{}) {
	c.ErrorDepth(1, args...)
}

func (c *componentData) Fatalln(args ...interface{}) {
	c.FatalDepth(1, args...)
}

func (c *componentData) V(l int) bool {
	return applog.Logger.V(l)
}

// Component creates a new component and returns it for logging. If a component
// with the name already exists, nothing will be created and it will be
// returned. SetLoggerV2 will panic if it is called with a logger created by
// Component.
func Component(componentName string) DepthLogger {
	if cData, ok := cache[componentName]; ok {
		return cData
	}
	c := &componentData{componentName}
	cache[componentName] = c
	return c
}
