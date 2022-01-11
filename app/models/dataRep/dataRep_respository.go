package dataRep

import "thh/helpers/db"

var dr *DataRepository

func init() {
	dr = new(DataRepository)
}

type DataRepository struct {
}

func GetDataRepository() *DataRepository {
	return dr
}

func (itself *DataRepository) Set(key string, value string) (err error) {
	var dataRep DataRep
	dataRep.Key = key
	dataRep.Value = value
	if err = DB.SqlDBIns().Where(&DataRep{Key: key}).First(&dataRep).Error; err != nil {
		dataRep.Key = key
		dataRep.Value = value
		if err = DB.SqlDBIns().Create(&dataRep).Error; err != nil {
			return err
		}
	} else {
		dataRep.Key = key
		dataRep.Value = value
		err = DB.SqlDBIns().Save(&dataRep).Error
	}
	return nil
}

func (itself *DataRepository) Get(key string) string {
	var dataRep DataRep
	DB.SqlDBIns().Where(&DataRep{
		Key: key,
	}).First(&dataRep)
	return dataRep.Value
}

func (itself *DataRepository) Del(key string) {
	DB.SqlDBIns().Delete(&DataRep{
		Key: key,
	})
}
