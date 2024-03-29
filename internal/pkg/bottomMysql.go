package pkg

import (
	"github.com/BarnabyCharles/frame/databases/mysql"
	"gorm.io/gorm"
)

type TypeBottom struct {
	gorm.Model
	Name  string `gorm:"INDEX,type:varchar(20)"`
	Url   string `gorm:"INDEX,type:text"`
	Image string `gorm:"INDEX,type:text"`
}

func NewBottom() *TypeBottom {
	return new(TypeBottom)
}

func SearchBottom() ([]TypeBottom, error) {
	Bottoms := NewBottom()
	BottomInfos := []TypeBottom{}
	err := mysql.DB.Model(Bottoms).Where("deleted_at IS NULL").Limit(5).Find(&BottomInfos).Error

	if err != nil {
		return nil, err
	}
	return BottomInfos, nil
}
