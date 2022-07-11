package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"sort"
	"thh/arms"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "y",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runY,
	})
}

type ysPeople struct {
	Name       string  `json:"name"`
	Sort       int     `json:"sort"`
	Level      int     `json:"level"`
	Expect     int     `json:"expect"`
	Skill      ysSkill `json:"skill"`
	ArmLevel   int     `json:"armLevel"`
	Equipments ysEQ    `json:"equipments"`
}

type ysSkill struct {
	One     int `json:"one"`
	OneEx   int `json:"oneEx"`
	Two     int `json:"two"`
	TwoEx   int `json:"twoEx"`
	Three   int `json:"three"`
	ThreeEx int `json:"threeEx"`
}

func (itself ysSkill) Ex() int {
	return -(itself.One - itself.OneEx + itself.Two - itself.TwoEx + itself.Three - itself.ThreeEx)
}

type ysEQ struct {
	Flower      int `json:"flower"`
	FlowerEx    int `json:"flowerEx"`
	Head        int `json:"head"`
	HeadEx      int `json:"headEx"`
	Necklace    int `json:"necklace"`
	NecklaceEx  int `json:"necklaceEx"`
	Hourglass   int `json:"hourglass"`
	HourglassEx int `json:"hourglassEx"`
	Plume       int `json:"plume"`
	PlumeEx     int `json:"plumeEx"`
}

func (itself ysEQ) Ex() int {
	return -(itself.Flower - itself.FlowerEx + itself.Head - itself.HeadEx + itself.Necklace - itself.NecklaceEx + itself.Hourglass - itself.HourglassEx + itself.Plume - itself.PlumeEx)
}

func (itself ysPeople) weight() int {
	lEx := itself.Level - 0
	skillEx := itself.Skill.Ex()
	if skillEx == 0 {
		skillEx = 1
	}

	eqEx := itself.Equipments.Ex()
	if eqEx == 0 {
		eqEx = 1
	}
	return lEx * skillEx * eqEx * itself.Expect
}

type IntSlice []ysPeople

func (s IntSlice) Len() int      { return len(s) }
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less
// 排序原则
func (s IntSlice) Less(i, j int) bool {
	return s[i].weight() > s[j].weight()
}

func runY(cmd *cobra.Command, args []string) {
	var all IntSlice
	jsonStr, err := arms.FileGetContents("y.json")
	json.Unmarshal(jsonStr, &all)
	for i, _ := range all {
		all[i].Sort = i
	}
	jsonData, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	arms.Put(jsonData, "y.json")

	sort.Stable(all)
	jsonData, err = json.MarshalIndent(all, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cast.ToString(jsonData))
	//helpers.Put(jsonData, "y.json")
}
