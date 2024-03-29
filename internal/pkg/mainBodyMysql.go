package pkg

import (
	"github.com/BarnabyCharles/frame/databases/mysql"
	"gorm.io/gorm"
)

type TypeMainBody struct {
	gorm.Model
	Name  string `gorm:"INDEX,type:varchar(20)"`
	Url   string `gorm:"INDEX,type:text"`
	Image string `gorm:"INDEX,type:text"`
}

func NewMainBody() *TypeMainBody {
	return new(TypeMainBody)
}

func SearchMainBody() ([]TypeMainBody, error) {
	mainBody := NewMainBody()
	mainBodyInfos := []TypeMainBody{}
	err := mysql.DB.Model(mainBody).Where("deleted_at IS NULL").Limit(10).Find(&mainBodyInfos).Error

	if err != nil {
		return nil, err
	}
	return mainBodyInfos, nil
}
