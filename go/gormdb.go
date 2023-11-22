package openapi

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"path/filepath"

	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (Cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Cfg.Host, Cfg.Port, Cfg.User, Cfg.Password, Cfg.Database, Cfg.SSLMode)

}

func OpenDb() error {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "cfd",
		SSLMode:  "disable",
	}

	var err error
	DB, err = gorm.Open(postgres.Open(cfg.String()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&MenuItem{}, CartItem{})

	return nil
}

func SeedMenuItems() {

	DB.Migrator().DropTable(&MenuItem{})
	DB.AutoMigrate(&MenuItem{})

	createMenuItem("Fresh from the tap", "Water", 1.99, 1, "water")
	createMenuItem("Chicken Wrap - Sandwich", "Chicken Wrap", 14.99, 2, "wrap")
	createMenuItem("A slow cooked stew", "Stew", 12.99, 3, "stew")
	createMenuItem("It looks good in the menu picture", "Tomato Soup", 4.99, 4, "soup")
	createMenuItem("A green salad", "Salad", 4.99, 5, "salad")

	var items []MenuItem
	DB.Find(&items)

	for _, item := range items {
		imagepath, _ := filepath.Abs("./go/images/" + item.ImageName + ".jpg")
		img, _ := decodeJpgImage(imagepath)
		buf := new(bytes.Buffer)
		jpeg.Encode(buf, img, nil)
		item.Image = buf.Bytes()
		DB.Save(&item)
	}

	DB.Debug().AutoMigrate(&MenuItem{})

	DB.Debug().AutoMigrate(&CartItem{})
}

func createMenuItem(desc string, name string, price float32, imageid int32, imagename string) {
	err := DB.Create(&MenuItem{
		Description: desc,
		Name:        name,
		Price:       price,
		ImageId:     imageid,
		ImageName:   imagename,
	}).Error
	if err != nil {
		panic(err.Error())
	}
}
