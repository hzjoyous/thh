package RolePermission

func Get(id any) (entity RolePermission) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity RolePermission) {
	builder().Save(&entity)
}

func UpdateAll(entities []RolePermission) {
	builder().Save(&entities)
}

func Delete(entity RolePermission) {
	builder().Delete(&entity)
}

func Save(entity RolePermission) {
	builder().Save(&entity)
}

func Create(entity *RolePermission) {
	builder().Create(&entity)

}
