maxlist(List,Max) :-
    List = [H|_],
    accMax(List,H,Max).

addone([], []).
addone([H1|T1], [H2| T2]) :- (is(H2,H1+1); is(H1,H2-1)), addone(T1, T2).