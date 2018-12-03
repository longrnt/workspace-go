package main

import "fmt"

/*
4. Create a new type called person which is STRUCT that stores fName and lName
Using the STRUCT “person”, using a composite literal create a value of type person and assign it to a variable with the identifier “p1”;
print out “p1”; print out just the field fName for “p1”

5. Take the STRUCT “person” in the previous exercise and add a field called “favFood” that stores a slice of string;
for the variable “p1” use a composite literal to add values to the field favFood;
print out the values in favFood;
also print out the values for “favFood” by using a for range loop

6. Add a method to type “person” with the identifier “walk”.
Have the func return this string: <person’s first name> is walking.
Remember, you make a func a method by giving the func a receiver.
	func (r receiver) identifier(parameters) (returns) {
		<code>
	}
To return a string, use fmt.Sprintln. Call the method assigning the returned string to a variable with the identifier “s”. Print out “s”.
 */

 type person1 struct {
 	fName string
 	lName string
 }

 type person2 struct {
 	person1
 	favFood []string
 }

 func (p person2) walk() string {
 	return fmt.Sprint(p.fName, " is walking.")
 }
func main() {
	p1 := person1{
		"Todd",
		"McLeod",
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)

	p2 := person2{
		person1{"Long", "Tran"},
		[]string{"Orange", "Kiwi", "Strawberry", "Apple"},
	}

	fmt.Println(p2.favFood)

	for i,v := range p2.favFood {
		fmt.Println(i,v)
	}

	walkingString := p2.walk()

	fmt.Println(walkingString)
}

