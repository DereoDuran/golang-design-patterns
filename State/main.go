package main

import "fmt"

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off")
	sw.State = &OffState{}
}

type OffState struct {
	BaseState
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on")
	sw.State = &OnState{}
}

func NewSwitch() *Switch {
	return &Switch{State: &OffState{}}
}

func main() {
	sw := NewSwitch()
	sw.Off()
	sw.On()
	sw.Off()
	sw.Off()
}
