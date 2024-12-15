functions, variables, constants, types, statement labels and packages
* begin with letter or _
* may have additional numbers, letters or underscores
* case matters

Twenty five key words (25) which can't be used as names
* break
* case
* const etc

Around 36 predeclared names like:
* constants
    - int
    - true
    - nil
    - itoa
* Types
    - int
    - int8
    - int16
    - int32
    - int64
    - float32
    - float64
    - complex128
    - complex64
    - bool
    - byte
    - string
    - rune
    - error
* functions
    - make
    - len
    - new
    - append
    - open/close
    - delete
    - copy
    - ...


## declarations

Names a program entity and specifies some or all of its properties.
* var
* const
* type
* func

A go program consists of one or more `.go` files.

Each source file is of this form:

* package declaration
* import declaration
* sequence of package-level declarations:
    - types
    - variables
    - constants
    - functions

Can be in any order.

### Variables

format

```
var name type = expression
```

> If expression is omitted, then the variable is initilized with the zero for the type.

zero values for the different types:

* numbers
    - 0 or 0.0, probably has a 0+0j for a complex
* strings
    - ""
* boolean
    - false
* interfaces and references (slice, pointer, map, channel, function)
    - nil
* arrays
    - zero value for the type of all elements
    - `var names []string`

Can also specify multiple variables at once.  If the type portion is omitted but the initial expression portion is specified, then variables of different types can be specified in one line.

```
var i,j = 1,0   // This is like c++. eg int i=0, j=0
var b, f, s = true, 2.3, "yo" // bool, float64, string
```

#### Short Variable Declarations

This style is used most often.  Its important to note it is a declaration and not an assinment.

format

```
variable_name := expression
```

The variable type is determined by the expression type.

```
anim := gif.GIF(...) // GIF struct
freq := rand.Float64() // float64
t := 0.0  // float64
```

> $ go doc doc.GIF


Can also do multiples

```
i,j := 0,1
```

demo

```
package main

import "fmt"


func main() {
	i, j := 0, 1
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", j, j)
}
```

output

```
int 0
int `
```

You can even swap variables in one line without using a temporary variable

NOTE: this is a tuple assignment and not a := declaration.

```
i, j = j, i
```

You can use this technique with functions that return multiple values.


```
f, err := os.Open(name)
if err != nil {
    return err
}
// ... use f
f.Close()
```

If a variable has already been declared using this
technique in the same lexical block, then upon reuse of this technique it becomes an assignment.

```
in, err := os.Open(infile)
// do stuff
// this time err is assigned rather than declared
// an assigned.
out, err := os.Create
```

The caveat is that this style must declare at least new variable.  If no new variable is declared, then simply use the assignment style.

```
in, err := os.Open(infile)
// this will not work
in, err := os.Open(anotherfile)
// use this instead
in, err = os.Open(anotherfile)
```
#### Pointers

Similar operations as c.

* `&` takes address of variable
* `*` returns value at address


```
 x := 1
 p := &x  // p, of type *int, points to x
 fmt.Println(*p) // "1"
 *p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

The zero value for ints is 0 and for pointers nil.

```
func main() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	var p *int
	fmt.Println(p)        // nil
	fmt.Println(p == nil) // "true "
	p = &x
	fmt.Println(p)        // not nil
	fmt.Println(p == nil) // "false"
}
```

output

```
true false false
<nil>
true
0xc000012028
false
```

functions can return pointers

This is interesting, In C, this would return the address of a value on the stack,
which only persists for the lifetime of the function.  I suppose since go does not have
a delete or free api and instead uses a garbage collector (GC) to free allocations
upon the last usage of a variable its ok to reurn the address of a variable
created/allocated in a function.

Each call of f() has a unique address of v.

```
func main() {

	var p = f()
	fmt.Println(" p : ", p)

}

func f() *int {
	v := 1
	fmt.Println(" v(addr): ", &v)
	return &v
}
```

Output

```
v(addr):  0xc0000a2010
p :  0xc0000a2010
```


Demo of unique address of v.

```
func main() {

	var p = f()
	fmt.Println(" p : ", p)

    fmt.Println(f() == f()) // "false"

}

func f() *int {
	v := 1
	fmt.Println(" v(addr): ", &v)
	return &v
}
```

Output

```
v(addr):  0xc000012028
p :  0xc000012028
v(addr):  0xc000012060
v(addr):  0xc000012068
false
```

Just like c/c++ it uses function pointers to manipulate variables
passed as args to a function.

```
func incr(p *int) int {
    *p++ // increments what p points to; does not change p
    return *p
}


     v := 1
     incr(&v)              // side effect: v is now 2
     fmt.Println(incr(&v)) // "3" (and v is 3)
```

The flags package (program args) uses pointers to not just get values, but to also set
values. See `projects/ch2/args/echo4.go`.


#### New

`new()` is a function that creates a variable containing an address of a type initialized to the zero value of the type.  The function returns a type of address to type parameter.

```
	p := new(int)
	fmt.Println(p)  // the address
	fmt.Println(*p) // the value
	*p = 2          // change the value to be 2 rather than the default value of 0
	fmt.Println(*p) // the value is now 2
}
```

output
```
0xc0000`2028
0
2
```

Here these two functions are equivalent.  Note in c/c++ the stack variable would
go away and this would be an error.  Here the new() function is not explictly used
but because go does not have a delete/free operation and uses an automatic garbage
collector these are equivalent.


```
func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	// test 1
	p1 := newInt()
	fmt.Println(p1)  // the address
	fmt.Println(*p1) // the value
	*p1 = 2          // change the value to be 2 rather than the default value of 0
	fmt.Println(*p1) // the value is now 2

	// test 2
	p2 := newInt2()
	fmt.Println(p2)  // the address
	fmt.Println(*p2) // the value
	*p2 = 2          // change the value to be 2 rather than the default value of 0
	fmt.Println(*p2) // the value is now 2
}
```

output

```
0xc0000aa010
0
2
0xc0000aa018
0
2
```

Note: The new function is relatively rarely used because the most common unnamed variables are of struct types, for which the struct literal syntax (§4.4.1) is more flexible.

Since `new()` is a function, it can be redefined.  In this example new is a parameter and within
the function `new()` is unavailable.

```
func delta(new, old int) int {
	// new() is unavailable here
	foo := new(int)
	*foo = 3
	return 3 - new - old
}
```

Can't do above.  The error message is:
```
invalid operation: cannot call non-function new (variable of type int)
```
#### Lifetime of variablesa

Once a variable is no longer accessible as in context stanza goes away, the garbage
collector deallocates a variable.

In this loop t persists the entire time of the context scope but x and y 
are allocated and deallocated each loop interation.

```
for t := 0.0; t <x; t+= res {
	x := math.Sin(t)
	y := math.Sin(t)
	// do stuff with x and y
}
```

The above variables are allocated on the stack.  Here is what happens when
a variable is allocated on the heap using new() or with a reference and assigned to
a global.


```
var global *int

func f() {
	var x int
	global = &x
}
```

Here x is on the stack and goes away when f() is complete.  However, the memory persists because
the address is assigned to global and its global.

```
func g() {
	y := new(int)
	*y = 1
}
```


In this case new() allocates on the heap, but the value of y is dereferenced and assigned a value without 
returning the value or the address.  Go, python and swift require the use of return to return a value, unlike R.

So, in this case the heap allocated variable is returned to memory by the garbage collector when g() completes.

#### Assignments

```
x = 1				// named variable
*p = true			// indirect variable
person.name = "bob"		// struct field
count[x] = count[x] * scale   // array or slice or map 
```

Each arithemetic and bitwise operator has an assignment operator.  As an example, an 
equivalent for the last statement is:

```
count[x] *= scale
```

Numerical variables can also be modified using postfix operators:

```
v := 1
v++   // v -> 2
v--   // v -> 1
```

#### Tuple Assignments

Several variables can be assigned at one time using tuple assignment.
All right hand variables are resolved before assigning on the left hand side.
This allows a one line swap assignment.

```
x, y = y, x
```

and when computing the greatest common divisor (GCD) of two integers

As an example:

```
func gcd(x, y int) int {

	// repeat until y result is not zero
	for y != 0 {
		fmt.Printf("loop with y: %d\n", y)
		// so for gcd(12, 6)
		// x, y = 6, 12%6 -> 0
		// x = 6
		// y = 0
		// return 6
		//
		// so for gcd(6, 12)
		// x, y = 12, 6%12 -> 6
		// x = 12
		// y = 6
		// 
		// y is not zero so it loops again
                //
		// but this is a for loop, so it assigns x=12 and y=6 so next iteration is
		// x, y = 6, 12%6 -> 0
		// x = 6
		// y = 0
		// break loop and return 6
		// return 6
		

		x, y = y, x%y
		fmt.Printf("Result of x:%d y:%d\n", x, y)
	}
	// return 12 and why not 6
	return x
}

func main() {
	println(gcd(12, 6))
	println(gcd(6, 12))
}
```

with output

```
loop with y: 6
Result of x:6 y:0
6
loop with y: 12
Result of x:12 y:6
loop with y: 6
Result of x:6 y:0
6
```


or when computing the n-th Fibonacci number iteratively : 

```
1.	0
2.	1
3.	1 (0 + 1)
4.	2 (1 + 1)
5.	3 (1 + 2)
6.	5 (2 + 3)
7.	8 (3 + 5)
8.	13 (5 + 8)
9.	21 (8 + 13)
10.	34 (13 + 21)
```

the loop counter could be done as n-1

```
     func fib(n int) int {
         x, y := 0, 1
         for i := 0; i < n; i++ {
             x, y = y, x+y
}
return x }
```

Certain expressions, such as a call to a function with multiple results, produce several values. 
When such a call is used in an assignment statement, the left-hand side must have as many variables 
as the function has results.

```
     f, err = os.Open("foo.txt")  // function call returns two values
```

Often, Functions use these additional results to indicate some kind of error, either by 
returning an error as in the call to os.Open, or a bool, usually called ok. As we’ll see in later chapters…



As with variable declarations, we can assign unwanted values to the blank identifier:

```
_, err = io.Copy(dst, src) // discard byte count
_, ok = x.(T)              // check type but discard result
```

### 2.5 Type Declarations

A type declaration/statement defines a new type which is identical to an existing type.

```
// type newtype_name underlying_type_name
type Celsius float64
```

Furthermore

```
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC	Celsius = -273.15
	FreezingC		Celsius = 0
	BoilingC		Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}
```

Every defined type has a automatic given conversion function `T(x)`.  Comparision and
assignment operations are automatic for two variables if they are of the same 
underlying type.

```
f Fahrenheit = 23
c Celsius = 32

f = c
f < c 
f == c 
```

New operations can be defined for a type.  Here a method returns a string representation 
of a celsius type.

This is a method named `String`, with a type and variable specified between the func keyword
and the method name.  

```
func (c Celsius)  String() string {
	return fmt.Sprintf("%g °C", c)
}
```

examples where it is used and not used

```
c := FtoC(212.0)

fmt.Println(c.String())  	// explicitly called

fmt.Printf("%v\n", c)   	// no need to call explictily
fmt.Printf("%s\n", c) 		// no need to call explictily
fmt.Println(c)				// no need to call explictily

fmt.Printf("%g\n", c)		// does not call String()
fmt.Println("float64(c))	// does not call String()
```




### 2.6 Packages and Files

The package tempconv has a document comment for the package.  Its the comment before the
package statement.  It should be a one line summary.  Its accessible via 

```
go doc tempconv
```

#### 2.6.1 Imports

Some terms are introduced

* package name  
  - This is the in the code with the "package" keyword.
  - It corresponds to the last entry in the `import path`. 
* import path
  - for packages which are imported from the file system, this is the path to the package.  
  - The patch includes the directory that contains the code.
  - By convention the directory name and the packae name in the code match.


Introduces two tools

* goimports
  - inserts/removes packages from the import section.
  - install
    * `go install golang.org/x/tools/cmd/goimports@latest`
    * This install to $GOPATH/bin.  Note: `$ go env` will show $GOPATH if not in ENV.  if this dir is not 
    in $PATH, simply add it.
  - usage
    * `goimports -w main.go` will rewrite main.go source file

* gofmt
  - formats code in canonical form.
  - `gofmt -w main.go` will rewrite the main.go source file

