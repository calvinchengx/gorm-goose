package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Up_20130106222315(txn *gorm.DB) {
	fmt.Println("Hello from migration 20130106222315 Up!")
}

func Down_20130106222315(txn *gorm.DB) {
	fmt.Println("Hello from migration 20130106222315 Down!")
}
