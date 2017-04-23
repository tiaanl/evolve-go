package evolve

type Migration struct {
}

func (m *Migration) Up() error {
	return nil
}

func (m *Migration) Down() error {
	return nil
}
