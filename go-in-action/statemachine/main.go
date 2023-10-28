package main

import "fmt"

// 状态机
// 1、StateMachine 结构，CurrentState, Transitions, Callback
// 2、Transition 结构，FromState, ToState, Event

// github.com/looplab/fsm
func main() {
	sm := NewStateMachine("", []Transition{
		{From: "", To: "stopped", Event: "create"},
		{From: "stopped", To: "running", Event: "start"},
		{From: "running", To: "paused", Event: "pause"},
		{From: "paused", To: "running", Event: "resume"},
		{From: "running", To: "running", Event: "restart"},
		{From: "running", To: "stopped", Event: "stop"},
	}, func(fromState, toState string, args []any) error {
		fmt.Printf("Transition from %s to %s with args %v\n", fromState, toState, args)
		return nil
	})

	sm.Trigger("create")
	sm.Trigger("start")
	sm.Trigger("pause")
	sm.Trigger("resume")
	sm.Trigger("restart")
	sm.Trigger("stop")
}
