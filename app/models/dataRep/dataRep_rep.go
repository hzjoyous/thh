package dataRep

import (
	"thh/conf/dbconnect"
)

func Set(key string, value string) (err error) {
	var dataRep DataRep
	dataRep.Key = key
	dataRep.Value = value
	if err = dbconnect.Std().Where(&DataRep{Key: key}).First(&dataRep).Error; err != nil {
		dataRep.Key = key
		dataRep.Value = value
		if err = dbconnect.Std().Create(&dataRep).Error; err != nil {
			return err
		}
	} else {
		dataRep.Key = key
		dataRep.Value = value
		err = dbconnect.Std().Save(&dataRep).Error
	}
	return nil
}

func Get(key string) string {
	var dataRep DataRep
	dbconnect.Std().Where(&DataRep{
		Key: key,
	}).First(&dataRep)
	return dataRep.Value
}

func Del(key string) {
	dbconnect.Std().Delete(&DataRep{
		Key: key,
	})
}
