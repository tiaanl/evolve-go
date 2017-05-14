package evolve

type MigrationList interface {
	Exists(name string) (bool, error)
	Add(migrationNames ...string) error
}

func NewMigrationList(backEnd BackEnd) MigrationList {
	return &migrationList{
		backEnd:            backEnd,
		existingMigrations: []string{},

		synced:                 false,
		migrationsTableCreated: false,
	}
}

type migrationList struct {
	backEnd            BackEnd
	existingMigrations []string

	synced                 bool
	migrationsTableCreated bool
}

func (ml *migrationList) Exists(name string) (bool, error) {
	if err := ml.ensureMigrationsTableExists(); err != nil {
		return false, err
	}

	// If the list of existing migrations is empty, we assume we haven't retrieved it from the database, so we update
	// it.
	if len(ml.existingMigrations) == 0 {
		// Get the list of migration names from the database.
		existingMigrations, err := ml.getMigrationsFromDatabase()
		if err != nil {
			return false, err
		}

		ml.existingMigrations = existingMigrations
	}

	// See if the migration exists in the list of existing migrations.
	for _, m := range ml.existingMigrations {
		if m == name {
			return true, nil
		}
	}

	return false, nil
}

func (ml *migrationList) getMigrationsFromDatabase() ([]string, error) {
	// If we already pulled the data from the database, then don't do it again.
	if ml.synced {
		return ml.existingMigrations, nil
	}

	migrations := []string{}

	sql := "SELECT name FROM migrations ORDER BY name ASC"

	rows, err := ml.backEnd.Connection().Query(sql)
	if err != nil {
		return migrations, err
	}
	defer rows.Close()

	for rows.Next() {
		var migrationName string
		if err := rows.Scan(&migrationName); err != nil {
			return []string{}, err
		}
		migrations = append(migrations, migrationName)
	}

	// Store the fact that we pulled the existing migrations from the db.
	ml.synced = true

	return migrations, nil
}

func (ml *migrationList) Add(migrationNames ...string) error {
	if err := ml.ensureMigrationsTableExists(); err != nil {
		return err
	}

	// Persist the new migration names in the database.
	for _, migrationName := range migrationNames {
		err := ml.backEnd.InsertData("migrations", []string{"name"}, []string{migrationName})
		if err != nil {
			return err
		}
	}

	// Add the new migration names to our internal cache of existing migration names.
	for _, m := range migrationNames {
		ml.existingMigrations = append(ml.existingMigrations, m)
	}

	return nil
}

func (ml *migrationList) ensureMigrationsTableExists() error {
	// If we already tried to create the table, then don't do it again.
	if ml.migrationsTableCreated {
		return nil
	}

	// Make sure the existingMigrations table exists.
	table := NewTable("migrations")
	table.String("name", 50).IsPrimary(true)

	err := ml.backEnd.CreateTableIfNotExists(table)
	if err != nil {
		return err
	}

	// Store the fact that we successfully made sure that the table exists.
	ml.migrationsTableCreated = true

	return nil
}
