# author: Rakshita Mathur 
# Student id: 300215340
# Date: 2023-03-24
# course: Csi 2120

% Base case: if the list is empty, the sum of odd numbers is zero
sum_odd_numbers([], 0).

% If the first element is odd, add it to the sum of the rest of the list
sum_odd_numbers([X|Xs], Sum) :-
    X mod 2 =:= 1, % X is odd
    sum_odd_numbers(Xs, RestSum),
    Sum is X + RestSum.

% If the first element is even, ignore it and sum the rest of the list
sum_odd_numbers([X|Xs], Sum) :-
    X mod 2 =:= 0, % X is even
    sum_odd_numbers(Xs, Sum).
