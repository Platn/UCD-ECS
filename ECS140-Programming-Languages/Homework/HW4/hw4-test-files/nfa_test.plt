:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    run_tests,
    halt.

:- begin_tests(nfa).

test(nfa01, [nondet]) :- reachable(fooTrans, 0, 1, [a]).
test(nfa02, [fail]) :- reachable(fooTrans, 0, 3, [a,b,c]).
test(nfa03, [nondet]) :- reachable(fooTrans, 0, 3, [a, b]).
test(nfa04, [fail]) :- reachable(fooTrans, 0, 3, [a, a, a]).
test(nfa05, [nondet]) :- reachable(fooTrans, 0, 3, [a, c]).

test(nfa11, [nondet]) :- reachable(barTrans, 10, 10, []).
test(nfa12, [fail]) :- reachable(barTrans, 10, 12, [b, b]).
test(nfa13, [nondet]) :- reachable(barTrans, 10, 11, [a, b, a]).
test(nfa14, [fail]) :- reachable(barTrans, 10, 12, [a, b]).

test(nfa21, [nondet]) :- reachable(mod3Trans, 51, 53, [y, x, y]).
test(nfa21, [fail]) :- reachable(mod3Trans, 51, 53, [y, x, y, x]).
test(nfa21, [nondet]) :- reachable(mod3Trans, 51, 52, [y, y, x, y]).
test(nfa21, [fail]) :- reachable(mod3Trans, 52, 52, [x, x, y]).
test(nfa21, [fail]) :- reachable(mod3Trans, 51, 53, [y, y, y, y, y]).
test(nfa21, [nondet]) :- reachable(helloTrans, 61, 72, [h,e,l,l,o,w]).
test(nfa21, [nondet]) :- reachable(helloTrans, 61, 72, [e,e,w,o,r,l,d]).
test(nfa21, [fail]) :- reachable(helloTrans, 61, 72, [e,e,w,o,r,l,d,e,e]).
test(nfa21, [nondet]) :- reachable(helloTrans, 61, 72, [h,e,w,o,o,d,o,l,d]).
test(nfa21, [fail]) :- reachable(helloTrans, 61, 72, [h,e,w,o,o,o,d,r,l,d]).

:- end_tests(nfa).
