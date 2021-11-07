package repository

import (
	"crypto/sha1"
	"strconv"
	"strings"

	"github.com/bitwormhole/gie/app/data/entity"
	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/util"
)

// BaseRepo ...
type BaseRepo struct {
	source         ptable.DataSource
	session        ptable.Session
	table          ptable.Table
	intIDGenerator uint64
}

func (inst *BaseRepo) initBase(ds ptable.DataSource, pto *ptable.TableOpen) error {

	session := ds.GetSession()

	pto.DoInit = true
	table, err := session.DB().OpenTable(pto)
	if err != nil {
		return err
	}

	inst.table = table
	inst.source = ds
	inst.session = session
	return nil
}

func (inst *BaseRepo) saveBase(row ptable.Row, e *entity.BaseEntity) {

	time1 := strconv.FormatInt(e.CreatedAt, 10)
	time2 := strconv.FormatInt(e.UpdatedAt, 10)

	row.SetValue("id", e.ID)
	row.SetValue("uuid", e.UUID)
	row.SetValue("owner", e.Owner)
	row.SetValue("creator", e.Creator)
	row.SetValue("createdat", time1)
	row.SetValue("updatedat", time2)
}

func (inst *BaseRepo) loadBase(row ptable.Row, e *entity.BaseEntity) {

	id, _ := row.GetValue("id")
	uuid, _ := row.GetValue("uuid")
	owner, _ := row.GetValue("owner")
	creator, _ := row.GetValue("creator")
	time1, _ := row.GetValue("createdat")
	time2, _ := row.GetValue("updatedat")

	e.ID = id
	e.UUID = uuid
	e.Owner = owner
	e.Creator = creator

	createdAt, _ := strconv.ParseInt(time1, 10, 64)
	updatedAt, _ := strconv.ParseInt(time2, 10, 64)

	e.CreatedAt = createdAt
	e.UpdatedAt = updatedAt
}

func (inst *BaseRepo) existsID(id uint64) bool {
	session := inst.session
	table := inst.table
	p := session.GetProperties(table)
	k1 := table.Name()
	k2 := strconv.FormatUint(id, 10)
	k3 := table.PrimaryKey()
	key := k1 + "." + k2 + "." + k3
	value := p.GetProperty(key, "")
	return len(value) > 0
}

func (inst *BaseRepo) generateNextID() string {
	id := inst.intIDGenerator
	if inst.existsID(id) {
		// seek for free id
		id = inst.generateNextIntID2(id)
	}
	inst.intIDGenerator = id + 1
	return strconv.FormatUint(id, 10)
}

func (inst *BaseRepo) generateNextIntID2(from uint64) uint64 {

	const idmax = 0xffffffff
	const idmin = 1000

	if from < idmin {
		from = idmin
	}

	// seed for to
	to := from * 2
	for ; to < idmax; to *= 2 {
		if !inst.existsID(to) {
			break
		}
	}

	for id := from; id < to; id++ {
		if inst.existsID(id) {
			continue
		} else {
			return id
		}
	}

	return to
}

func (inst *BaseRepo) generateNextUUID(now int64, tag string) string {
	if now < 1 {
		now = util.CurrentTimestamp()
	}
	const nl = "\n"
	builder := strings.Builder{}
	builder.WriteString(nl)
	builder.WriteString(inst.table.Name())
	builder.WriteString(nl)
	builder.WriteString(strconv.FormatInt(now, 10))
	builder.WriteString(nl)
	builder.WriteString(strconv.FormatUint(inst.intIDGenerator, 10))
	builder.WriteString(nl)
	builder.WriteString(tag)
	builder.WriteString(nl)
	str := builder.String()
	sum := sha1.Sum([]byte(str))
	return util.StringifyBytes(sum[:])
}
