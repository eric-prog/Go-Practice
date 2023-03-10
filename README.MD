# GO Practice & API Tutorial

## Tutorial
<br>

### Naming Rules
Camel Case <br>
Pascal Case <br>
Snake Case <br>

### Variables
var a, b, c, d int = 1, 3, 5, 7 <br> <br>
c, d := 7, "World!" <br> <br>
const CONSTNAME type = value <br> <br>
arr2 := [...]int{4,5,6,7,8} array with inferred length <-- array NOT slice<br> <br>
slice can grow and shrink as you see fit --> []datatype{values} format <br> 
var a [5]int --> initialized with a length of 5 and filled with zeros <br> <br>
Maps in go can hold key value pairs and similar to dictionaries in Python. map[type_of_key]type_of_value --> Ex: vertices := make(map[string]int) <-- make is a built-in function --> adding to map --> vertices["triangle"] = 3 <br> <br>

### Loops

The only type of loop in Go is for loop (so no while loops). <br>
```
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
``` 
<br>
For loop can also act as while loop

```
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```
<br>
Looping through an array

```
arr := []string{"1", "2", "3"}
for index, value := range arr {
    fmt.Println("index", index, "value", value)
}
```
<br>
Looping through a map

```
m := make(map[string]int)
m["a"] = 2

for key, value := range arr {
    fmt.Println("key", key, "value", value)
}
``` 
<br>

### Functions

```
func sum (x int, y int) int {
    return x + y
}
```
<br>

Functions can multiple return types <br>

```
function sqrt(x float64) (float64, error) { // returns float64 and error
    if x < 0 {
        return 0, error.New("Undefined for negative numbers") // return 0, and error
    }

    return math.Sqrt(x), nil // still need to return 2 values so return nil not error
}

```
<br>

### Structs 

```
type person struct {
    name string
    age int
}

p := person{name: "bob", age: 12}
```
<br>

### Pointers

```
func main () {
    i := 7
    fmt.Println(&i) // outputs memory address
}
```
<br>

New example: <br>
```
func main () {
    i := 7
    inx(i) // <-- doing this will just copy i and i will not be updated
    inx(&i) // <-- pass in memory address
    fmt.Println(&i) // outputs memory address
}

// NEW we pass in pointer to variable and the x gets updated
func inc (x *int) { // <-- * gets pointer
    *x++ // <-- * infront actually gets value in address and updates that value
}

// OLD that will not update because we update the copy
func inc (x int) {
    x++
}
```
<br>

### API dependencies

```
go mod init example/Go-Api-Tutorial 
go get github.com/gin-gonic/gin
```