package backend

import (
	"errors"
	"encoding/json"
)

type EntityId int64
type NodeId int64

// read data from nil
type Unsafe struct {
	Id      *EntityId `db: "id, primarykey, autoincrement" json: ",omitempty" `
	NodeID  *NodeId `db: "node_id" json: ",omitempty"`
	Payload *string `db: "payload" json: ",omitempty"`
}

func (u *Unsafe) FromInterface(raw interface{}) (err error) {
	jsonData, err := json.Marshal(raw); if err != nil {
		return errors.New("Can't marshal raw")
	}
	return json.Unmarshal(jsonData, &u)
}
func (u *Unsafe) toString() string {
	j, _ := json.Marshal(u)
	return string(j)
}

type Entity struct {
	Id      EntityId `db: "id, primarykey, autoincrement" json: ",omitempty" `
	NodeID  NodeId `db: "node_id" json: ",omitempty"`
	Payload string `db: "payload" json: ",omitempty"`
}

func (m *Entity) ApplyRawIO(raw Unsafe) (err error) {
	jsonData, err := json.Marshal(raw); if err != nil {
		return errors.New("Can't marshal raw")
	}
	return json.Unmarshal(jsonData, &m)
}

func (m *Entity) toString() string {
	j, _ := json.Marshal(m)
	return string(j)
}