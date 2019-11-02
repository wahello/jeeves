// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BankItemsColumns holds the columns for the "bank_items" table.
	BankItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "item_id", Type: field.TypeString, Unique: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "guild_id", Type: field.TypeInt, Nullable: true},
	}
	// BankItemsTable holds the schema information for the "bank_items" table.
	BankItemsTable = &schema.Table{
		Name:       "bank_items",
		Columns:    BankItemsColumns,
		PrimaryKey: []*schema.Column{BankItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bank_items_guilds_bank",
				Columns: []*schema.Column{BankItemsColumns[3]},

				RefColumns: []*schema.Column{GuildsColumns[0], GuildsColumns[1]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GuildsColumns holds the columns for the "guilds" table.
	GuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "id", Type: field.TypeInt, Unique: true, Increment: true},
	}
	// GuildsTable holds the schema information for the "guilds" table.
	GuildsTable = &schema.Table{
		Name:        "guilds",
		Columns:     GuildsColumns,
		PrimaryKey:  []*schema.Column{GuildsColumns[0], GuildsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GuildChannelsColumns holds the columns for the "guild_channels" table.
	GuildChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "channel", Type: field.TypeString},
		{Name: "guild_id", Type: field.TypeInt, Nullable: true},
	}
	// GuildChannelsTable holds the schema information for the "guild_channels" table.
	GuildChannelsTable = &schema.Table{
		Name:       "guild_channels",
		Columns:    GuildChannelsColumns,
		PrimaryKey: []*schema.Column{GuildChannelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "guild_channels_guilds_channels",
				Columns: []*schema.Column{GuildChannelsColumns[2]},

				RefColumns: []*schema.Column{GuildsColumns[0], GuildsColumns[1]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BankItemsTable,
		GuildsTable,
		GuildChannelsTable,
	}
)

func init() {
	BankItemsTable.ForeignKeys[0].RefTable = GuildsTable
	GuildChannelsTable.ForeignKeys[0].RefTable = GuildsTable
}
