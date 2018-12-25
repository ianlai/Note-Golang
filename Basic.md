### Multiple results
```
func swap(x, y string) (string, string) {
	return y, x
}
func main(){
a, b := swap("hello", "world")
}
```

### Named return values 
```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

### For loop 
```
for i := 0; i < 10; i++ {
    sum += i
}
```

### While loop
```
for sum < 1000 {
	sum += sum
}
```

### If
```
if x < 0 {
	return sqrt(-x) + "i"
}
```

### Switch
```
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
```

### Defer
##### Execute the deferred function until the surrounding function returns.
```
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	for i:=0;i<10000;i++{
		fmt.Println(i)
	}
}
```

### Struct 
```
type Vertex struct {
    X int
    Y int
    Z int
}

func main() {
    v := Vertex{1, 2, 3} //struct initializer 
    fmt.Println(v)

    v.X = 4
    fmt.Println(v.X)

    /* pointer */
    p := &v

    (*p).X = 5
    fmt.Println(v)

    p.X = 1e9     //可以不需要dereferencing   也不需要特別使用箭頭  p->X
    fmt.Println(v)

    var (
        v1 = Vertex{1, 2} 
        v2 = Vertex{X: 1} //struct literals (創建的時候就初始化了) 不用整個物件全部的欄位都使用 也不需要照順序 可以單獨用指定的 ex. X:1 如此一來Y是0
        v3 = Vertex{} // X:0 and Y:0
        p  = &Vertex{1, 2} // has type *Vertex
    )
}

//pointer是不能做運算的 不像C可以用一個位址找到下一個
```

### Array
```
var myint int
var myarr [5]string
myarr[0] = "Hi"
myarr[1] = "there"
fmt.Println(myarr[0], myarr[1])

```

### Slice 
- A slice does not store any data, it just describes a section of an underlying array.

```
names := [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}
fmt.Println(names)

a := names[0:2]  //slice 
b := names[1:3]  //slice
fmt.Println(a, b)

b[0] = "XXX"
fmt.Println(a, b)
fmt.Println(names)
```

- Creating a slice with make
- The make function allocates a zeroed array and returns a slice that refers to that array.

```
a := make([]int, 5)
printSlice("a", a)
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```


### Slice literals
```
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
```


### Range 
```
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {       //range returns (1)index, (2)element
        fmt.Printf("2**%d = %d\n", i, v)
    }   
    for _, value := range pow {   //use underscore to skip the return 
        fmt.Printf("%d\n", value)  
    }   
}
```

### Map 
```
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) //make function returns an initialized map 
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["AAA"] = Vertex{
		11, 22,
	}
	
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["AAA"])
}
```

### Map literals 
```
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
var m = map[string]Vertex{ //type can be omitted 
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

### Function values
- Functions are values too. They can be passed around just like other values.

```
func compute(fn func(int, int) int, a int, b int) int {  //3 arguments(func, int, int) 
	return fn(a, b)
}
func main() {
	square_add := func(x, y int) int {
		return x*x + y*y
	}
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(compute(add, 2, 3))
	fmt.Println(compute(square_add, 2, 3))
}

```

### Function closures
```
func adder() func(int) int {
	sum := 0   //bound to the returned function 
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
output: 
```
0 0
1 -2
3 -6
6 -12
10 -20
```

### Methods
##### A method is a function with a receiver argument.
Form:
```
func (receiver) func_name(argument) return_type 
```
```
type Vertex struct {
	X, Y float64
}

/* Regular function */
func AbsRegular(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* Method with receiver */
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(AbsRegular(v))
	fmt.Println(v.AbsMethod())
}
```
##### We can declare a method on non-struct types, too.
However, we can only declare a method with a receiver whose type is defined in the same package as the method (which means not applicable to int). 
```
type MyFloat float64  //MyFloat is defined in this package so we can declar a method on it 

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
### Pointer Receivers 
##### Methods with pointer receivers can modify the value to which the receiver points
##### With a value receiver, the method operates on a copy of the original value
```
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(100)
	fmt.Println(v.Abs())   //result is 500; result becomes 5 if Scale function's receiver is a value receiver
}
```

### Interface 
##### An interface type is defined as a set of method signatures.
```
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())
	
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```
