package evolve

func newAlterTableColumns() *alterTableColumns {
	return &alterTableColumns{
		toDrop:  []string{},
		toAdd:   []*Column{},
		toAlter: []*Column{},
	}
}

type alterTableColumns struct {
	toDrop  []string
	toAdd   []*Column
	toAlter []*Column
}

func (atc *alterTableColumns) dropColumn(names ...string) {
	for _, name := range names {
		atc.toDrop = append(atc.toDrop, name)
	}
}

func (atc *alterTableColumns) addColumns(columns ...*Column) {
	for _, column := range columns {
		atc.toAdd = append(atc.toAdd, column)
	}
}

func (atc *alterTableColumns) alterColumn(columns ...*Column) {
	for _, column := range columns {
		atc.toAlter = append(atc.toAlter, column)
	}
}
