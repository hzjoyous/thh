package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
	"thh/arms"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "yzw_json_tool",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runYzwJsonTool,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

type vSort []v

func (s vSort) Len() int      { return len(s) }
func (s vSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less
// 排序原则
// 完成现阶段目标的当置于中间位置
// 不期望的当滞后
// 离目标越接近的越靠前
// 离目标同样接近的期望越高越高前
func (s vSort) Less(i, j int) bool {
	return s[i].City < s[j].City
}

func runYzwJsonTool(cmd *cobra.Command, args []string) {
	vJson := arms.StorageGet("final.json")
	vList := arms.JsonDecode[vSort](vJson)

	newVList := vSort{}

	arms.StoragePut("t.csv", "城市,博士点,招生单位,院系,专业,方向,学习方式,招生人数,考试方式,政治,英语,专业课1,专业科2,信息源\n", false)
	fmt.Println(len(vList))
	sort.Stable(vList)
	hMap := map[string]v{}
	for _, vItem := range vList {
		if _, ok := hMap[vItem.Name]; ok {
			continue
		}
		hMap[vItem.Name] = vItem
		newVList = append(newVList, vItem)
		for _, vtItem := range vItem.SpecialityList {
			for _, sItem := range vtItem.Info {
				t := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,研招网2022信息\n",
					replaceDot(vItem.City),      // 城市
					replaceDot(vtItem.IsDoctor), // 博士点
					replaceDot(sItem.Zsdw),      // 招生单位
					replaceDot(sItem.Yx),        // 院系
					replaceDot(sItem.Zy),        // 专业
					replaceDot(sItem.Fx),        // 方向
					replaceDot(sItem.StudyFunc), // 学习方式
					replaceDot(sItem.Number),    // 招生人数
					replaceDot(sItem.Kxfs),      // 考试方式
					replaceDot(sItem.Zz),        // 政治
					replaceDot(sItem.Wy),        // 英语
					replaceDot(sItem.B1),        // 专业课1
					replaceDot(sItem.B2),        // 专业科2
				)
				arms.StoragePut("t.csv", t, true)
			}
		}
	}
	vListString := arms.JsonEncode(newVList)
	arms.StoragePut("final2.json", vListString, false)
}

func replaceDot(s string) string {
	return strings.NewReplacer("\n", "", ",", "_", "，", "_").Replace(s)
	//return strings.ReplaceAll(s, ",", "")
}
