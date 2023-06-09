// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "email", Type: field.TypeString, Size: 100},
		{Name: "address", Type: field.TypeString, Size: 200},
		{Name: "phone_number", Type: field.TypeString, Size: 20},
		{Name: "identify_number", Type: field.TypeString, Size: 12},
		{Name: "date_of_birth", Type: field.TypeTime},
		{Name: "member_code", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CustomersTable,
	}
)

func init() {
}
