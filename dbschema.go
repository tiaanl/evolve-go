package evolve

func NewSchemaFromBackEnd(backEnd BackEnd) (Schema, error) {
	return backEnd.BuildSchema()
}
