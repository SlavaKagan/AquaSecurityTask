package db

import (
	"AquaSecurityChallenge/pkg/models"
	"database/sql"
	"github.com/nu7hatch/gouuid"
)

/* Get all containers from the aqua db with SQL */

func GetAllContainersFromDB() []models.Container {
	db := GetDB()
	rows, err := db.Query("SELECT c.id, c.host_id, c.name,c.image_name, h.name FROM containers AS c JOIN hosts AS h on h.id = c.host_id")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	checkErr(err)

	return ConvertToContainerArray(rows)
}

/* Get containers by id (from the user) from the aqua db with SQL */

func GetContainerByIDFromDB(containerID int) *models.Container {
	db := GetDB()
	rows, err := db.Query(`SELECT c.id, c.host_id, c.name,c.image_name, h.name FROM containers AS c JOIN hosts AS h on h.id = c.host_id WHERE c.id = ?`, containerID)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	checkErr(err)

	var container *models.Container = nil

	if rows.Next() {
		container = ConvertToContainer(rows)
	}
	return container
}

/* Convert co container model - array */

func ConvertToContainerArray(rows *sql.Rows) []models.Container {
	containers := make([]models.Container, 0)
	for rows.Next() {
		container := ConvertToContainer(rows)
		containers = append(containers, *container)
	}
	return containers
}

/* Convert to container model */

func ConvertToContainer(row *sql.Rows) *models.Container {
	container := &models.Container{}
	err := row.Scan(&container.ID, &container.Host_ID, &container.Name, &container.Image_Name, &container.Host_name)
	checkErr(err)
	return container
}

/* Create a new container in the aqua db with SQL */

func CreateContainer(host_id int, image_name string) bool {
	db := GetDB()
	rows, err := db.Query("SELECT 1 FROM hosts WHERE hosts.id = ?", host_id)
	checkErr(err)
	res := false
	found := rows.Next()
	err = rows.Close()
	if err != nil {
		return false
	}
	if found {
		stmt, err := db.Prepare("INSERT INTO containers(host_id, name, image_name) values(?,?,?)")
		checkErr(err)

		u, err := uuid.NewV4()
		_, err = stmt.Exec(host_id, u.String(), image_name)
		checkErr(err)
		res = true
	}
	return res

}

/* Get all container for specific host by id (from the user) from the aqua db with SQL */

func GetContainersForHost(host_id int) []models.Container {
	db := GetDB()
	rows, err := db.Query(`SELECT c.id, c.host_id, c.image_name,c.name, h.name FROM containers AS c JOIN hosts AS h on h.id = c.host_id WHERE c.host_id = ?`, host_id)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	checkErr(err)

	return ConvertToContainerArray(rows)
}
