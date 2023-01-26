:- initialization main.

main :-
    consult(['matrix.pl', 'matrices.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(mat).

test(are_adjacent01, [fail]) :- are_adjacent([], 1, 1).
test(are_adjacent02, [fail]) :- are_adjacent([], 1, 1).
test(are_adjacent03, [fail]) :- are_adjacent([1], 1, 1).
test(are_adjacent04, [nondet]) :- are_adjacent([1, 1], 1, 1).
test(are_adjacent05, [nondet]) :- are_adjacent([1, 2, 3], 2, 1).
test(are_adjacent06, [nondet]) :- are_adjacent([1, 2, 3], 1, 2).
test(are_adjacent07, [nondet]) :- are_adjacent([1, 2, 3], 3, 2).
test(are_adjacent08, [fail]) :- are_adjacent([1, 2, 3], 1, 1).

test_transpose(M, MT) :-
	C1 =.. [M, MX], call(C1),
	C2 =.. [MT, MXT], call(C2),
	matrix_transpose(MX, X),
	MXT == X.

test(transpose01, [nondet]) :- test_transpose(m1, m1t).
test(transpose02, [nondet]) :- test_transpose(m2, m2t).
test(transpose03, [nondet]) :- test_transpose(m3, m3t).
test(transpose04, [nondet]) :- test_transpose(m4, m4t).
test(transpose05, [nondet]) :- test_transpose(m5, m5t).
test(transpose06, [nondet]) :- test_transpose(m6, m6t).
test(transpose07, [nondet]) :- test_transpose(m7, m7t).
test(transpose08, [nondet]) :- test_transpose(m8, m8t).

test(are_neighbors01, [fail]) :- m2(M2), are_neighbors(M2, 0, 1).
test(are_neighbors02, [nondet]) :- m2(M2), are_neighbors(M2, 0, 2).
test(are_neighbors03, [nondet]) :- m2(M2), are_neighbors(M2, 0, 3).
test(are_neighbors04, [nondet]) :- m2(M2), are_neighbors(M2, 0, 4).
test(are_neighbors05, [nondet]) :- m2(M2), are_neighbors(M2, 0, 5).
test(are_neighbors06, [fail]) :- m2(M2), are_neighbors(M2, 3, 2).
test(are_neighbors07, [fail]) :- m2(M2), are_neighbors(M2, 4, 3).
test(are_neighbors08, [nondet]) :- m2(M2), are_neighbors(M2, 1, 1).
test(are_neighbors09, [fail]) :- m2(M2), are_neighbors(M2, 4, 5).
test(are_neighbors10, [fail]) :- m2(M2), are_neighbors(M2, 3, 5).

:- end_tests(mat).
