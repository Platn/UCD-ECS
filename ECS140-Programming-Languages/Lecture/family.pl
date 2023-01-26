
% parent(X, Y) means X is Y's parent
parent(kim,holly).
parent(margaret,kent).
parent(margaret,kim).
parent(esther,margaret).
parent(herbert,margaret).
parent(herbert,jean).

% X is Y's child
child(X, Y) :- parent(Y, X).

% X is Y's grandparent.
grandparent(GrandParent,GrandChild) :- 
    parent(GrandParent,Parent), parent(Parent,GrandChild). 

% Try the following query:
% grandparent(X,_).
% How would we get rid of repetitions?

siblings1(S1, S2) :- parent(P, S1), parent(P, S2).

siblings(S1, S2) :- parent(P, S1), parent(P, S2), not(S1=S2).
