# Workshop exercises


1. Create two packages in the exercises module, `rick` and `morty`. Define a exported function with no arguments in the `rick` module and call it in the other package, in another exported function. Call this function from the main.go file.

    Define a secret string in the `rick` package and use print it with the function you previously created. Take a moment to check that you cannot use it directly in the `morty` package. 
    In the `morty` package you should create a function with an integer parameter which should print `Get in the portal Rick, it's only a X minutes adventure!`, where X is the interger received.

    Adapt this function to also receive the name of the person who should go with Morty through the portal. :D

1. Generate 10 random numbers in a private function in the main.go file. Write an add function that computes the expression `x^2 + 3*x + 5` and apply it on every one of these generated numbers.

1. Implemment func f(x, y int, s1, s2 string) string that searches for the second string in the first string and only returns the start indexes of the matches that are in the [x,y] interval.

   Hint: use package strings & documentation

1. Go into the fixme module and make it build! :D

1. Create a cmd app that receives 2 parameters an integer and a character ( '+' or '*' ) and generates all the prime numbers less or equal than the first argument and applies the operation given by the second argument on all of them.
    The script should print only the result (either the sum or the product of these numbers).

    Hint: os.Args

1. Create a few constants for the area codes in the region and use a switch statement to return the area code for a given city name.

Bonus:

1. Import the ftp module linked below, connect to the test server https://dlptest.com/ftp-test/ using the credentials in the url page and create a file from random bytes, upload it under it's sha256 name and download it to check if the sha256 matches.
    
    Ftp golang module: https://pkg.go.dev/github.com/jlaffaye/ftp 

