package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	DEBUGGER    bool
	MENU_CHOICE int64
}

func tryingGoBasics() {
	settings := Settings{DEBUGGER: true, MENU_CHOICE: 1}

	menu := "Main menu:\n  Enter:\n 0: Stop\t\t 1: print\t\t 2: input\n 3: loops\t\t 4: array\t\t 5: map\n 6: function\t\t 7: mutableImmutable\t 8: pointers\n 9: struct\t\t \n"
	scanner := bufio.NewScanner(os.Stdin)
	var inputCmd int64
	stop := false
	for !stop {
		fmt.Println(menu)
		if !settings.DEBUGGER {
			scanner.Scan()
			inputCmd, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		} else {
			stop = true
			inputCmd = int64(settings.MENU_CHOICE)
		}
		switch inputCmd {
		case 0:
			stop = true
		case 1:
			printStuff()
		case 2:
			inputStuff()
		case 3:
			loops()
		case 4:
			arrayAndSliceStuff()
		case 5:
			mapStuff()
		case 6:
			functionStuff()
		case 7:
			mutableImmutableStuff()
		case 8:
			pointersStuff()
		case 9:
			structStuff()
		default:
			fmt.Println("not valid input")
		}
	}
	fmt.Println("Bye bye")
}

func printStuff() {

	/* Declaration: */
	fmt.Println("#########################")
	fmt.Println("Declaration of the variables ( var a int = 10 or  a := 10) ")
	var a int = 10
	/* := is a shortcut for var and the type */
	//a1 := 10
	b := 3
	/* Casting: */
	fmt.Println("#########################")
	fmt.Println("Type casting variables: c := float64(a) / float64(b)")
	c := float64(a) / float64(b)

	fmt.Println("#########################")
	fmt.Printf("Type of the variable (%%T): a = %T b = %T c = %T \n", a, b, c)
	var out string = fmt.Sprintf("Printing the numbers (%%d or %%f etc...): a = %d , b = %d , c = %5f", a, b, c)
	fmt.Println(out)

	stringStuff()
}
func stringStuff() {
	fmt.Println("#########################")
	var str string = "/father/child1/Your name is Heisenberg"
	splitter := func(c rune) bool {
		return c == '/'
	}
	fmt.Println("initial string: " + str)
	fields := strings.FieldsFunc(str, splitter)
	fmt.Printf("Fields are: %q\n", fields)
	fmt.Printf("Field[2]: %q\n", fields[2])
}
func inputStuff() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your birth date: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Input: %q\n", input)
	birthYear, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2030\n", 2030-birthYear)
}

/**
 * This count how many odds or even there are between 2 numbers
 */
func loops() {
	fmt.Println("##################")
	fmt.Println("even/odd counter with loops:")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a:\t")
	scanner.Scan()
	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Print("Enter b:\t")
	scanner.Scan()
	b, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if b < a {
		tmp := b
		b = a
		a = tmp
	}

	evenCounter := 0
	oddCounter := 0
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			evenCounter++
		} else {
			oddCounter++
		}
	}
	switch {
	case evenCounter > oddCounter:
		fmt.Println("more even than odd")
	case evenCounter < oddCounter:
		fmt.Println("more odd than even")
	case evenCounter == 0 || oddCounter == 0:
		fmt.Println("strange case")
	}
	fmt.Printf("Between %d and %d there are %d odds and %d evens\n", a, b, oddCounter, evenCounter)
}

func arrayAndSliceStuff() {
	fmt.Println("#################")
	fmt.Println("Array declaration (var arr [5]int) it initializes the array with all 0")
	const n int = 5
	var arr [n]int
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	arr2 := [7]int{12, 31, 22, 1, 3, 4, 5}
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Printf("sum of the arr2 values: %d\n", sum)
	fmt.Println("#####################")
	fmt.Printf("Slices declaration (var s [empty_here]int = x[from:to])")
	var s []int = arr2[1:3]
	fmt.Println("s (from 1 to 3 of var2) : ")
	fmt.Println(s)
	fmt.Printf("capacity: %d\nlenght: %d\n", cap(s), len(s))
	fmt.Println(s[:5])

	fmt.Println("Declaring a slice directly (but it will create an array first)")
	var s1 []int = []int{5, 6, 7, 8, 9}
	fmt.Println(s1, cap(s1), len(s1))
	fmt.Println("If I want a slice with more elements I can use append func but it will return a new slice")
	s2 := append(s1, 10, 11, 12)
	fmt.Println(s2)

	fmt.Println("make function can be used to declare a slice")
	s3 := make([]int, 5)
	fmt.Printf("This is a slice because the square brackets are empty: %T \n", s3)

	var s4 []int = []int{1, 3, 4, 56}
	fmt.Println("This is how range works:\n it returns the position and the element of the slice from the start to the end of the slice")
	for i, element := range s4 {
		fmt.Printf("s4[%d]: %d\n", i, element)
	}

	fmt.Println("Example: Looking for duplicates in a slice")

	var s5 []int = []int{1, 3, 4, 5, 4, 2, 1, 1}
	for i, element := range s5 {
		for j, element2 := range s5 {
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}
	}

}

func mapStuff() {
	fmt.Println("########################")
	fmt.Println("make can be used to declare empty maps")
	mp1 := make(map[string]int)
	fmt.Println(mp1)
	fmt.Println("Declaring map (var mp map[keyType]valueType = map[keyType]valueType{key1:value1, ...}")
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pen":    6,
		"orange": 7,
	}
	fmt.Println(mp)
	fmt.Println("Changing value of apple, adding pinapple key and deleting orange key")
	mp["apple"] = 9000
	mp["pinapple"] = 42
	delete(mp, "orange")
	fmt.Println(mp)

	fmt.Println("Take values from map (val, ok := mp[key] ) ")
	fmt.Println("if the key doesn't exist val will be the default value for that value typ and ok will be false")
	val1, ok1 := mp["watermelon"]
	fmt.Println("searching for watermelon:\t var, ok = ", val1, ok1)
	val2, ok2 := mp["apple"]
	fmt.Println("searching for apple:\t var, ok = ", val2, ok2)

}

func functionStuff() {
	fmt.Println("########################")
	fmt.Println("The syntax is func functionName(parameter1,parameter2 type1, parameter3 type2) (result1,result2 type3)")
	fmt.Println("To return values assign the values to result1,2.. variables")
	fmt.Println("calling testFunction(4,5,6,3)")
	z1, z2, z3 := testFunction(4, 5, 6, 3)
	fmt.Printf("4+5 = %d\n6/3 = %d 4*5 = %d", z1, z2, z3)
}

/* z1=x+y z2=f1/f2 z3=x*y */
func testFunction(x, y int, f1, f2 float64) (z1, z2, z3 int) {

	fmt.Println("I can have a func variable myFunc := func(parameter1,par2 type1) result1 type2 { func body }")
	mul := mulDeferFunction
	myfunc := func(a int) {
		fmt.Printf("a = %d\n", a)
	}
	fmt.Println("I can pass function to another function as parameter with this syntax:")
	fmt.Println("func myFunc2(myFunc1 func(MF1parameters) MF1results, MF2parameter2) MF2result2")
	myfunc(7)
	z1 = x + y
	defer fmt.Println("the defer keyword allows me to execute the action before returning values")
	z2 = int(f1 / f2)

	z3 = testInner(mul, x, y)
	return

}
func testInner(myFunc func(int, int) int, a, b int) int {
	return myFunc(a, b)
}

func mulDeferFunction(x, y int) int {
	// defer fmt.Println("Adding x and y")
	sum := 0
	summing := func(a int) {
		fmt.Printf("Adding %d at sum that is = %d\n", a, sum)
		sum += a
	}

	for i := 0; i < y; i++ {
		defer summing(x)
	}

	return sum
}

func mutableImmutableStuff() {
	fmt.Println("##############################")
	fmt.Println("standard variables and arrays are immutable \"var1 := var2 when i modify var1 var2 wont be modified\"")
	fmt.Println("int:")
	x1 := 10
	y1 := x1
	y1 = 12
	fmt.Println("x1 = ", x1, "\ty1 = ", y1)
	fmt.Println("string:")
	str1 := "Hello"
	str2 := str1
	str2 = "byebye"
	fmt.Println("str1 = ", str1, "\tstr2 = ", str2)
	fmt.Println("array:")
	a1 := [3]int{3, 4, 5}
	a2 := a1
	a2[0] = 100
	fmt.Println("a1 = ", a1, "\ta2 = ", a2)

	fmt.Println("slice and map variables are mutable \"var1 := var2 when i modify var1 var2 will be modified\"")
	fmt.Println("slice:")
	var s1 []int = []int{3, 4, 5}
	s2 := s1
	s2[0] = 100
	fmt.Println("s1 = ", s1, "\ts2 = ", s2)
	fmt.Println("map:")
	var m1 map[string]int = map[string]int{"Hello": 3}
	m2 := m1
	m2["Bye"] = 100
	fmt.Println("m1 = ", m1, "\tm2 = ", m2)

}

func pointersStuff() {
	fmt.Println("######################")
	fmt.Println("pointers and functions")
	change1 := func(str string) string {
		fmt.Println("this is the input of the function " + str)
		str = "Changed!"
		return str
	}
	change2 := func(str *string) {
		*str = "Changed!"
	}
	str1 := "Hello!"
	fmt.Printf("str1 = %q   \t", str1)
	change1(str1)
	fmt.Printf(" str1 after change1 = %q\n", str1)
	fmt.Printf("str1 = %q   \t", str1)
	change2(&str1)
	fmt.Printf(" str1 after change2 = %q\n", str1)

}

// Point is : x and y are the coordinates, name is the name of the pointer
type Point struct {
	x    int32
	y    int32
	name string
}

// Shape has area
type Shape interface {
	area() float64
	//perimeter() float64
}

// Circle is : radius and center to represent a circle
type Circle struct {
	radius float64
	// center is the origin of the circle
	center *Point
}

func (c *Circle) setRadius(r float64) {
	c.radius = r
}
func (c Circle) getRadius() float64 {
	return c.radius
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle is : height and width
type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}
func garbageCollectorIssue() *Point {
	p1 := &Point{x: 10, y: 20}
	p1.name = "MyPointer"
	fmt.Println("Inside the function")
	fmt.Printf("%p -> ", p1)
	fmt.Println(*p1)

	return p1
}

func structStuff() {

	fmt.Println("###################")
	fmt.Println("define a struct: type StructName struct { body }")
	fmt.Println("istantiate an istance of the struct: var varName StructName := StructName{par1,par2,...}")
	fmt.Println("I can also use this: varName := StructName{par1:0} // if I don't put all the parameters they will be the default value ")
	var p1 Point = Point{1, 2, "Point 1"}
	p2 := &Point{x: 10}
	fmt.Println(p1, p2)
	p2.name = "Point 2"

	changeX1 := func(pt Point) {
		pt.x = 100
	}
	changeX2 := func(pt *Point) {
		pt.x = 100
	}
	fmt.Println("p2 = ", p2)
	changeX1(*p2)
	fmt.Println("changing x with changeX1 : ", p2)
	changeX2(p2)
	fmt.Println("changing x with changeX2 : ", p2)

	c1 := Circle{4.56, &p1}
	fmt.Println("This is your circle: ", c1)
	fmt.Println("The center is: ", c1.center, "\tThe x = ", c1.center.x, "\tThe y = ", c1.center.y)
	fmt.Println("Radius before set  = ", c1.getRadius())
	c1.setRadius(100)
	fmt.Println("Radius after set  = ", c1.getRadius())

	r1 := Rectangle{6.9, 8.9}
	fmt.Println(r1)
	shapes := []Shape{c1, r1}
	for _, shape := range shapes {
		fmt.Println(shape.area())
		// fmt.Println(getArea(shape))

	}
	fmt.Printf("area of shape[0] : %f", getArea(shapes[0]))
	fmt.Println("Garbage collector example!:")
	p3 := garbageCollectorIssue()
	fmt.Println("outside the function")
	fmt.Printf("%p -> ", p3)
	fmt.Println(*p3)
	fmt.Println("The compiler recognize that the pointer of MyPointer is used outside the function and allocate it in the heap instead of the stack")
	fmt.Println("In this way the garbage collector doesn't clean the memory pointed by the variable when the function ends")
}
