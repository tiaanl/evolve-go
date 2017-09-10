package evolve

func newAlterTableColumns() *alterTableColumns {
	return &alterTableColumns{}
}

type alterTableColumns struct {
	toDrop []string
	toAdd  []*Column
}

func (atc *alterTableColumns) dropColumn(name string) {
	atc.toDrop = append(atc.toDrop, name)
}

func (atc *alterTableColumns) addColumns(columns ...*Column) {
	for _, column := range columns {
		atc.toAdd = append(atc.toAdd, column)
	}
}
