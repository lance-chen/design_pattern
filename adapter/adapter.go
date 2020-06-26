// 适配器模式示例：
// 对象适配器
// 客户端需要获取第三方组件的配置信息，已调用获取配置信息的接口为GetConfig；组件升级后，原接口被废弃，对外提供GetConfigV2
package main

import "fmt"

// Target
type TargetInterface interface {
	GetConfig() string
}

// Adaptee
type ComponentV2 struct{}

func (c ComponentV2) GetConfigV2() string {
	return fmt.Sprintln("config v2")
}

// Adapter
type Adapter struct {
	c ComponentV2
}

func (a Adapter) GetConfig() string {
	srcFormat := a.c.GetConfigV2()
	dstFormat := fmt.Sprintf("config: converted from %s", srcFormat)
	return dstFormat
}

func main() {
	var a TargetInterface
	a = Adapter{}
	fmt.Println(a.GetConfig())
}
