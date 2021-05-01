package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var (
//	DB *gorm.DB
//
//	username string = "root"
//	password string = "root"
//	dbName   string = "spiders"
//)
//
//func init() {
//	var err error
//	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(10.10.10.253:13306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)) //nolint:lll
//	if err != nil {                                                                                                          //nolint:wsl
//		log.Fatalf(" gorm.Open.err: %v", err)
//	}
//
//	DB.SingularTable(true)
//	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
//		return "sp_" + defaultTableName
//	}
//}
