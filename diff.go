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

		for _, currentColumn := range currentTable.Columns() {
			targetColumn := targetTable.Column(currentColumn.Name)

			// Check if we have to drop the column in the current table.
			if targetColumn == nil {
				atc.dropColumn(currentColumn.Name)
				continue
			}

			// Check if we have to alter the current column.
			if !currentColumn.Equals(targetColumn) {
				atc.changeColumn(targetColumn)
			}
		}

		// Check if we have to create the column in the current table.
		for _, targetColumn := range targetTable.Columns() {
			currentColumn := currentTable.Column(targetColumn.Name)
			if currentColumn == nil {
				atc.addColumns(targetColumn)
			}
		}

		changeSet.AlterTable(currentTable.Name(), atc)
	}

	return changeSet, nil
}
