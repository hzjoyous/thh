package ActivityLimitConfig

func Get(id any) (entity ActivityLimitConfig) {
	builder().Where(pid, id).First(&entity)
	return
}

func Update(entity ActivityLimitConfig) {
	builder().Save(&entity)
}

func UpdateAll(entities []ActivityLimitConfig) {
	builder().Save(&entities)
}

func Delete(entity ActivityLimitConfig) {
	builder().Delete(&entity)
}

func Save(entity ActivityLimitConfig) {
	builder().Save(&entity)
}

func Create(entity *ActivityLimitConfig) {
	builder().Create(&entity)

}
