package pkg

import (
	"github.com/BarnabyCharles/frame/databases/mysql"
	"gorm.io/gorm"
)

type TypeTop struct {
	gorm.Model
	Name string `gorm:"INDEX,type:varchar(20)"`
	Url  string `gorm:"INDEX,type:text"`
}

func NewTop() *TypeTop {
	return new(TypeTop)
}

func SearchTopType() ([]TypeTop, error) {
	top := NewTop()
	topInfos := []TypeTop{}
	err := mysql.DB.Model(top).Where("deleted_at IS NULL").Limit(5).Find(&topInfos).Error

	if err != nil {
		return nil, err
	}
	return topInfos, nil
}
