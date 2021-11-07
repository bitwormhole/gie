package vo

import "github.com/bitwormhole/gie/app/web/dto"

type DBA struct {
	BaseVO

	Action     string //  start,stop
	Running    bool
	Cancelling bool
	Snapshots  []*dto.DBASnapshot
	Latest     dto.DBASnapshot
}
