package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	// 南平，这里的db的驱动如果是seata XA不能AutoMigrate
	return db.AutoMigrate(&Tag{})
}
