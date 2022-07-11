package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:signal", Short: "demo:signal", Run: rbacHandle})
}

func rbacHandle(cmd *cobra.Command, args []string) {
	admin := map[string]any{
		"admin": []string{
			"admin1",
			"admin2",
		},
		"adminRole": []string{
			"id",
		},
		"adminPermission": []string{
			"id",
		},
		"adminRoleRelation": []any{
			struct {
				adminId int64
				roleId  int64
			}{},
		},
		"rolePermissionRelation": []any{
			struct {
				permission int64
				roleId     int64
			}{},
		},
	}

	fmt.Println(admin)
}
