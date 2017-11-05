package evolve

var (
	testTable = NewTable("test_table")

	testColumnPrimary  = PrimaryColumn("id")
	testColumnString   = StringColumn("username", 50)
	testColumnInteger  = IntegerColumn("age")
	testColumnDateTime = DateTimeColumn("last_login")

	testColumnInvalid = &Column{
		Type: 2345,
	}
)

func init() {
	testTable.AddColumns(testColumnPrimary, testColumnString, testColumnInteger, testColumnDateTime)
}
