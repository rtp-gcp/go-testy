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

