package db

import (
	"AquaSecurityChallenge/pkg/models"
	"database/sql"
)

func GetAllHostsFromDB() []models.Host {
	db := GetDB()

	rows, err := db.Query("SELECT * FROM hosts")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	checkErr(err)

	return ConvertToHostArray(rows)
}

func GetHostByIDFromDB(hostId int) *models.Host {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM hosts WHERE id = $1 ", hostId)
	defer rows.Close()
	checkErr(err)

	var host *models.Host = nil

	if rows.Next() {
		host = ConvertToHost(rows)
	}
	return host
}

func ConvertToHostArray(rows *sql.Rows) []models.Host {
	hosts := make([]models.Host, 0)
	for rows.Next() {
		host := ConvertToHost(rows)
		hosts = append(hosts, *host)
	}
	return hosts
}

func ConvertToHost(row *sql.Rows) *models.Host {
	host := &models.Host{}
	err := row.Scan(&host.ID, &host.Uuid, &host.Name, &host.Ip_address)
	checkErr(err)

	return host
}
