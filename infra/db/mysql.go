package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

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