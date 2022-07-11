package Permission

func Get(id any) (entity Permission) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity Permission) {
	builder().Save(&entity)
}

func UpdateAll(entities []Permission) {
	builder().Save(&entities)
}

func Delete(entity Permission) {
	builder().Delete(&entity)
}

func Save(entity Permission) {
	builder().Save(&entity)
}

func Create(entity *Permission) {
	builder().Create(&entity)

}
