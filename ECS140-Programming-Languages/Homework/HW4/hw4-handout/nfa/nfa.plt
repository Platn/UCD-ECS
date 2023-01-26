:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(nfa).

test(nextState0, [fail]) :- containsState(0, [1, 2, 3, 4]).
test(nextState1, [nondet]) :- containsState(1, [1, 2, 3, 4]).
test(nextState2, [nondet]) :- containsState(3, [1, 2, 3, 4]).

test(nfaExp0a, [nondet]) :- reachable(expTransitions, 0, 0, []).
test(nfaExp0b, [fail]) :- reachable(expTransitions, 0, 2, []).
test(nfaExp0c, [fail]) :- reachable(expTransitions, 0, 1, [b]).
test(nfaExp1, [nondet]) :- reachable(expTransitions, 0, 2, [a]).
test(nfaExp2, [nondet]) :- reachable(expTransitions, 0, 2, [b]).
test(nfaExp3, [nondet]) :- reachable(expTransitions, 0, 1, [a, b, a]).
test(nfaExp4, [fail])   :- reachable(expTransitions, 0, 1, [a, b, a, b]).
test(nfaExp5, [nondet]) :- reachable(expTransitions, 0, 2, [a, b, a]).

test(nfaFoo1, [nondet]) :- reachable(fooTransitions, 0, 3, [a, b]).
test(nfaFoo2, [nondet]) :- reachable(fooTransitions, 0, 3, [a, c]).
test(nfaFoo3, [nondet]) :- reachable(fooTransitions, 0, 1, [a]).
test(nfaFoo4, [fail])   :- reachable(fooTransitions, 0, 3, [a, a]).
test(nfaFoo5, [fail])   :- reachable(fooTransitions, 0, 3, [a]).
test(nfaFoo6, [fail])   :- reachable(fooTransitions, 0, 1, [b]).

test(nfaLang1, [nondet]) :- reachable(langTransitions, 0, 0, [a, b, b]).
test(nfaLang2, [nondet]) :- reachable(langTransitions, 0, 1, [a, a, b]).
test(nfaLang3, [nondet]) :- reachable(langTransitions, 0, 0, [a, a, a, a, a]).
test(nfaLang4, [fail])   :- reachable(langTransitions, 0, 1, [a, a]).
test(nfaLang5, [fail])   :- reachable(langTransitions, 0, 0, [a, b, a, a]).

:- end_tests(nfa).
