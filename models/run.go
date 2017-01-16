package models

import "time"

//Run : A instance when scribd was run
type Run struct {
	ID          string
	Machinename string
	Start       time.Time
	End         time.Time
}

//Runlist : A list of runs
type Runlist struct {
	Runs []Run
}
