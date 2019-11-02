package main

import (
	"github.com/jinzhu/gorm"
	"github.com/swkoubou/torch/server/model/types"
)

func Migrate(db *gorm.DB) error {
	// AreaInfo
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(
			&types.AreaInfo{},
		).
		Error
	if err != nil {
		return err
	}

	// SpotInfo
	err = db.Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(&types.SpotInfo{}).
		AddForeignKey("area_info_id", "area_infos(id)", "CASCADE", "CASCADE").
		Error
	if err != nil {
		return err
	}

	// LikeArea
	err = db.Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(&types.LikeArea{}).
		AddForeignKey("area_info_id", "area_infos(id)", "CASCADE", "CASCADE").
		Error
	if err != nil {
		return err
	}

	// LikeSpot
	err = db.Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(&types.LikeSpot{}).
		AddForeignKey("spot_info_id", "spot_infos(id)", "CASCADE", "CASCADE").
		Error
	if err != nil {
		return err
	}

	return nil
}
