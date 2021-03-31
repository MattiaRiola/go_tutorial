package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Decorations struct {
	Host_uuid string
	Hostname  string
}

type DecorationsJSON struct {
	Host_uuid string `json:"host_uuid"`
	Hostname  string `json:"hostname"`
}

func (d *Decorations) MarshalJSON() ([]byte, error) {
	return json.Marshal(DecorationsJSON{
		d.Host_uuid,
		d.Hostname,
	})
}

func (d *Decorations) UnmarshalJSON(b []byte) error {
	temp := &DecorationsJSON{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	d.Host_uuid = temp.Host_uuid
	d.Hostname = temp.Hostname
	return nil
}

type Snaplog struct {
	Snapshot       []json.RawMessage
	Action         string
	Name           string
	QueryName      string //Extra: this is the name of the collection
	HostIdentifier string
	CalendarTime   string
	UnixTime       int
	Ip             string //Extra: I've no access to this information here in the log I've to query the fleet sql db
	Dec            Decorations
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

func (sl *Snaplog) MarshalJSON() ([]byte, error) {

	raw_decorations, err := json.Marshal(&Decorations{sl.Dec.Host_uuid, sl.Dec.Hostname})
	if err != nil {
		log.Println(err)
	}
	return json.Marshal(SnaplogJSON{
		sl.Snapshot,
		sl.Action,
		sl.Name,
		sl.HostIdentifier,
		sl.CalendarTime,
		sl.UnixTime,
		-1,    //sl.Epoch,
		-1,    //sl.Counter,
		false, //sl.Numerics,
		json.RawMessage(raw_decorations),
	})
}

func (sl *Snaplog) UnmarshalJSON(b []byte) error {
	temp := &SnaplogJSON{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	sl.Snapshot = temp.Snapshot
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
	tmp_dec := new(Decorations)
	b_decorations := []byte(temp.Decorations)
	err = json.Unmarshal(b_decorations, &tmp_dec)
	if err != nil {
		log.Println(err)
	}
	sl.Dec = *tmp_dec
	return nil
}

// type DifferentialLog struct {
// 	//TODO: DifferentialLog structure
// }
