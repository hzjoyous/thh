package example

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func init() {
	addConsole("suggestStrOperation", "suggestStrOperation", suggestStrOperation)
	addConsole("suggestFuncParamsOperation", "suggestFuncParamsOperation", suggestFuncParamsOperation)
}

func suggestFuncParamsOperation() {

	type name struct {
		c func()
	}
	c1 := func() {
		fmt.Println("c1")
	}

	c2 := func() {
		fmt.Println("c2")
	}

	c3 := func() {
		fmt.Println("c3")
	}
	nameList := make(map[string]name, 0)
	nameList["c1"] = name{c: c1}
	nameList["c2"] = name{c: c2}
	nameList["c3"] = name{c: c3}

	funcList := make(map[string]func())
	fmt.Println("这是一段可能会导致bug的代码")
	for key, value := range nameList {
		fmt.Println("bug的起因是下面的闭包函数出现了一个静态的变量，" +
			"虽然value每次都是新的，" +
			"但是func中保存了第一个，" +
			"如果想要达到预期的目的，" +
			"可以创建一个高阶函数")
		funcList[key] = func() {
			value.c()
		}

		fmt.Println("与上面的闭包不同，" +
			"上面的value.c()是一个直接执行的函数，" +
			"value又是一个变量但是这个变量依赖于循环。" +
			"所以等到要执行funcList[key]()的时候会无法确定value（个人感觉是可以实现的，但是编译器目前没有实现）" +
			"但是下面的代码在参数部分明确了闭包执行的内容，所以可以规避，个人感觉如果以后支持上面的写法，运行逻辑其实同下面的逻辑")
		_ = func(value func()) func() {
			return func() {
				value()
			}
		}(value.c)
	}
}

func suggestStrOperation() {
	fmt.Println("this is a suggestStrOperation")
	k := 5
	d := [5]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(30000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			// 每次声明一个bytes.buffer 仍然比正常的字符串拼接快很多
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

	}

	if index == 4 { // for stable bytes.Buffer
		s = buf.String()
	}
	fmt.Println(len(s)) // consume s to avoid compiler optimization
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}
