package evolve

func NewChangeSetFromSchameDiff(current, target Schema) (ChangeSet, error) {
	changeSet := NewChangeSet()

	// Check if we need to drop table in the current schema.
	for _, currentTable := range current.Tables() {
		targetTable := target.GetTable(currentTable.Name())

		if targetTable == nil {
			changeSet.DropTable(currentTable.Name())
		}
	}

	// Go through the target's tables to see if we need to create or alter any tables in the current schema.
	for _, targetTable := range target.Tables() {
		currentTable := current.GetTable(targetTable.Name())

		if currentTable == nil {
			changeSet.CreateTable(targetTable)
			continue
		}

		// We have to alter the current table.
		atc := newAlterTableColumns()

		// Check if we have to drop any columns.
		for _, currentColumn := range currentTable.Columns() {
			targetColumn := targetTable.Column(currentColumn.Name)
			if targetColumn == nil {
				atc.dropColumn(targetColumn.Name)
			}
		}

		changeSet.AlterTable(currentTable.Name(), atc)
	}

	return changeSet, nil
}
