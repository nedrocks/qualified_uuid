package quuid

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type QualifiedUUID struct {
	base      uuid.UUID
	qualifier string
}

func (q QualifiedUUID) String() string {
	return fmt.Sprintf("%s.%s", q.qualifier, q.base.String())
}

func NewQualifiedUUIDV4(qualifier string) QualifiedUUID {
	return QualifiedUUID{
		base:      uuid.NewV4(),
		qualifier: qualifier,
	}
}

func QualifyUUID(qualifier string, base uuid.UUID) QualifiedUUID {
	return QualifiedUUID{
		base:      base,
		qualifier: qualifier,
	}
}
