package factories

import (
	"thh/app/models/about"
)

func MakeAbouts(count int) []about.About {

	var objs []about.About

	// 设置唯一性，如 About 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		aboutModel := about.About{
			// FIXME()
		}
		objs = append(objs, aboutModel)
	}

	return objs
}
