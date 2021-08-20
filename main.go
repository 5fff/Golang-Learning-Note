package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	pointers_demo()
	structs_demo()
	struct_literals()
	arrays_demo()
	slices_demo()
	slice_literals_demo()
	slice_default_demo()
	slices_of_slices()
	range_demo()
	range_demo2()
	map_demo()
	func_as_value()
	func_closure_demo()
	fibonacciDriver()
	method_demo()
}

func pointers_demo() {
	// There's no pointer arithmetics in Go
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func structs_demo() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

//############ struct literals ##############
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func struct_literals() {
	fmt.Println(v1, p, v2, v3)
}

func arrays_demo() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices_demo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	/*
		Slices are like references to arrays
		A slice does not store any data, it just describes a section of an underlying array.
		Changing the elements of a slice modifies the corresponding elements of its underlying array.
		Other slices that share the same underlying array will see those changes.

	*/
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}

/*
	A slice literal is like an array literal without the length.
	This is an array literal:
	[3]bool{true, true, false}
	And this creates the same array as above, then builds a slice that references it:
	[]bool{true, true, false}
*/
func slice_literals_demo() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	//array of struct on the fly
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slice_default_demo() {
	/*
		When slicing, you may omit the high or low bounds to use their defaults instead.
		The default is zero for the low bound and the length of the slice for the high bound.

		For the array

		var a [10]int

		these slice expressions are equivalent:

		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = s[:]
	fmt.Println(s)

	//create slice with make
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func slices_of_slices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// https://blog.golang.org/go-slices-usage-and-internals
func append_slice() {
	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func range_demo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}

func range_demo2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

/*
Issues encountered but don't know why:
func Pic(dx, dy int) [][]uint8 {
	counter := uint8(0)
	img := make([][]uint8, dx)
	for _, s := range img {
		s = make([]uint8, dy)
	}
	for i := 0; i < dx; i++ {
		for j :=0; j < dy; j++ {
			counter++
			img[i][j]=counter
		}
	}
	return img
}

ERROR s declared but not used

Below works without error
func Pic(dx, dy int) [][]uint8 {
	counter := uint8(0)
	img := make([][]uint8, dx)
	for i := range img {
		img[i] = make([]uint8, dy)
	}
	for i := 0; i < dx; i++ {
		for j :=0; j < dy; j++ {
			counter++
			img[i][j]=counter
		}
	}
	return img
}

Found the answer: range copies the values from the slice you're iterating over
https://golang.org/ref/spec#RangeClause


*/

type GeoVertex struct {
	Lat, Long float64
}

func map_demo() {
	var m map[string]GeoVertex
	/*
		Map types are reference types, like pointers or slices, and so the
		value of m above is nil; it doesn't point to an initialized map.
		A nil map behaves like an empty map when reading, but attempts to write
		to a nil map will cause a runtime panic; don't do that.
		To initialize a map, use the built in make function:
	*/
	m = make(map[string]GeoVertex)
	m["Bell Labs"] = GeoVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	var m2 = map[string]GeoVertex{
		"Bell Labs": GeoVertex{
			40.68433, -74.39967,
		},
		"Google": GeoVertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	//Mutating Maps
	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

//Map exercise
func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, exists := ans[word]; exists {
			ans[word]++
		} else {
			ans[word] = 1
		}
	}
	return ans
}

/*
Funtion Values
Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func func_as_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

/*
Brain Alert!!

Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
For example, the adder function returns a closure. Each closure is bound to its own sum variable.

*/
func adder() func(int) int {
	sum := 0 //declare a new variable, sum
	return func(x int) int {
		sum += x //the sum is still referenced in this inner function after the adder function returns
		return sum
	}
}

/*
Explaination:
	note that every time adder() executes, a new "sum" is created and is uniquely referenced
	by the function it returned. Each execution returns a function and each of those
	functions references to their own sum, which keeps it's value after the function's execution
*/
func func_closure_demo() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
*/

/*
Exercise: Fibonacci closure
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	numA := 0
	numB := 1
	return func() int {
		tmp := numA
		numA = numB
		numB = tmp + numB
		return tmp
	}
}
func fibonacciDriver() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
Methods
Methods are functions with a "receiver" argument
*/
type Vertex2 struct {
	X, Y float64
}

//below is a method, "(v Vertex)" is the receiver argument
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// see how this method is called below
func method_demo() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs()) //it's like a member function
}
