// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"

	"github.com/AlecAivazis/jeeves/db/guildchannel"
	"github.com/AlecAivazis/jeeves/db/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GuildChannelDelete is the builder for deleting a GuildChannel entity.
type GuildChannelDelete struct {
	config
	predicates []predicate.GuildChannel
}

// Where adds a new predicate to the delete builder.
func (gcd *GuildChannelDelete) Where(ps ...predicate.GuildChannel) *GuildChannelDelete {
	gcd.predicates = append(gcd.predicates, ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GuildChannelDelete) Exec(ctx context.Context) (int, error) {
	return gcd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GuildChannelDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GuildChannelDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gcd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(guildchannel.Table))
	for _, p := range gcd.predicates {
		p(selector)
	}
	query, args := builder.Delete(guildchannel.Table).FromSelect(selector).Query()
	if err := gcd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// GuildChannelDeleteOne is the builder for deleting a single GuildChannel entity.
type GuildChannelDeleteOne struct {
	gcd *GuildChannelDelete
}

// Exec executes the deletion query.
func (gcdo *GuildChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{guildchannel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GuildChannelDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}
