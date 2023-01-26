% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
% base cases
are_adjacent([A,B|_], A, B).
are_adjacent([B,A|_], A, B).
% recursive case
are_adjacent([_|Rest],A, B) :- are_adjacent(Rest, A, B).

% matrix_transpose(Matrix, Answer) returns true iff Answer is the transpose of
% the 2D matrix Matrix.
% base cases
matrix_transpose([], []).
matrix_transpose([], _, []).
% reformat to 3 arg form which takes first row as arguement
matrix_transpose([Row|RestRows], Result) :- matrix_transpose(Row, [Row|RestRows], Result).
matrix_transpose([_|RestRow], Matrix, [ResultRow|ReresultRestRows]) :- first_col_to_row(Matrix, ResultRow, RestMatrix), matrix_transpose(RestRow, RestMatrix, ReresultRestRows).

% first_col_to_row(Matrix, Rest, RestMatrix)
first_col_to_row([], [], []).
first_col_to_row([[H|T]|Rest], [H|Hs], [T|Ts]) :- first_col_to_row(Rest, Hs, Ts).

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
are_neighbors(Matrix, A, B) :- matrix_transpose(Matrix, MatrixTransposed), append(Matrix, MatrixTransposed, ConcatMatrix), are_row_neighbors(ConcatMatrix, A, B).

% are_row_neighbors(Matrix, A, B) returns true if A and B are adjacent in any of the rows
are_row_neighbors([Row|_], A, B) :- are_adjacent(Row, A, B).
are_row_neighbors([_|RestRows], A, B) :- are_row_neighbors(RestRows, A, B).


