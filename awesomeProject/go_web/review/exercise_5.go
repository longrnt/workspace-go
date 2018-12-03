package main

import "fmt"

/*
11. Create a new type called “gator”.
The underlying type of “gator” is an int.
Using var, declare an identifier “g1” as type gator (var g1 gator). Assign a value to “g1”.
Print out “g1”. Print the type of “g1” using fmt.Printf(“%T\n”, g1)

12. Adding onto the code of previous exercise;
Using var, declare an identifier “x” as type int (var x int).
Print out “x”. Print the type of “x” using fmt.Printf(“%T\n”, x)

13. Adding onto the code of previous exercise;
Can you assign “g1” to “x”? Why or why not?
=> No, because they are two different types

14. Adding onto the code of previous exercise;
You will now learn about CONVERSION.
This is called “CASTING” in a lot of other languages.
Since “g1” is of type “gator” but its underlying type is an “int”, we can use “CONVERSION” to convert the value to an int.
Now let's do it.

15. Now add a method to type gator with this signature ...
	greeting()
… and have it print “Hello, I am a gator”. Create a value of type gator. Call the greeting() method from that value.

16. Adding onto the code of previous exercise;
create another type called “flamingo”.
Make the underlying type of “flamingo” bool. Give “flamingo” a method with this signature …
	greeting()
… and have it print “Hello, I am pink and beautiful and wonderful.”
Now create a new type “swampCreature”. The underlying type of “swapCreature” is interface.
The behavior which the “swampCreature” interface defines is such that any type which has this method …
	greeting()
… will implicitly implement the “swampCreature” interface.
Create a func called “bayou” which takes a value of type “swampCreature” as an argument.
Have this func print out the greeting for whatever “swampCreature” might be passed in.

*/

func bayou(swamp swampCreature) {
	swamp.greeting()
}

type swampCreature interface {
	greeting()
}
type flamingo bool

type gator int

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}


func (g gator) greeting() {
	fmt.Println("Hello, I am a gator.")
}
func main() {
	var g1 gator
	g1 = 10
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g1.greeting()

	var g2 gator
	var f1 flamingo
	bayou(g2)
	bayou(f1)

}
