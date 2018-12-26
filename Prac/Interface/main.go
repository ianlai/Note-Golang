package main

import(
    "fmt"
    "math"
)
type ICalc interface {
	size() int
}
type Vector struct {
	x, y int
}
type String struct{
    str string
}
func (s String) size() int {
    return len(s.str)
}
func (v Vector) size() int {
    return int(math.Sqrt(float64(v.x*v.x+v.y*v.y)))
}
func main() {
    var str1 ICalc = String{"Hi"}   //ok
    var str2 = String{"Taiwan"}     //ok 
    fmt.Println(str1.size())
    fmt.Println(str2.size())

    var v1 ICalc = Vector{3,4}      //ok
    var v2 = Vector{y:12, x:5}      //ok 
    fmt.Println(v1.size())
    fmt.Println(v2.size())
}
