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

		// Check if there are any schema differences between the two tables.

	}

	return changeSet, nil
}
