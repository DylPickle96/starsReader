package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type starValues struct {
	location string
	world    string
	minTime  string
	maxTime  string
}

func main() {
	db, err := sql.Open("sqlite3", "./stardb.db")
	if err != nil {
		log.Printf("OPEN ERROR: %v", err)
		return
	}
	rows, err := db.Query("SELECT location, world, minTime, maxTime FROM data WHERE maxTime > strftime('%s', 'now') ORDER BY minTime")
	// rows, err := db.Query("SELECT location, world, minTime, maxTime FROM data")
	if err != nil {
		log.Printf("PREPARE ERROR: %v", err)
		return
	}
	svSlice := []starValues{}
	sv := starValues{}
	for rows.Next() {
		rows.Scan(&sv.location, &sv.world, &sv.minTime, &sv.maxTime)
		svSlice = append(svSlice, sv)
	}
	for _, sv := range svSlice {
		minParsed, _ := strconv.ParseInt(sv.minTime, 10, 64)
		maxParsed, _ := strconv.ParseInt(sv.maxTime, 10, 64)
		minTime := time.Unix(minParsed, 0)
		maxTime := time.Unix(maxParsed, 0)

		fmt.Printf("STAR LOCATION: %s is between: %v AND %v. On world: %s\n", locationConverter(sv.location), minTime, maxTime, sv.world)
		fmt.Println()
	}
}

func locationConverter(location string) string {
	switch location {
	case "0":
		return "Asgarnia"
	case "1":
		return "Crandor Karamja"
	case "2":
		return "Feldip Hills or on the Isle of Souls"
	case "3":
		return "Fossil Island or on Mos Le'Harmless"
	case "4":
		return "Fremennik Lands or on Lunar Isle"
	case "5":
		return "Great Kourend"
	case "6":
		return "Kandarin"
	case "7":
		return "Kebos Lowland"
	case "8":
		return "Kharidian Desert"
	case "9":
		return "Misthalin"
	case "10":
		return "Morytania"
	case "11":
		return "Piscatoris or the Gnome Stronghold"
	case "12":
		return "Tirannwn"
	case "13":
		return "Wilderness"
	case "14":
		return "Unknown"
	default:
		return "you fucked up nerd"
	}
}
