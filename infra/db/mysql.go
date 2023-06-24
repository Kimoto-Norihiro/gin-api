package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB
var dsn = "n000r111:password@tcp(127.0.0.1:3306)/message_api?charset=utf8mb4&parseTime=True&loc=Local"

func CreateMySqlDB(dsn string) (*gorm.DB, error) {
    if db != nil {
        return db, nil
    }
    var err error
    
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}