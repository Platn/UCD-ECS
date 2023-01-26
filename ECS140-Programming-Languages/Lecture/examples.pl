
% myappend(List1, List2, List3) is true if List1 append List2 is List3.
myappend([],L,L).
myappend(L, [], L).

myappend([X|L1], L2, [X|L3]) :- myappend(L1, L2,L3).


same(X, X).



% mymember(Elem, List) is true if Elem is a member of the list List.
mymember(X, [X|_]). 
mymember(X, [_|Y]) :- mymember(X,Y).
  




%% myreverse(L1, L2).
myreverse([], []). 
myreverse([H|T], L1) :- myreverse(T,L2), myappend(L2,[H],L1).

% max(A,B,B) :- A =< B, !.
% max(A,B,A) :- A > B.

% max(A,B,B) :- A =< B, !.
% max(A,B,A).

max(A,B,C) :- A =< B, !, B=C.
max(A,B,A).


% len(List, N) iff length of the list List is N.
len([],0).
len([_|T], N) :- N>0, X is N-1, len(T,X).


% Y is a suffix of X.
suffix(X, Y) :-  append(_, Y, X).

prefix(X, Y) :-  append(Y, _, X).

% Y is a sublist of X
sublist(Xs, Ys) :-  suffix(Xs, Zs), prefix(Zs, Ys).


% sublist([a, b, c, d, e], [c, d]).
% sublist([a, b, c, d, e], Ys).
% sublist(Xs, Ys).


% another length predicate using accumulators.
% Tail recursive!
accLen([_|T],A,L) :-  Anew is A+1, accLen(T,Anew,L).
accLen([],A,A).

leng(List,Length) :- accLen(List,0,Length).

% length of a list that works when the list is supplied.
len1([], 0).
len1([_|L], N) :- len1(L, N1), N is N1 + 1.

% length of a list that works when the length is supplied.
len2([], 0).
len2([_|L], N) :- N > 0, N1 is N - 1, len2(L, N1).

% length that works in both modes. 
lenBoth(L,N) :- ( var(N) ->  len1(L,N); len2(L,N) ).
% Note that var(Term) is true if Term is a free variable.
% Other predicates to verify type of a term can be found at:
% http://www.swi-prolog.org/pldoc/man?section=typetest 

% Try the following queries:
% lenBoth(X, 2).
% lenBoth([1, 2,3], X).
% lenBoth([1|T], 2).


accMax([H|T],A,Max) :-
    H > A,
    accMax(T,H,Max).
 
 accMax([H|T],A,Max) :-
    H =< A,
    accMax(T,A,Max).
 
 accMax([],A,A).

% accMax([1,0,5,4],0,Max).

% Only works for non-empty lists.
maxlist(List,Max) :-
    List = [H|_],
    accMax(List,H,Max).

addone([], []).
addone([H1|T1], [H2| T2]) :- (is(H2,H1+1); is(H1,H2-1)), addone(T1, T2).

% addone([[1,2,3], X).
% addone([X, [1,2,3]).

% height(Tree, N) N is the height of the tree.
height(void, 0).

height(tree(_, L, R), H) :- height(L, H1), 
                            height(R, H2), 
                            H1 > H2, 
                            H is H1+1. 

height(tree(_, L, R), H) :- height(L, H1), 
                            height(R, H2), 
                            H1 =< H2, 
                            H is H2+1. 

% height(tree(2,void,void), X).
% height(tree(4,tree(2,void,void),tree(10,void,void)), X).
                            
% insertTree(Key, BeforeTree, AfterTree).
insertTree(K,void,tree(K,void,void)).   % base case 
   
insertTree(K,tree(K,L,R),tree(K,L,R)).  % no duplicates

insertTree(K,tree(N,L,R),tree(N,Lnew,R)) :- K < N, insertTree(K, L, Lnew).  % to the left

insertTree(K,tree(N,L,R),tree(N,L,Rnew)) :- K>N, insertTree(K, R, Rnew). % to the right

% insertTree(2, tree(2,void,void), X).
% insertTree(2, tree(4,tree(2,void,void),tree(10,void,void)), X).

memberTree(K, tree(K,_,_)). 
memberTree(K, tree(N,L,_)) :- K < N, memberTree(K, L).
memberTree(K, tree(N,_,R)) :- K > N, memberTree(K, R).  

% pivot(List, P, LessThanEq, Gt).

% If pivot element is >= than X, add X to the left list.  
pivot([X|Xs], Y, [X|Ls], Rs) :- X =< Y, pivot(Xs, Y, Ls, Rs).
% If pivot element is <= than X, add X to the right list.  
pivot([X|Xs], Y, Ls, [X|Rs]) :- X > Y, pivot(Xs, Y, Ls, Rs).
% base case.
pivot([],_,[],[]).

% quicksort(X, Y) is true when Y is the sorted list for X.
quicksort([X|Xs],Ys) :-
    pivot(Xs,X,Left,Right),
    quicksort(Left,Ls),
    quicksort(Right,Rs),
    append(Ls,[X|Rs],Ys).

quicksort([],[]).
  


  
