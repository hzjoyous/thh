package Role

func Get(id any) (entity Role) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity Role) {
	builder().Save(&entity)
}

func UpdateAll(entities []Role) {
	builder().Save(&entities)
}

func Delete(entity Role) {
	builder().Delete(&entity)
}

func Save(entity Role) {
	builder().Save(&entity)
}

func Create(entity *Role) {
	builder().Create(&entity)

}
