package main

import (
	"fmt"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number or 0: %f", e) 
}

func Sqrt(x float64) (float64, error) {
	/* input less than 0 */
	if x<=0{
		err := ErrNegativeSqrt(x)
		return x, err
	}
	/* input larger than 0 */
	z := 1.0
	count := 1 
	for ;count<10;count++{
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}

func main() {
	fmt.Println("[Errors]")
	
	max   := 30
	start := float64(-max/2)
	/* Test */ 
	for i:=0; i<max; i++{
		cur := float64(start) + float64(i)
		ans, err := Sqrt(cur)
		if err==nil{
			fmt.Println(cur, ":  sqrt(", cur, ")=>", ans)
		}else{ 
			fmt.Println(cur, ":  ", err)
		}
	}
	//fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))
}
