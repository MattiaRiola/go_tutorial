package main

import (
	"encoding/json"
	"time"
)

type Collection struct {
	Mode      string `json: "mode"`
	LogResult string `json: "log_result"`
	//	TODO: write a Log structure
}

type GeneralLog interface {
	new(json.RawMessage)
}

type SnapshotLog struct {
	Mode          string          `json:"mode"`
	QueryResponse json.RawMessage `json: "query_response"`
	QueryName     string          `json: "query_name"` // this will be the name of the collection
	UnixTime      time.Time       `json: "unix_time"`
	HostId        string          `json: "host_id"` //host_uuid or hostIdentifier (the id of the machine) or hostname
	HostName      string          `json: "host_name"`
	Ip            string          `json: "ip"`
	//TODO: SnapshotLog structure
	//from RawMessage To SnapshotLog
}

func (SnapLog *SnapshotLog) new(json.RawMessage) {
	//TODO:

}

func (DiffLog *DifferentialLog) new(json.RawMessage) {
	//TODO:
}

type DifferentialLog struct {
	//TODO: DifferentialLog structure
}
