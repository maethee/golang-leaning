
Naming Conventions for Golang Variables
These are the following rules for naming a Golang variable:

A name must begin with a letter, and can have any number of additional letters and numbers.
A variable name cannot start with a number.
A variable name cannot contain spaces.
If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
Variable names are case-sensitive (car, Car and CAR are three different variables).


//zero value