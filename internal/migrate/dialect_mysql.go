// Code generated by go generate; DO NOT EDIT.
package migrate

var DialectMySQL = Migrations{
	{
		ID:         "20211121-094727",
		Dialect:    "mysql",
		Statements: []string{"DROP INDEX uix_places_place_label ON `places`;"},
	},
	{
		ID:         "20211124-120008",
		Dialect:    "mysql",
		Statements: []string{"DROP INDEX idx_places_place_label ON `places`;", "DROP INDEX uix_places_label ON `places`;"},
	},
}
