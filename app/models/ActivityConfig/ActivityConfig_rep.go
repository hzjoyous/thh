package ActivityConfig

func Get(id any) (entity ActivityConfig) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity ActivityConfig) {
	builder().Save(&entity)
}

func UpdateAll(entities []ActivityConfig) {
	builder().Save(&entities)
}

func Delete(entity ActivityConfig) {
	builder().Delete(&entity)
}

func Save(entity ActivityConfig) {
	builder().Save(&entity)
}

func Create(entity *ActivityConfig) {
	builder().Create(&entity)

}
