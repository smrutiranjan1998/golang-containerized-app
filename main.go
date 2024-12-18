package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Fixlet struct {
	ID                    uint `gorm:"primaryKey"`
	SiteID                string
	FixletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

func main() {
	dsn := "host=db user=postgres password=postgres dbname=fixletdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Fixlet{})
	fmt.Println("Database connected and migrated successfully!")

	file, err := os.Open("fixlets.csv")
	if err != nil {
		log.Fatal("Failed to open CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV file:", err)
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		count, _ := strconv.Atoi(row[4])
		fixlet := Fixlet{
			SiteID:                row[0],
			FixletID:              row[1],
			Name:                  row[2],
			Criticality:           row[3],
			RelevantComputerCount: count,
		}
		db.Create(&fixlet)
	}

	fmt.Println("CSV data inserted into the database successfully!")
}
