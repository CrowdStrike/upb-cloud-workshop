# Workshop exercises

1. Create a new directory and a main.go file inside it.
From terminal, run the following command: "go get golang.org/x/tour/wc"
Define the function WordCount which takes as parameter string s and returns a map that stores the frequency(number of occurences) of every word from string s.
You must import "golang.org/x/tour/wc", and the main() function should only contain the following line:
    wc.Test(WordCount)

2. Again, new directory, new main.go file. Write a function that takes as input a list(slice) of strings and returns a map
that stores the number of vowels for each string in the slice. 

e.g.: strings_list = ["ana", "mere", "peree"] will return a map with the following key-value pairs:
    "ana": 2,
    "mere": 2,
    "peree": 3

3. Create a new directory and a main.go file inside it.
Starting from an interface called Shape which defines the methods
    getName() string
    accept(ShapeVisitor)
and will be implemented by 3 structs with different fields:
    Square
      side
    Rectangle
      width
      length
    Circle
      radius

ShapeVisitor interface will contain methods:
    visitForSquare(*Square)
    visitForCircle(*Circle)
    visitForrectangle(*Rectangle)

and will be implemented by 2 structs with different functionalities:
   SurfaceCalculator
     surface
   CenterCoordinates
      X
      Y

For testing, you will instantiate one object for each defined struct and test the interaction between ShapeVisitor(SurfaceCalculator, CenterCoordinates) and Shape(Rectangle,Square,Circle)

