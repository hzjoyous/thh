package dataRep

type DataRep struct {
	Key   string `gorm:"type:varchar(255);not null;primaryKey;default:'';"  json:"key"`
	Value string `gorm:"type:varchar(255);not null;default:'';"  json:"value"`
}
