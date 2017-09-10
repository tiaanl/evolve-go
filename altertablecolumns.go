package evolve

func newAlterTableColumns() *alterTableColumns {
	return &alterTableColumns{
		toDrop:   []string{},
		toAdd:    []*Column{},
		toChange: []*Column{},
	}
}

type alterTableColumns struct {
	toDrop   []string
	toAdd    []*Column
	toChange []*Column
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

func (atc *alterTableColumns) changeColumn(columns ...*Column) {
	for _, column := range columns {
		atc.toChange = append(atc.toChange, column)
	}
}
