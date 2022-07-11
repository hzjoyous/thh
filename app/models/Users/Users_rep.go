package Users

func Get(id any) (entity Users) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity Users) {
	builder().Save(&entity)
}

func UpdateAll(entities []Users) {
	builder().Save(&entities)
}

func Delete(entity Users) {
	builder().Delete(&entity)
}
