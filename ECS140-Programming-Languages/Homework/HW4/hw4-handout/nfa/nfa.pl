reachable(Nfa, FinalState, FinalState, []).
reachable(Nfa, StartState, FinalState, [CurInput]) :- transition(Nfa, StartState, CurInput, NextStates), containsState(FinalState, NextStates).
reachable(Nfa, StartState, FinalState, [CurInput|RestInput]) :- transition(Nfa, StartState, CurInput, NextStates), reachableSplit(Nfa, NextStates, FinalState, RestInput).

reachableSplit(Nfa, [FirstStartState|_], FinalState, Input) :- reachable(Nfa, FirstStartState, FinalState, Input).
reachableSplit(Nfa, [_|RestStartState], FinalState, Input) :- reachableSplit(Nfa, RestStartState, FinalState, Input).

containsState(State, [H|T]) :- State = H.
containsState(State, [H|T]) :- containsState(State, T).
