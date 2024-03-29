package pkg

import (
	"github.com/BarnabyCharles/frame/databases/mysql"
	"gorm.io/gorm"
)

type TypeSlideshow struct {
	gorm.Model
	Name string `gorm:"INDEX,type:varchar(20)"`
	Img  string `gorm:"INDEX,type:text"`
	Url  string `gorm:"INDEX,type:text"`
}

func NewSlideshow() *TypeSlideshow {
	return new(TypeSlideshow)
}

func SearchSlideshow() ([]TypeSlideshow, error) {
	slideshow := NewSlideshow()
	slideshowInfos := []TypeSlideshow{}
	err := mysql.DB.Model(slideshow).Where("deleted_at IS NULL").Limit(6).Find(&slideshowInfos).Error

	if err != nil {
		return nil, err
	}
	return slideshowInfos, nil
}
