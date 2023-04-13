package config

import "github.com/dimasyudhana/alterra-group-project-2/entities"

func Migrate(c *Config) {
	db, err := GetConnection(c)
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(entities.User{}, entities.Book{}); err != nil {
		panic(err)
	}
}
