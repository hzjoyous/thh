package example

import (
	"context"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/go-playground/validator/v10"
	"math"
	"runtime"
	"strconv"
	"thh/helpers"
	"time"
)

func init() {
	addConsole("practiceArrAndMap", "", practiceArrAndMap)
	addConsole("bitCompute", "", bitCompute)
	addConsole("algorithmEatFood", "", algorithmEatFood)
	addConsole("practiceChan", "", practiceChan)
	addConsole("practiceDefer", "", practiceDefer)
	addConsole("practiceContext", "", practiceContext)
	addConsole("practiceGo", "", practiceGo)
	addConsole("practiceFuncCase1", "", practiceFuncCase1)
	addConsole("practiceValidate", "", practiceValidate)
	addConsole("practiceGG", "", practiceGG)

}

func practiceDefer() {
	practiceDeferFunc1 := func() {
		practiceDeferFunc2 := func() {
			defer func() {
				fmt.Println("this is practiceDeferFunc2 defer")
			}()
			fmt.Println("this is practiceDeferFunc2")
		}

		defer func() {
			fmt.Println("this is practiceDeferFunc1 defer")
		}()
		fmt.Println("this is practiceDeferFunc1")
		practiceDeferFunc2()
	}

	defer func() { fmt.Println("a") }()
	defer func() { fmt.Println("b") }()
	defer func() { fmt.Println("c") }()
	practiceDeferFunc1()

	for {
		time.Sleep(time.Second)
	}
}

func bitCompute() {

	// 10进制数2进制打印
	var ii int64 = -5
	fmt.Printf("%b\n", ii)

	var i int64 = 5
	fmt.Printf("%b\n", i)

	// 8进制
	var a int64 = 011
	fmt.Println("a=", a)

	// 16进制
	var j = 0x11
	fmt.Println("j=", j)

}

func practiceArrAndMap() {
	var (
		intArr  [2]int
		intArr2 [2][2]int
		intArr3 []int
		mapArr  map[int]string
	)

	fmt.Println(intArr)
	fmt.Println(intArr2)
	fmt.Println(mapArr)

	if mapArr == nil {
		fmt.Println("it is nil")
	} else {
		fmt.Println("it not not nil")
	}

	fmt.Println("mapArr len is", len(mapArr))

	mapArr = make(map[int]string, 0) // 必要，否则panic
	fmt.Println(mapArr)
	if mapArr == nil {
		fmt.Println("it is nil")
	} else {
		fmt.Println("it not not nil")
	}

	fmt.Println("map len4 is", len(mapArr))

	mapArr[1] = "string"

	fmt.Println(mapArr)

	fmt.Println("综上表现，goMap的初始化的设计有些鸡肋")

	practiceArrAndMapPrintSlice := func(x []int) {
		fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	}
	practiceArrAndMapPrintSlice(intArr3)
	fmt.Println(len(intArr3), cap(intArr3))
	intArr3 = append(intArr3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(intArr3), cap(intArr3))
	practiceArrAndMapPrintSlice(intArr3)
	intArr3[0] = 9
	practiceArrAndMapPrintSlice(intArr3)
}

func algorithmEatFood() {

	fmt.Println("盘子")

	plate := make(chan int)
	apple := make(chan int)
	orange := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			dadPlate := <-plate
			fmt.Println("盘子可以放入一个水果 dad Plate", dadPlate)
			apple <- 1
			fmt.Println("dad放入一个苹果")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			momPlate := <-plate
			fmt.Println("盘子可以放入一个水果mom Plate", momPlate)
			orange <- 1
			fmt.Println("mom放入一个橘子")
		}

	}()

	go func() {
		for {
			time.Sleep(time.Second)
			sonOrange := <-orange
			fmt.Println("儿子取走一个橘子 sonOrange Plate", sonOrange)
			plate <- 1
			fmt.Println("儿子清空盘子")
		}

	}()
	go func() {
		for {
			time.Sleep(time.Second)
			daughterApple := <-apple
			fmt.Println("女儿取走一个苹果 daughterApple Plate", daughterApple)
			plate <- 1
			fmt.Println("女儿清空盘子")
		}
	}()
	plate <- 1
	for {
		time.Sleep(time.Second)
	}
}

func practiceChan() {
	chanVar := make(chan int)
	go func() {
		for {
			varDad := <-chanVar
			fmt.Println("dad", varDad)
		}
	}()
	go func() {
		for {
			varMom := <-chanVar
			fmt.Println("mom", varMom)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second)
			chanVar <- i
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}

func practiceContext() {
	defer func() {
		fmt.Println("销毁了")
	}()
	fmt.Println("当前go数量")
	fmt.Println(runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响
	gen := func() <-chan int {
		ch := make(chan int)
		go func() {
			var n int
			for {
				ch <- n
				n++
				time.Sleep(time.Second)
				fmt.Println(n)
			}
		}()
		return ch
	}
	for n := range gen() {
		fmt.Println("print n from for:", n)
		fmt.Println("当前go数量:", runtime.NumGoroutine())
		if n == 5 {
			break
		}
	}

	for i := 0; i < 5; i++ {
		go func(i int) {
			counter := 1
			for {
				counter += 1
				fmt.Println(i, counter)
				time.Sleep(time.Second * 1)
			}
		}(i)

		fmt.Println("当前go数量:", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 5)

	defer cancel() // 避免其他地方忘记 cancel，且重复调用不影响
	genCanQuite := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			var n int
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- n:
					n++
					time.Sleep(time.Second)
				}
			}
		}()
		return ch
	}

	for n := range genCanQuite(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

}

func practiceGo() {
	go func() {
		fmt.Println("A:我是一个来自闭包的携程，end")
	}()
	numData1 := make(chan int)
	numData2 := make(chan int)

	goB := func() {
		fmt.Println("B:我是一个来自函数的携程,end")
	}

	goC := func(numData1 chan int) {
		n := 0
		for {
			n++
			numData1 <- n
			fmt.Println("C:say2:我是一个来自函数的携程,我向通道中发出了一条数据：" + strconv.Itoa(n))
			time.Sleep(1 * time.Second)
		}
	}

	goD := func(numData2 chan int) {
		for {
			fmt.Println("我是一个会sleep的go，看看会不会阻塞，结果是没有阻塞")
			time.Sleep(time.Second * 3)
			resultNumData2 := <-numData2
			fmt.Println("D:say2:我是一个来自函数的携程，接收到通道内数据：" + strconv.Itoa(resultNumData2))
		}
	}

	goE := func() {
		for {
			fmt.Println("E:我是一个来自函数的携程，我只会这一句")
			time.Sleep(1 * time.Millisecond * 500)
		}
	}

	goF := func() {
		fmt.Println("F:我是一个来自函数的携程,我预感我将不会执行完毕end")
	}

	go goB()
	go goC(numData1)
	go goD(numData2)
	go goE()

	for i := 1; i <= 10; i++ {
		resultNumData1 := <-numData1
		numData2 <- resultNumData1
		fmt.Println("Main:我是主程,我接受到了来自C程序的信息" + strconv.Itoa(resultNumData1))
	}
	go goF()
}

func practiceFuncCase1() {
	fmt.Println("多值返回")
	a, b := func() (a string, b int) {
		a = "str"
		b = 1
		return
	}()
	fmt.Println(a, b)

	startTime := helpers.GetMicroTime()
	cmdFuncCounter = 1
	fmt.Println(fibonacci(46))
	endTime := helpers.GetMicroTime()
	fmt.Println(float64(endTime-startTime) / 1000.0)
	fmt.Println("计数器统计方法执行次数", cmdFuncCounter)

	r := fibonacciTailCall(1000, 0, 1)
	fmt.Println(r)
	// 求第46位斐波那契数列将调用方法5942430146次
}

var (
	cmdFuncCounter int
)

func fibonacci(n int) (res int) {
	cmdFuncCounter++
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 尾递归
func fibonacciTailCall(n int64, result int64, preValue int64) int64 {
	if n == 0 {
		return result
	}
	return fibonacciTailCall(n-1, preValue, result+preValue)
}

func practiceValidate() {
	var validate *validator.Validate
	validateVariable := func() {
		myEmail := "abc@tencent.com"
		errs := validate.Var(myEmail, "required,email")
		if errs != nil {
			fmt.Println(errs)
			return
			//停止执行
		}
		// 验证通过，继续执行
		fmt.Println("end")
	}
	validate = validator.New()
	validateVariable()
}

func practiceGG() {
	type Point struct {
		X, Y float64
	}

	Polygon := func(n int, x, y, r float64) []Point {
		result := make([]Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
		}
		return result
	}

	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	err := dc.SavePNG("out.png")
	if err != nil {
		fmt.Println(err)
	}
}
