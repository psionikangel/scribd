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
	err := client.QueryRow(`insert into run (runid, starttime, endtime, machinename) values ($1,$2,$3,$4) returning id;`, r.ID, r.Start, r.End, r.Machinename).Scan(&lastInsertedID)
	if err != nil {
		panic(err)
	}
}

//GetAllRuns : Returns a list of all the runs collected by the server
func GetAllRuns() *models.Runlist {
	client := getPQClient()
	defer client.Close()
	rows, err := client.Query(`select runid, starttime, endtime, machinename, filescount from run order by machinename asc, starttime desc;`)
	if err != nil {
		panic(err)
	}
	list := new(models.Runlist)
	defer rows.Close()
	for rows.Next() {
		run := new(models.Run)
		err := rows.Scan(&run.ID, &run.Start, &run.End, &run.Machinename, &run.FilesCount)
		list.Runs = append(list.Runs, *run)
		if err != nil {
			panic(err)
		}
	}
	return list
}

//EndRun : Ends a run by updating the end time
func EndRun(r models.Run) {
	client := getPQClient()
	defer client.Close()
	var lastUpdatedID string
	err := client.QueryRow(`update run set endtime = $1, filescount = $2 where runid = $3 returning id;`, r.End, r.FilesCount, r.ID).Scan(&lastUpdatedID)
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

//FetchMetadataPerRun : Returns all the metadata for a single run
func FetchMetadataPerRun(runid string) *models.MetadataList {
	client := getPQClient()
	defer client.Close()
	rows, err := client.Query(`select filepath, filesize, lastmodified, filename, extension, checksum from metadata where runid = $1;`, runid)
	if err != nil {
		panic(err)
	}
	metas := new(models.MetadataList)
	defer rows.Close()
	for rows.Next() {
		meta := new(models.Metadata)
		err := rows.Scan(&meta.Path, &meta.Filesize, &meta.LastModified, &meta.Filename, &meta.Extension, &meta.Checksum)
		if err != nil {
			panic(err)
		}
		metas.Meta = append(metas.Meta, *meta)
	}
	return metas
}

//FetchDuplicatesPerRun : Returns all duplicates files found in a given run
func FetchDuplicatesPerRun(runid string) *models.MetadataList {
	client := getPQClient()
	defer client.Close()
	rows, err := client.Query("select filepath, lastmodified, metadata.checksum, filename, filesize, extension, runid, numOcc from metadata inner join (select checksum, count(checksum) as numOcc from metadata where runid = '$1' group by checksum having (count(checksum) > 1 )) as duplicates on metadata.checksum = duplicates.checksum where runid = '$1' and metadata.checksum != '' order by checksum asc;", runid)
	if err != nil {
		panic(err)
	}
	metas := new(models.MetadataList)
	defer rows.Close()
	for rows.Next() {
		meta := new(models.Metadata)
		err := rows.Scan(&meta.Path, &meta.LastModified, &meta.Checksum, &meta.Filename, &meta.Filesize, &meta.Extension)
		if err != nil {
			panic(err)
		}
		metas.Meta = append(metas.Meta, *meta)
	}
	return metas
}
