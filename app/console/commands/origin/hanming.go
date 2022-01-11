package origin

import (
	"fmt"
	"thh/base"
	"thh/helpers"
)

func init() {
	c := base.Console{Signature: "o:hanming", Description: "this is a reflectDemo", Handle: hanming}
	commandList[c.Signature] = c
}

func hanming() {
	var data [4]int
	data[0] = 1
	data[1] = 0
	data[2] = 1
	data[3] = 0

	var check [3]int
	check[0] = 0
	check[1] = 0
	check[2] = 0

	var haimingCode [7]int
	haimingCode[0] = 0
	haimingCode[1] = 0
	haimingCode[2] = 0
	haimingCode[3] = 0
	haimingCode[4] = 0
	haimingCode[5] = 1
	haimingCode[6] = 1
	start := 1

	for {
		fmt.Println(start)
		t := 0
		for i := 1; i <= 7; i++ {
			fmt.Println(i & start)
			if i&start >= 1 {
				t = t ^ haimingCode[i-1]
			}
		}
		fmt.Println(t)
		haimingCode[start-1] = t
		start = start << 1
		if start > len(haimingCode) {
			break
		}
	}

	fmt.Println(haimingCode)

	w1 := func(a int, b int) int {
		return helpers.ToInt(!(a == b))
	}
	w2 := func(a int, b int) int {
		return helpers.ToInt(a != b)
	}
	fmt.Print(w1(1, 1))
	fmt.Print(w1(1, 0))
	fmt.Print(w1(0, 0))
	fmt.Println(w1(0, 1))

	fmt.Print(w2(1, 1))
	fmt.Print(w2(1, 0))
	fmt.Print(w2(0, 0))
	fmt.Println(w2(0, 1))

	fmt.Print(1 ^ 1)
	fmt.Print(1 ^ 0)
	fmt.Print(0 ^ 0)
	fmt.Println(0 ^ 1)
}
