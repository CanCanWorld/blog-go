package model

func migration() {
	// 自动迁移
	DB.Set("gorm:table_option", "charset=uft8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
