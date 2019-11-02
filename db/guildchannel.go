// Code generated by entc, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// GuildChannel is the model entity for the GuildChannel schema.
type GuildChannel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel string `json:"channel,omitempty"`
}

// FromRows scans the sql response data into GuildChannel.
func (gc *GuildChannel) FromRows(rows *sql.Rows) error {
	var vgc struct {
		ID      int
		Channel sql.NullString
	}
	// the order here should be the same as in the `guildchannel.Columns`.
	if err := rows.Scan(
		&vgc.ID,
		&vgc.Channel,
	); err != nil {
		return err
	}
	gc.ID = vgc.ID
	gc.Channel = vgc.Channel.String
	return nil
}

// QueryGuild queries the guild edge of the GuildChannel.
func (gc *GuildChannel) QueryGuild() *GuildQuery {
	return (&GuildChannelClient{gc.config}).QueryGuild(gc)
}

// Update returns a builder for updating this GuildChannel.
// Note that, you need to call GuildChannel.Unwrap() before calling this method, if this GuildChannel
// was returned from a transaction, and the transaction was committed or rolled back.
func (gc *GuildChannel) Update() *GuildChannelUpdateOne {
	return (&GuildChannelClient{gc.config}).UpdateOne(gc)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gc *GuildChannel) Unwrap() *GuildChannel {
	tx, ok := gc.config.driver.(*txDriver)
	if !ok {
		panic("db: GuildChannel is not a transactional entity")
	}
	gc.config.driver = tx.drv
	return gc
}

// String implements the fmt.Stringer.
func (gc *GuildChannel) String() string {
	var builder strings.Builder
	builder.WriteString("GuildChannel(")
	builder.WriteString(fmt.Sprintf("id=%v", gc.ID))
	builder.WriteString(", channel=")
	builder.WriteString(gc.Channel)
	builder.WriteByte(')')
	return builder.String()
}

// GuildChannels is a parsable slice of GuildChannel.
type GuildChannels []*GuildChannel

// FromRows scans the sql response data into GuildChannels.
func (gc *GuildChannels) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vgc := &GuildChannel{}
		if err := vgc.FromRows(rows); err != nil {
			return err
		}
		*gc = append(*gc, vgc)
	}
	return nil
}

func (gc GuildChannels) config(cfg config) {
	for _i := range gc {
		gc[_i].config = cfg
	}
}
