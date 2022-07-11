package tool

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "z:makedaytodo", Short: "this is a z:makedaytodo", Run: makeDayTodo})
}
func makeDayTodo(cmd *cobra.Command, args []string) {
	now := time.Now().Unix()
	formatStr := "2006/01/02"
	sumDay := 30
	chineseWeekDay := map[time.Weekday]string{
		time.Sunday:    "Sunday",    //"星期日",
		time.Monday:    "Monday",    //"星期一",
		time.Tuesday:   "Tuesday",   //"星期二",
		time.Wednesday: "Wednesday", //"星期三",
		time.Thursday:  "Thursday",  //"星期四",
		time.Friday:    "Friday",    //"星期五",
		time.Saturday:  "Saturday",  //"星期六",
	}
	for i := sumDay; i > -16; i-- {
		timeEntity := time.Unix(now+cast.ToInt64(i*83600), 0)

		fmt.Println("# ", timeEntity.Format(formatStr), " ", chineseWeekDay[timeEntity.Weekday()])
	}
}
