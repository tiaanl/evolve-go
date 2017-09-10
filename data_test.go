package evolve

var (
	testTable = NewTable("test_table")

	testColumnPrimary  = NewColumnPrimary("id")
	testColumnString   = NewColumnString("username", 50)
	testColumnInteger  = NewColumnInteger("age")
	testColumnDateTime = NewColumnDateTime("last_login")

	testColumnInvalid = &Column{
		Type: 2345,
	}
)

func init() {
	testTable.AddColumns(testColumnPrimary, testColumnString, testColumnInteger, testColumnDateTime)
}
