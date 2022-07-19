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
	appendCommand(&cobra.Command{Use: "d:gdcRun", Short: "", Run: gdcRun})
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

func gdcRun(_ *cobra.Command, _ []string) {
	fmt.Println(_gcd(10, 20))
}

func _gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return _gcd(b, a%b)
}
