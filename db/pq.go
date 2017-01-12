package db

import (
	"database/sql"

	_ "github.com/lib/pq" //pq lib will register itself against std lib package
	"github.com/psionikangel/scribd/models"
)

func getPQClient() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://scribd:c0g1n0v@postgres:5432/scribd?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

//NewRun : Inserts a new run that is starting
func NewRun(r models.Run) {
	client := getPQClient()
	defer client.Close()
	var lastInsertedID string
	err := client.QueryRow(`insert into run (id, starttime, endtime, machinename) values ($1,$2,$3,$4) returning id;`, r.ID, r.Start, r.End, r.Machinename).Scan(&lastInsertedID)
	if err != nil {
		panic(err)
	}
}

//EndRun : Ends a run by updating the end time
func EndRun(r models.Run) {
	client := getPQClient()
	defer client.Close()
	var lastUpdatedID string
	err := client.QueryRow(`update run set endtime = $1 where id = $2 returning id;`, r.End, r.ID).Scan(&lastUpdatedID)
	if err != nil {
		panic(err)
	}
}

//AddMetadata : Adds a row of metadata from a run
func AddMetadata(m models.Metadata) {
	client := getPQClient()
	defer client.Close()
	var lastInsertedID string
	err := client.QueryRow(`insert into metadata (filepath, filesize, lastmodified, filename, extension, checksum, runid) values ($1,$2,$3,$4,$5,$6,$7) returning id;`,
		m.Path, m.Filesize, m.LastModified, m.Filename, m.Extension, m.Checksum, m.RunID).Scan(&lastInsertedID)
	if err != nil {
		panic(err)
	}
}
