package auto

import (
	"astro/database"
	"context"
	"database/sql"
	"log"
	"time"
)

//Load is
func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("this is an error :", err)
	}
	defer db.Close()

	// Creating Astro Table
	createAstro := `CREATE TABLE IF NOT EXISTS Astro (
							id int NOT NULL AUTO_INCREMENT,
							urlSlug varchar(255),
							namePrefix varchar(255),
							firstName varchar(255),
							lastName varchar(255),
							aboutMe varchar(255),
							profliePicUrl varchar(255),
							experience FLOAT,
							rating int,
							minimumCallDuration FLOAT,
							minimumCallDurationCharges FLOAT,
							additionalPerMinuteCharges FLOAT,
							isAvailable bool,
							isOnCall bool,
							PRIMARY KEY (id)
						)`

	createImageInfo := `CREATE TABLE IF NOT EXISTS ImageInfo (
							id int,
							imageUrl varchar(255),
							pid int,
							imageSize varchar(255),
							imageType varchar(255)
						)`

	createDays := `CREATE TABLE IF NOT EXISTS Days (
							id int,
							day varchar(255)
						)`

	createSlotInfo := `CREATE TABLE IF NOT EXISTS SlotInfo (
							id int,
							toFormat varchar(255),
							fromFormat varchar(255),
							fromTime varchar(255),
							toTime varchar(255)
						)`
	createSkills := `CREATE TABLE IF NOT EXISTS Skills (
							id int,
							skill varchar(255)
						)`
	createLanguage := `CREATE TABLE IF NOT Exists Language (
							id int,
							language varchar(255)
						)`
	createBannerData := `CREATE TABLE IF NOT Exists BannerData (
							id int NOT NULL AUTO_INCREMENT,
							name varchar(255),
							landingUrl varchar(255),
							orderType varchar(255),
							productCode varchar(255),
							productName varchar(255),
							professionalSlug varchar(255),
							PRIMARY KEY (id)
						)`
	

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	ExecuteQuery(createAstro, ctx, db)
	ExecuteQuery(createImageInfo, ctx, db)
	ExecuteQuery(createDays, ctx, db)
	ExecuteQuery(createSlotInfo, ctx, db)
	ExecuteQuery(createSkills, ctx, db)
	ExecuteQuery(createLanguage, ctx, db)
	ExecuteQuery(createBannerData, ctx, db)
}

func ExecuteQuery(query string, ctx context.Context, db *sql.DB) {
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
	}
	log.Printf("Rows affected when creating table: %d", rows)
}
