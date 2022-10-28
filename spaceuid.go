package spaceuid

import (
	"fmt"
)

type SpaceUID interface {
	New(string) (SpaceUID, error)
	SUID() string   //identity-space unique ID
	Space() string  //identity-space name
	String() string // "[space]_[id]"
}

type SpaceUIDs []SpaceUID

type SpaceUIDsMap map[string]SpaceUID

func (ids SpaceUIDsMap) GetID(space string, idStr string) (SpaceUID, error) {
	id, ok := ids[space]
	if !ok {
		return nil, fmt.Errorf("unknown identity space:%s", space)
	}
	return id.New(idStr)
}
