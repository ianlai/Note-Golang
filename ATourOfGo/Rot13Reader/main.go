package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(arr []byte) (int, error) {
	p, err := r13.r.Read(arr) //p is the number of byte be read

	//fmt.Println("[Before decode]")
	//s := string(arr[:p])  //change byte array back to string
	//fmt.Print(s)

	for i := 0; i < p; i++ {
		if (arr[i] >= 'A' && arr[i] <= 'M') || (arr[i] >= 'a' && arr[i] <= 'm') {
			arr[i] += 13
		} else if (arr[i] >= 'M') || (arr[i] >= 'm') {
			arr[i] -= 13
		}
		//fmt.Print("p:", p, " len:", len(bytes))
		//fmt.Print("%b", bytes[i])
	}
	//fmt.Println("[After decode]")
	//s = string(arr[:p])  //change byte array back to string
	//fmt.Print(s)

	return p, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") //You cracked the code!
	//s := strings.NewReader("Lbh")
	r := rot13Reader{s}
	fmt.Println("[Result]")
	io.Copy(os.Stdout, &r)
}
