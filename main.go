package main

import (
	"fmt"
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
