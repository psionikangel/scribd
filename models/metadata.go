package models

import "time"

//Metadata : A file's metadata
type Metadata struct {
	Path         string
	Filesize     int64
	LastModified time.Time
	Filename     string
	Extension    string
	Checksum     string
	RunID        string
}

//MetadataList : A list of metadata
type MetadataList struct {
	Meta []Metadata
}
