package demo

import (
	"fmt"
)

func init() {
	addConsole("demo:signal", "demo:signal", rbacHandle)
}

func rbacHandle() {
	admin := map[string]interface{}{
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
		"adminRoleRelation": []interface{}{
			struct {
				adminId int64
				roleId  int64
			}{},
		},
		"rolePermissionRelation": []interface{}{
			struct {
				permission int64
				roleId     int64
			}{},
		},
	}

	fmt.Println(admin)
}
