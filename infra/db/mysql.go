package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Db *gorm.DB

func CreateMySqlDB(dsn string) error {
    if Db != nil {
        return nil
    }
    var err error
    
    Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    return nil
}