package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"reflect"
	"sort"
	"strings"
	"sync"
	"thh/arms"
	"thh/arms/app"
	"thh/arms/str"
	"time"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tmp",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runTmp,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

type SortTestStruct struct {
	id int
}

type String4Pl string

func (itself String4Pl) String() string {
	return "this is " + string(itself)
}

type buildYml struct {
	Name  string
	Age   int
	Other struct {
		Ok   string
		Name string
	}
}

func runTmp(cmd *cobra.Command, args []string) {

	listArr := []int{3, 4, 5, 6, 7, 2, 3}
	a4list := sync.WaitGroup{}
	for _, v := range listArr {
		a4list.Add(1)
		go func(v int) {
			defer a4list.Done()
			time.Sleep(time.Millisecond * cast.ToDuration(v))
			fmt.Println(v)
		}(v)
	}
	a4list.Wait()
	return

	b := buildYml{}
	//v.WriteConfig()
	aa, _ := yaml.Marshal(b)
	fmt.Println(cast.ToString(aa))

	var t time.Duration
	t = 60_000_000_000
	fmt.Println(t)

	fmt.Println(app.GetRunTime())

	var s String4Pl
	s = "123141"
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	return
	var err1 error
	if err1 == nil {
		fmt.Println("刚才修改的没问题")
		fmt.Println(err1)
	}

	tmpPassWord := "asdasdas"
	password := arms.MakePassword(tmpPassWord)
	err := arms.VerifyPassword(password, tmpPassWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("密码正确")
	}
	err = arms.VerifyPassword(password, "tmpPassWord")
	if err != nil {
		fmt.Println(err)
	}

	return
	tmpSplice := []SortTestStruct{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}
	sort.Slice(tmpSplice, func(i, j int) bool {
		return tmpSplice[i].id < tmpSplice[j].id
	})

	var Pln = fmt.Println
	Pln(tmpSplice)

	fmt.Println(str.Camel(strings.ToLower(`ISSUE_PREPAID_DISTRIBUTOR_COMMISSION_FUND`)))
	// strings
	fmt.Println(strings.TrimSpace(" kong kong kong \t\n\r\n"))

	fmt.Println(strings.Join([]string{"1", "2"}, "-"))
	fmt.Println(cast.ToBool(1))

	println()
	var a interface{}
	setT(1)
	a = getT()
	fmt.Println(a)
	setT("aaa")
	a = getT()
	fmt.Println(a)

}

var t interface{}

func setT(vatT interface{}) {
	t = vatT
}
func getT() interface{} {
	return t
}
