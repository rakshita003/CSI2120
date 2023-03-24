# author: Rakshita Mathur 
# Student id: 300215340
# Date: 2023-03-24
# course: Csi 2120

parent(john, mary).
parent(john, tom).
parent(mary, ann).
parent(mary, fred).
parent(tom, liz).

male(john).
male(tom).
male(fred).

female(mary).
female(ann).
female(liz).

parent(X, Y) :- 
    father(X, Y); mother(X, Y).
father(X, Y) :- 
    male(X), parent(X, Y).
mother(X, Y) :- 
    female(X), parent(X, Y).

sibling(X, Y) :- 
    parent(Z, X), parent(Z, Y), X \= Y.

grandparent(X, Y) :- 
    parent(X, Z), parent(Z, Y).

ancestor(X, Y) :- 
    parent(X, Y).
ancestor(X, Y) :- 
    parent(X, Z), ancestor(Z, Y).


