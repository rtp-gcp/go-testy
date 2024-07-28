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




