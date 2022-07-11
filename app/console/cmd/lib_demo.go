package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"math/big"
	"regexp"
	"sort"
	"strings"
	"sync"
	"thh/arms"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "p:TimeFormat", Short: "", Run: practiceTimeFormat})
	appendCommand(&cobra.Command{Use: "p:BigNumber", Short: "", Run: practiceBigNumber})
	appendCommand(&cobra.Command{Use: "p:Panic", Short: "", Run: practicePanic})
	appendCommand(&cobra.Command{Use: "p:SyncOnce", Short: "", Run: practiceSyncOnce})
	appendCommand(&cobra.Command{Use: "p:Reg", Short: "", Run: practiceReg})
	appendCommand(&cobra.Command{Use: "eolTest", Short: "", Run: eolTest})
	appendCommand(&cobra.Command{Use: "suggestStrOperation", Short: "suggestStrOperation", Run: suggestStrOperation})
	appendCommand(&cobra.Command{Use: "suggestFuncParamsOperation", Short: "suggestFuncParamsOperation", Run: suggestFuncParamsOperation})
	appendCommand(&cobra.Command{Use: "p:ArrAndMap", Short: "", Run: practiceArrAndMap})
}

func eolTest(cmd *cobra.Command, args []string) {

	peopleList := make(map[int]*people, 2000)
	peopleWaitList := make(map[int]people, 2000)

	for i := 1; i <= 2000; i++ {
		weight := int(i / 10 / 2)
		if i > 1900 {
			weight += 100
		}
		peopleList[i] = &people{Id: i, Weight: weight}
	}

	wg4team := &sync.Mutex{}
	//wg4pool := &sync.Mutex{}

	go func() {
		for {
			// 等待池子进入people
			wg4team.Lock()

			if len(peopleWaitList) < 10 {
				// 如果排队池中人数小于10任务结束，放弃当前任务
				wg4team.Unlock()
				time.Sleep(time.Microsecond * 10)
				continue
			}

			teamA := make(map[int]people, 5)
			teamB := make(map[int]people, 5)

			// 池子中抽取人
			//delete(peopleWaitList, )

			tl := make(tmpList, 0)
			for _, peopleEntity := range peopleWaitList {
				tl = append(tl, peopleEntity)
			}
			sort.Sort(tl)

			waitStartPeople := tl[0:10]

			groupI := 1
			for _, peopleEntity := range waitStartPeople {
				groupI += 1
				if groupI%2 == 1 {
					teamA[peopleEntity.Id] = peopleEntity
				} else {
					teamB[peopleEntity.Id] = peopleEntity
				}
				delete(peopleWaitList, peopleEntity.Id)
			}

			aRate := 1
			bRate := 1
			for _, peopleEntity := range teamA {
				aRate += peopleEntity.Weight
			}
			for _, peopleEntity := range teamB {
				bRate += peopleEntity.Weight
			}
			victoryTeam := teamA
			failTeam := teamB
			if arms.RandTrue(bRate, aRate+bRate) {
				victoryTeam = teamB
				failTeam = teamA

			}
			// 胜利的队伍 比赛场次 +1 连胜 场次 +1
			for _, peopleEntity := range victoryTeam {
				peopleList[peopleEntity.Id].Total += 1
				peopleList[peopleEntity.Id].Victory += 1
				peopleList[peopleEntity.Id].Score += 1
				peopleList[peopleEntity.Id].Continuity += 1
			}

			for _, peopleEntity := range failTeam {
				peopleList[peopleEntity.Id].Total += 1
				peopleList[peopleEntity.Id].Fail += 1
				peopleList[peopleEntity.Id].Score -= 1
				peopleList[peopleEntity.Id].Continuity = 0
			}

			//dataByte, _ := json.Marshal(peopleList)
			//fmt.Println(string(dataByte))

			//dataByte, _ := json.Marshal(*peopleList[1])
			//fmt.Println(string(dataByte))
			//dataByte, _ = json.Marshal(*peopleList[1000])
			//fmt.Println(string(dataByte))
			//dataByte, _ = json.Marshal(*peopleList[2000])
			//fmt.Println(string(dataByte))
			fmt.Println(
				peopleList[2000].Weight,
				fmt.Sprintf("%.2f", cast.ToFloat64(peopleList[2000].Victory)/cast.ToFloat64(peopleList[2000].Total)),
				peopleList[2000].Score,
				"#\t",
				peopleList[1500].Weight,
				fmt.Sprintf("%.2f", cast.ToFloat64(peopleList[1500].Victory)/cast.ToFloat64(peopleList[1500].Total)),
				peopleList[1500].Score,
				"#\t",
				peopleList[1000].Weight,
				fmt.Sprintf("%.2f", cast.ToFloat64(peopleList[1000].Victory)/cast.ToFloat64(peopleList[1000].Total)),
				peopleList[1000].Score,
				"#\t",
				peopleList[300].Weight,
				fmt.Sprintf("%.2f", cast.ToFloat64(peopleList[300].Victory)/cast.ToFloat64(peopleList[300].Total)),
				peopleList[300].Score,
				"#\t",
				peopleList[290].Weight,
				fmt.Sprintf("%.2f", cast.ToFloat64(peopleList[290].Victory)/cast.ToFloat64(peopleList[290].Total)),
				peopleList[290].Score,
			)
			wg4team.Unlock()
		}
	}()

	for {

		for i := 1; i <= 2000; i++ {
			// 等待队伍选人完成
			wg4team.Lock()
			// 1/2 的可能打算玩
			if arms.RandTrue(1, 2) {
				peopleWaitList[i] = *peopleList[i]
			}
			wg4team.Unlock()

		}
	}

}

type people struct {
	Id         int `json:"Id"`
	Weight     int `json:"Weight"` // 权重
	Victory    int `json:"Victory"`
	Fail       int `json:"Fail"`
	Total      int `json:"Total"`
	Continuity int `json:"Continuity"`
	WaitTime   int `json:"waitTime"`
	Score      int `json:"Score"`
}
type tmpList []people

func (p tmpList) Less(i, j int) bool {
	return (p[i].Score + p[i].Continuity) < (p[j].Score + p[i].Continuity)
}
func (p tmpList) Len() int {
	return len(p)
}
func (p tmpList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func practiceTimeFormat(cmd *cobra.Command, args []string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("01/02 03:04:05PM 2006 -0700"))
	fmt.Println(time.Now().Hour())
}

func practiceBigNumber(cmd *cobra.Command, args []string) {
	x := big.NewInt(1)
	y := big.NewInt(1)
	for i := 1; i <= 1000; i++ {
		x = big.NewInt(0).Add(x, y)
		y = x
	}
	fmt.Println(y)
}
func practicePanic(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("运行结束", err)
		}
	}()

	myPanic := func() {
		var x = 30
		var y = 0
		for i := 0; i < 5; i++ {

			//if i == 3 {
			//	panic("我就是一个大错误！")
			//}

			c := x / y
			fmt.Println(c)
		}
	}
	myPanic()
	fmt.Printf("end！")
}

func practiceSyncOnce(cmd *cobra.Command, args []string) {
	type runOne struct {
		once sync.Once
	}
	run := new(runOne)
	run.once.Do(func() {
		fmt.Println("第一次")
	})
	run.once.Do(func() {
		fmt.Println("第一次")
	})

}

func practiceReg(cmd *cobra.Command, args []string) {
	regUnit := func(regStr string, matchStr string, unMatchStr string) {
		defer func() {
			fmt.Println("Enter defer function.")
			if p := recover(); p != nil {
				fmt.Printf("panic: %s\n", p)
			}
			fmt.Println("Exit defer function.")
		}()
		reg := regexp.MustCompile(regStr)
		//根据规则提取关键信息
		result1 := reg.FindAllStringSubmatch(matchStr, -1)
		fmt.Println(regStr, matchStr, "match", len(result1))
		//根据规则提取关键信息
		result2 := reg.FindAllStringSubmatch(unMatchStr, -1)
		fmt.Println(regStr, unMatchStr, "match", len(result2))

	}
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`a[0-9]c`)
	if reg == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)

	fmt.Println("reg start")

	// 结尾匹配
	regUnit(`^[abcdef]*$`, `accede`, `beam`)

	// 结尾匹配
	regUnit(`[a-z]k$`, `Mick`, `Nickneven`)

	// 单词结尾
	regUnit(`fu\b`, `tofu`, `futz`)

	// ()重复出现 就是 匹配 （$字符串1）匹配 $字符串1
	// allochirally 可以匹配到 (all) ochir all 前面的all在后面再次出现了
	// go 目前不支持\1
	//regUnit(`(...).*\1`,`allochirally`,`anticker`)

	// go 目前不支持?!
	regUnit(`^(?!.*(.)(.)\2\1.*)`, `acritan`, `anallagmatic`)
	//
	//regUnit(``,``,``)
	//
	//regUnit(``,``,``)

}

func suggestFuncParamsOperation(cmd *cobra.Command, args []string) {

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

func suggestStrOperation(cmd *cobra.Command, args []string) {
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
