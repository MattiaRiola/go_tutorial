package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Snaplog struct {
	Snapshot       []interface{}
	Action         string
	Name           string
	QueryName      string //Extra: this is the name of the collection
	HostIdentifier string
	CalendarTime   string
	UnixTime       int
	Ip             string //Extra: I've no access to this information here in the log I've to query the fleet sql db
	Dec            interface{}
	// Dec            Decorations
	// Epoch          int
	// Counter        int
	// Numerics       bool
}

type SnaplogJSON struct {
	Snapshot       []json.RawMessage `json:"snapshot"`
	Action         string            `json:"action"`
	Name           string            `json:"name"`
	HostIdentifier string            `json:"hostIdentifier"`
	CalendarTime   string            `json:"calendarTime"`
	UnixTime       int               `json:"unixTime"`
	Epoch          int               `json:"epoch"`
	Counter        int               `json:"counter"`
	Numerics       bool              `json:"numerics"`
	Decorations    json.RawMessage   `json:"decorations"`
}

func (sl *Snaplog) UnmarshalJSON(b []byte) error {
	temp := &SnaplogJSON{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	var result []interface{}
	for _, q := range temp.Snapshot {
		var r interface{}
		err := json.Unmarshal(q, &r)
		if err != nil {
			log.Println(err)
		}
		log.Println(r)
		result = append(result, r)
	}

	log.Println("THIS IS THE RESULT:")
	log.Println(result)
	sl.Snapshot = result
	sl.Action = temp.Action
	sl.Name = temp.Name
	splitter := func(c rune) bool {
		return c == '/'
	}
	tmpPathFields := strings.FieldsFunc(temp.Name, splitter)
	sl.QueryName = tmpPathFields[2] // in this way I select only the query name from the "query path"
	sl.HostIdentifier = temp.HostIdentifier
	sl.CalendarTime = temp.CalendarTime
	sl.UnixTime = temp.UnixTime
	// sl.Epoch = temp.Epoch
	// sl.Counter = temp.Counter
	// sl.Numerics = temp.Numerics
	sl.Ip = "TODO"
	//TODO: find and insert the ip during the translation

	var tmp_dec interface{}
	err = json.Unmarshal(temp.Decorations, &tmp_dec)
	if err != nil {
		log.Println(err)
	}
	sl.Dec = tmp_dec
	return nil
}

// type DifferentialLog struct {
// 	//TODO: DifferentialLog structure
// }
