package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"sort"
	"strings"
	"thh/arms"
	"thh/arms/str"
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

func runTmp(cmd *cobra.Command, args []string) {
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
