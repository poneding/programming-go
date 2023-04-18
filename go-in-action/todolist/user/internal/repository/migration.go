package repository

import "user/pkg/util"

func migration() {
	if err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}); err != nil {
		util.LogrusObject.Errorln("migrate table failed")
		panic(err)
	}

	util.LogrusObject.Infoln("migrate table success")
}
