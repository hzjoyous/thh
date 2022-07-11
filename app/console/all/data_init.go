package all

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/app/models/Permission"
	"thh/app/models/Role"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "d:dataInit", Short: "", Run: dataInit})
}

func dataInit(_ *cobra.Command, _ []string) {
	p := Permission.Permission{}
	p.Name = "pNam2e"
	Permission.Create(&p)
	fmt.Println("end", p)
	r := Role.Role{}
	r.Name = time.Now().Format("2006-01-02 15:04:05")
	Role.Create(&r)
}
