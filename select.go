package main

import (
	"context"
	"fmt"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Select(pool *pgxpool.Pool) {
	query := `
select id,
    name,
    array(
        select sequence
        from content
        where food_id = food.id
    ) as sequences
from food
where id = any(@id)`
	parsedQuery, args, err := pgx.NamedArgs{"id": []int{1, 2}}.RewriteQuery(context.TODO(), nil, query, nil)
	if err != nil {
		log.Fatalf("NamedArgs(): err %s", err)
	}
	fmt.Println(parsedQuery)
	fmt.Println(args...)
	bind := []BindObject{}
	err = pgxscan.Select(context.TODO(), pool,
		&bind,
		parsedQuery,
		args...)
	if err != nil {
		log.Fatalf("Select(): err %s", err)
	}

	fmt.Println(bind)
}

type BindObject struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Sequenses []int  `db:"sequences"`
}
