package main

import (
	"fmt"
)

type StateMachine interface {
	Trigger(event string, args ...any) error
}

var _ StateMachine = (*stateMachine)(nil)

type stateMachine struct {
	state        string
	transitions  []Transition
	eventHandler EventHandler
}

// Trigger implements StateMachine.
func (sm *stateMachine) Trigger(event string, args ...any) error {
	t := sm.findTransition(sm.state, event)
	if t == nil {
		return fmt.Errorf("no transition found for event %s", event)
	}
	err := sm.eventHandler(sm.state, t.To, args)
	if err != nil {
		return fmt.Errorf("event handler failed: %w", err)
	}
	sm.toState(t.To)
	return nil
}

type EventHandler func(fromState, toState string, args []any) error

type Transition struct {
	From  string
	To    string
	Event string
}

func NewStateMachine(initState string, transitions []Transition, eventHandler EventHandler) StateMachine {
	return &stateMachine{
		state:        initState,
		transitions:  transitions,
		eventHandler: eventHandler,
	}
}

func (sm *stateMachine) toState(toState string) {
	sm.state = toState
}

func (sm *stateMachine) findTransition(fromState, event string) *Transition {
	for _, t := range sm.transitions {
		if t.Event == event && t.From == fromState {
			return &t
		}
	}
	return nil
}
