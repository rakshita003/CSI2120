# author: Rakshita Mathur 
# Student id: 300215340
# Date: 2023-03-24
# course: Csi 2120

% Facts about pets
pet(fido, dog, 3).
pet(spot, dog, 5).
pet(mittens, cat, 2).
pet(tweety, bird, 1).

% Facts about pet gender
male(fido).
male(spot).
female(mittens).

% Predicates
pet(Name, Species, Age) :-
    pet(Name, Species, Age, _).

species(Species, Count) :-
    findall(_, pet(_, Species, _), List),
    length(List, Count).

age_range(MinAge, MaxAge, Count) :-
    findall(_, (pet(_, _, Age), Age >= MinAge, Age =< MaxAge), List),
    length(List, Count).



