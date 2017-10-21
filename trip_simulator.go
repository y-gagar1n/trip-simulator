package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	n := 10
	k := 1
	j := 7

	var a []int32 = make([]int32, n)
	var c int32 = 0
	for {
		c++
		fmt.Print(c)
		fmt.Print(": ")
		for i := 0; i < n; i++ {
			step := r.Int31n(int32(j+1-k) + int32(k))
			fmt.Print(step)
			fmt.Print(" ")
			if i > 0 {
				step = int32(math.Min(float64(a[i-1]), float64(step)))
				a[i-1] -= step
			}
			if i < n {
				a[i] += step
			}
		}
		fmt.Println()
		fmt.Println(a)
		fmt.Print("|")
		for i := 0; i < n; i++ {
			for q := int32(0); q < a[i]; q++ {
				fmt.Print(" ")
			}
			if i < n-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		fmt.Println(float64(a[n-1]) / float64(c))
		fmt.Println()
		_, _, _ = keyboard.GetKey()

	}

	fmt.Println((j + k) / 2)
	fmt.Println(a)

}
