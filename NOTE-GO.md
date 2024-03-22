- For starters : Declare _main_ package in same directory.

- Then import _fmt_ package as well, which mainly contains all the pre-written function that will be used in go language. :smile

- Running go file

```
$ go run .
```

- In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

- **Comments** in _go_:

  - Single line comments starts with `//`.
  - multline comments `/*` and `*/`.

- Sometimes we need to **build** our files. - `go build <flename>` command is used for it. - And you can run it via using `./<filename>`. Yup, just like that. :) -[Gobyexample](https://gobyexample.com/)

- **Variables** in _go_.

  - There are **2** way you can _declare_ variables in go.

    1. using _var_ keyword.

    ```go
    // var <variablename> type = value;
    var greeting String = "Hello World!";
    ```

    > There are lots of _comments_ I am reading that says that we have to specify either `type` or `value` or sometmes **both**.

    2. using `:=` sign.

    ```go
    // <variablename> := value
    greeting := "Devanshu";
    ```

    > The `:=` syntax is shorthand property of _declaration_ and _initializating_ a variable.

    > This syntax is only available **inside a function**. (I really don't know about this.)

  - Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0. `var a int`

  - We can declare mulitple variables.

  ```go
  var p, q, r, s int = 1, 3, 4, 5
  fmt.Println(p, q, r, s)

  var x, y = 6, "Hello"
  n, z := 7, "World!"
  fmt.Println(x, y, n, z)
  ```

  - If the `int`<type> keyword is not specified, you can declare different types of variables in the same line.
  - We can also declare them in _block_ as well.

  ```go
  var (
   a int
   b int = 1
   c string = "hello"
  )
  ```

  - so, there are some **Naming** convention to _follow_ while declaring _variablename_ in go. you can look it upto _w3school.com_.

- **Constant** in _go_.

  - The constant of value must be assigned.
  - _Syntax_ : const <CONSTNAME> type = value
  - They normally written in **Uppercases** as to differenciate between normal variables and constant.
  - _Constant can be written both outside and inside of function_.
  - They are 2 types of constant to say.
    1. Typed constant.
    ```go
    const a int = 3;
    ```
    2. Untyped constant.
    ```go
    const a = "hello";
    ```
  - As you can see, the compiler will decide on it's own for 2nd type. **Pretty much same as variable.**.

- **Data Types** in _go_.

  - Go is statically typed, meaning that once the data types is defined, it can only store that type of data.

- **Arrays** in _go_.

  ```go
  array_name = [length]<data_type>{here goes values} // The length should be pre-defined.
  **OR**
  array_name := [...]<data_type>{here goes values} // Here the length will *inferred* (means that the compiler decides the length of the array, based on the number of values.)
  ```

- **Slices** in _go_.

  - In simple Terms _slices_ are very similar to _Arrays_.
  - In _go_, there are several ways to create slice:

    1. Using `[]datatype{values}` format

    ```go
        mySlice := []int{1, 3, 4}
        // The code above declares a slice of integers of length 3 and also the capacity of 3.
    ```

    - `len()`: returns the length of the slice (the number of elements in the slice)
    - `cap()`: returns the capacity of the slice (the number of elements the slice can grow or shrink to)

    2. Create a slice from _Array_.

    ```go
    arr := [3]int{2, 4, 5}
    mySlice := arr[1:2] // array[start:end]
    ```

    3. Create a slice with make() function.

    - syntax : `slice_name := make([]type, length, capacity)`

    ```go
        myslice1 := make([]int, 5, 10)
    ```

- **Operators** in _go_.

  - There are _Arithmetic_, _assignment_, _comparison_, _logical_, _bitwise_.

- **Conditions** in _go_.

  - Use `if`, condition is _true_
  - Use `else`, condition is _false_
  - Use `else if` to specify a new condition to test, if the first condition is _false_
  - Use `switch` to specify many alternative blocks of code to be executed
  - [if-else statements](https://yourbasic.org/golang/if-else-statement/)
  - [gobyexample : if-else](https://gobyexample.com/if-else)
  - [golang programs: If-else](https://www.golangprograms.com/golang-if-else-statements.html)

- **Switch** in _go_.

  - The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a `break` statement.

- **Loops** in _go_.
  - **`FOR` loop is only available loop in go.**
  - The _for_ loop can take upto _3_ statements.
  - Syntax :
    ```go
        for statement1, statement2, statement3 {
            // code to be executed for each iteration.
        }
    ```
    - Here,
      1. statement1 = Initializes the loop counter value.
      2. statement2 = Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
      3. statement3 = Increases the loop counter value.
  - Note: These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.
  - The *continue statement* is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
  - The *break statement* is used to break out of the interaions in the loop.
  - **Note**: continue and break are usually used with conditions.
  - [Effective Go: For loop](https://go.dev/doc/effective_go#for)
  - [Go by Example: For loop](https://gobyexample.com/for)
  - [5 basic for loop patterns](https://yourbasic.org/golang/for-loop/)