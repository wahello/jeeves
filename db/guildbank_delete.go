// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"

	"github.com/AlecAivazis/jeeves/db/guildbank"
	"github.com/AlecAivazis/jeeves/db/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GuildBankDelete is the builder for deleting a GuildBank entity.
type GuildBankDelete struct {
	config
	predicates []predicate.GuildBank
}

// Where adds a new predicate to the delete builder.
func (gbd *GuildBankDelete) Where(ps ...predicate.GuildBank) *GuildBankDelete {
	gbd.predicates = append(gbd.predicates, ps...)
	return gbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gbd *GuildBankDelete) Exec(ctx context.Context) (int, error) {
	return gbd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gbd *GuildBankDelete) ExecX(ctx context.Context) int {
	n, err := gbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gbd *GuildBankDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gbd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(guildbank.Table))
	for _, p := range gbd.predicates {
		p(selector)
	}
	query, args := builder.Delete(guildbank.Table).FromSelect(selector).Query()
	if err := gbd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// GuildBankDeleteOne is the builder for deleting a single GuildBank entity.
type GuildBankDeleteOne struct {
	gbd *GuildBankDelete
}

// Exec executes the deletion query.
func (gbdo *GuildBankDeleteOne) Exec(ctx context.Context) error {
	n, err := gbdo.gbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{guildbank.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gbdo *GuildBankDeleteOne) ExecX(ctx context.Context) {
	gbdo.gbd.ExecX(ctx)
}
