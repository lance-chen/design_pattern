// 策略模式示例：
// 有三个程序员，每一个都能进行编码、调试工作。选择一个程序员，来完成一项任务。
package main

import "fmt"

type Programmer interface {
	Code()
	Debug()
}

// ProgrammerA
type ProgrammerA struct{}

func (pa *ProgrammerA) Code() {
	fmt.Println("A is coding")
}
func (pa *ProgrammerA) Debug() {
	fmt.Println("A is debugging")
}

// ProgrammerB
type ProgrammerB struct{}

func (pa *ProgrammerB) Code() {
	fmt.Println("B is coding")
}
func (pa *ProgrammerB) Debug() {
	fmt.Println("B is debugging")
}

// ProgrammerC
type ProgrammerC struct{}

func (pa *ProgrammerC) Code() {
	fmt.Println("C is coding")
}
func (pa *ProgrammerC) Debug() {
	fmt.Println("C is debugging")
}

// Software Develop
type SoftwareDevelop struct {
	programmer Programmer
}

func (sd *SoftwareDevelop) DoTask() {
	sd.programmer.Code()
	sd.programmer.Debug()
}
func NewSoftwareDevelop(p Programmer) *SoftwareDevelop {
	return &SoftwareDevelop{
		programmer: p,
	}
}

func main() {
	sd1 := NewSoftwareDevelop(new(ProgrammerA))
	sd1.DoTask()
	sd2 := NewSoftwareDevelop(new(ProgrammerB))
	sd2.DoTask()
	sd3 := NewSoftwareDevelop(new(ProgrammerC))
	sd3.DoTask()
}
