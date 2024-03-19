- For starters : Declare _main_ package in same directory.

- Then import _fmt_ package as well, which mainly contains all the pre-written function that will be used in go language. :smile

- Running go file
```
$ go run .
```

- In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

- **Comments** in *go*: 
    - Single line comments starts with `//`.
    - multline comments `/*` and `*/`.

- Sometimes we need to **build** our files. 
    - `go build <flename>` command is used for it.
    - And you can run it via using `./<filename>`. Yup, just like that. :)
-[Gobyexample](https://gobyexample.com/)

- **Variables** in *go*.
    - There are **2** way you can *declare* variables in go.
        1. using *var* keyword.
        ```go
        // var <variablename> type = value;
        var greeting String = "Hello World!";
        ```
        > There are lots of *comments* I am reading that says that we have to specify either `type` or `value` or sometmes **both**.

        2. using `:=` sign.
        ```go
        // <variablename> := value
        greeting := "Devanshu";
        ```
        > The `:=` syntax is shorthand property of *declaration* and *initializating* a variable.

        >  This syntax is only available **inside a function**. (I really don't know about this.)

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
    - We can also declare them in *block* as well.
    ```go
    var (
     a int
     b int = 1
     c string = "hello"
   )
   ```
   - so, there are some **Naming** convention to *follow* while declaring *variablename* in go. you can look it upto *w3school.com*. 


- **Constant** in *go*.
    - The constant of value must be assigned.
    - *Syntax* : const <CONSTNAME> type = value
    - They normally written in **Uppercases** as to differenciate between normal variables and constant.
    - *Constant can be written both outside and inside of function*.
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

- **Data Types** in *go*.
    - Go is statically typed, meaning that once the data types is defined, it can only store that type of data.

- **Arrays** in *go*.
```go
array_name = [length]<data_type>{here goes values} // The length should be pre-defined.
**OR**
array_name := [...]<data_type>{here goes values} // Here the length will *inferred* (means that the compiler decides the length of the array, based on the number of values.)
```


- **Slices** in *go*.
    - In simple Terms *slices* are very similar to *Arrays*.
    - In *go*, there are several ways to create slice:
        1. Using `[]datatype{values}` format
        ```go
            mySlice := []int{1, 3, 4} 
            // The code above declares a slice of integers of length 3 and also the capacity of 3.
        ```
        - `len()`: returns the length of the slice (the number of elements in the slice)
        - `cap()`: returns the capacity of the slice (the number of elements the slice can grow or shrink to)

        2. Create a slice from *Array*.
        ```go
        arr := [3]int{2, 4, 5}
        mySlice := arr[1:2] // array[start:end]
        ```

        3. Create a slice with make() function.
        - syntax : `slice_name := make([]type, length, capacity)`
        ```go
            myslice1 := make([]int, 5, 10)
        ```

- **Operators** in *go*.
    - There are *Arithmetic*, *assignment*, *comparison*, *logical*, *bitwise*.

- **Conditions** in *go*.
    - Use `if`, condition is *true*
    - Use `else`, condition is *false*
    - Use `else if` to specify a new condition to test, if the first condition is *false*
    - Use `switch` to specify many alternative blocks of code to be executed