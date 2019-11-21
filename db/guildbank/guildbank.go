// Code generated by entc, DO NOT EDIT.

package guildbank

const (
	// Label holds the string label denoting the guildbank type in the database.
	Label = "guild_bank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChannelID holds the string denoting the channelid vertex property in the database.
	FieldChannelID = "channel_id"
	// FieldDisplayMessageID holds the string denoting the displaymessageid vertex property in the database.
	FieldDisplayMessageID = "display_message_id"
	// FieldBalance holds the string denoting the balance vertex property in the database.
	FieldBalance = "balance"

	// Table holds the table name of the guildbank in the database.
	Table = "guild_banks"
	// ItemsTable is the table the holds the items relation/edge.
	ItemsTable = "bank_items"
	// ItemsInverseTable is the table name for the BankItem entity.
	// It exists in this package in order to avoid circular dependency with the "bankitem" package.
	ItemsInverseTable = "bank_items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "bank_id"
	// GuildTable is the table the holds the guild relation/edge.
	GuildTable = "guild_banks"
	// GuildInverseTable is the table name for the Guild entity.
	// It exists in this package in order to avoid circular dependency with the "guild" package.
	GuildInverseTable = "guilds"
	// GuildColumn is the table column denoting the guild relation/edge.
	GuildColumn = "guild_id"
)

// Columns holds all SQL columns are guildbank fields.
var Columns = []string{
	FieldID,
	FieldChannelID,
	FieldDisplayMessageID,
	FieldBalance,
}
