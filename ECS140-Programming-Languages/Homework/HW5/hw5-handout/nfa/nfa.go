package nfa

import (
	"sync"
)

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.

func ReachableConcurrent(
	transitions TransitionFunction,
	start, final state,
	input []rune,
	t chan bool,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	if len(input) == 0 {
		if start == final {
			t <- true

		} else {
			t <- false
		}
		return
	}

	nextStates := transitions(start, input[0])
	if len(nextStates) >= 1 {
		for i := 0; i < len(nextStates); i++ {
			wg.Add(1)
			go ReachableConcurrent(transitions, nextStates[i], final, input[1:], t, wg)
		}
	} else {
		t <- false
	}
}

func Reachable(
	transitions TransitionFunction,
	start, final state,
	input []rune,
) bool {

	quit := make(chan bool)
	var wg sync.WaitGroup

	go Expect(transitions, start, final, input, quit, &wg)
	for i := range quit {
		if i {
			return true
		}
	}

	return false

}

func Expect(transitions TransitionFunction, start state, final state, input []rune, quit chan bool, wg *sync.WaitGroup) {
	wg.Add(1)
	go ReachableConcurrent(transitions, start, final, input, quit, wg)
	for {
		select {
		case <-quit:

		default:
			wg.Wait()
			close(quit)
		}
	}
}
