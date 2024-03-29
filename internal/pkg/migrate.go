package pkg

import "github.com/BarnabyCharles/frame/databases/mysql"

func Migrate() error {
	top := NewTop()
	err := mysql.DB.AutoMigrate(top)
	if err != nil {
		return err
	}

	slideshow := NewSlideshow()
	err = mysql.DB.AutoMigrate(slideshow)
	if err != nil {
		return err
	}

	mainBody := NewMainBody()
	err = mysql.DB.AutoMigrate(mainBody)
	if err != nil {
		return err
	}

	bottom := NewBottom()
	err = mysql.DB.AutoMigrate(bottom)
	if err != nil {
		return err
	}
	return nil
}
