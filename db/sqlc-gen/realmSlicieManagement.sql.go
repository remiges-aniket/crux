// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: realmSlicieManagement.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const copyConfig = `-- name: CopyConfig :execresult
INSERT INTO
    config (
        realm, slice, name, descr, val, ver, setby
    )
SELECT realm, $2, name, descr, val, ver, $3
FROM config
WHERE
    config.slice = $1
`

type CopyConfigParams struct {
	Slice   int32  `json:"slice"`
	Slice_2 int32  `json:"slice_2"`
	Setby   string `json:"setby"`
}

func (q *Queries) CopyConfig(ctx context.Context, arg CopyConfigParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, copyConfig, arg.Slice, arg.Slice_2, arg.Setby)
}

const copyRuleset = `-- name: CopyRuleset :execresult
INSERT INTO
    ruleset (
        realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdby
    )
SELECT
    realm,
    $2,
    app,
    brwf,
    class,
    setname,
    schemaid,
    is_active,
    is_internal,
    ruleset,
    $3
FROM ruleset
WHERE
    ruleset.slice = $1
    AND (
        $4::text [] is null
        OR app = any ($4::text [])
    )
`

type CopyRulesetParams struct {
	Slice     int32    `json:"slice"`
	Slice_2   int32    `json:"slice_2"`
	Createdby string   `json:"createdby"`
	App       []string `json:"app"`
}

func (q *Queries) CopyRuleset(ctx context.Context, arg CopyRulesetParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, copyRuleset,
		arg.Slice,
		arg.Slice_2,
		arg.Createdby,
		arg.App,
	)
}

const copySchema = `-- name: CopySchema :execresult
INSERT INTO
    schema (
        realm, slice, app, brwf, class, patternschema, actionschema, createdby
    )
SELECT
    realm,
    $2,
    app,
    brwf,
    class,
    patternschema,
    actionschema,
    $3
FROM schema
WHERE
    schema.slice = $1
    AND (
        $4::text [] is null
        OR app = any ($4::text [])
    )
`

type CopySchemaParams struct {
	Slice     int32    `json:"slice"`
	Slice_2   int32    `json:"slice_2"`
	Createdby string   `json:"createdby"`
	App       []string `json:"app"`
}

func (q *Queries) CopySchema(ctx context.Context, arg CopySchemaParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, copySchema,
		arg.Slice,
		arg.Slice_2,
		arg.Createdby,
		arg.App,
	)
}

const createNewSliceBY = `-- name: CreateNewSliceBY :one
INSERT INTO
    realmslice (
        realm, descr, active, activateat, deactivateat
    )
SELECT
    realm,
    (
        $3::VARCHAR is null
        OR descr = $3::VARCHAR
    ),
    -- descr,
    active,
    activateat,
    deactivateat
FROM realmslice
WHERE
    realmslice.id = $1
    AND realmslice.realm = $2
RETURNING
    realmslice.id
`

type CreateNewSliceBYParams struct {
	ID    int32  `json:"id"`
	Realm string `json:"realm"`
	Descr string `json:"descr"`
}

func (q *Queries) CreateNewSliceBY(ctx context.Context, arg CreateNewSliceBYParams) (int32, error) {
	row := q.db.QueryRow(ctx, createNewSliceBY, arg.ID, arg.Realm, arg.Descr)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createRealmSlice = `-- name: CreateRealmSlice :one
INSERT INTO
    realmslice (
        realm, descr, active
    )
VALUES (
    $1, $2, TRUE
)
RETURNING
   realmslice.id
`

type CreateRealmSliceParams struct {
	Realm string `json:"realm"`
	Descr string `json:"descr"`
}

func (q *Queries) CreateRealmSlice(ctx context.Context, arg CreateRealmSliceParams) (int32, error) {
	row := q.db.QueryRow(ctx, createRealmSlice, arg.Realm, arg.Descr)
	var id int32
	err := row.Scan(&id)
	return id, err
}
