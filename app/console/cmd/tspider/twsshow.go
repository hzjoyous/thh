package tspider

import (
	"fmt"
	"github.com/spf13/cobra"
	"path"
	"sort"
	"thh/arms"
	"thh/arms/config"
)

func init() {
	appendCommand(CmdTwsshow)
}

var CmdTwsshow = &cobra.Command{
	Use:   "twsshow",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runTwsshow,
	//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

type vTwSlice []vTw

func (s vTwSlice) Len() int      { return len(s) }
func (s vTwSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less
// 排序原则
// 完成现阶段目标的当置于中间位置
// 不期望的当滞后
// 离目标越接近的越靠前
// 离目标同样接近的期望越高越高前
func (s vTwSlice) Less(i, j int) bool {
	return len(s[i].VerList) > len(s[j].VerList)
}

func runTwsshow(cmd *cobra.Command, args []string) {
	jsonStr, _ := arms.FileGetContents("./storage/tmp/" + config.GetString("t_Statistics_filename"))
	baseName := "./storage/tmp/" + path.Base(config.GetString("t_Statistics_filename"))
	counter := map[string]int{}
	arms.FilePutContents(baseName+".txt", []byte(""), false)
	tList := arms.JsonDecodeBl[vTwSlice](jsonStr)
	sort.Sort(tList)
	for _, vTwItem := range tList {
		arms.FilePutContents(baseName+".txt", []byte(fmt.Sprintln(vTwItem.Name, ":", "https://twitter.com/"+vTwItem.ScreenName)), true)
		if vTwItem.Desc != "" {
			arms.FilePutContents(baseName+".txt", []byte(fmt.Sprintln("desc:", vTwItem.Desc)), true)
		}
		for verName, verLink := range vTwItem.VerList {
			arms.FilePutContents(baseName+".txt", []byte(fmt.Sprintln(verName, "_转发地址:", verLink)), true)
			if count, ok := counter[verName]; ok {
				counter[verName] = count + 1
			} else {
				counter[verName] = 1
			}
		}
		arms.FilePutContents(baseName+".txt", []byte("##################################\n"), true)
	}
	fmt.Println(counter)
}
