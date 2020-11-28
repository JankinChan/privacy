// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// JdColumns holds the columns for the "jd" table.
	JdColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "id_number", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeInt64},
	}
	// JdTable holds the schema information for the "jd" table.
	JdTable = &schema.Table{
		Name:        "jd",
		Columns:     JdColumns,
		PrimaryKey:  []*schema.Column{JdColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// QqColumns holds the columns for the "qq" table.
	QqColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "qq_number", Type: field.TypeInt64},
		{Name: "phone_number", Type: field.TypeInt64},
	}
	// QqTable holds the schema information for the "qq" table.
	QqTable = &schema.Table{
		Name:        "qq",
		Columns:     QqColumns,
		PrimaryKey:  []*schema.Column{QqColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SfColumns holds the columns for the "sf" table.
	SfColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// SfTable holds the schema information for the "sf" table.
	SfTable = &schema.Table{
		Name:        "sf",
		Columns:     SfColumns,
		PrimaryKey:  []*schema.Column{SfColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JdTable,
		QqTable,
		SfTable,
	}
)

func init() {
}