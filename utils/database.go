package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

func Init(models ...interface{}) error {
    var err error // エラーをキャッチする変数を追加します。
    dsn := "n000r111:password@tcp(127.0.0.1:3306)/message_api?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    db.Debug()

    err = db.AutoMigrate(models...)
    if err != nil {
        return err
    }
    return nil
}

func GetDB() *gorm.DB {
    return db
}
