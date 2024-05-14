// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: stepworkflow.sql

package sqlc

import (
	"context"
)

const getWorkflowNameForStep = `-- name: GetWorkflowNameForStep :one
SELECT workflow,step FROM stepworkflow WHERE step = $1
`

type GetWorkflowNameForStepRow struct {
	Workflow string `json:"workflow"`
	Step     string `json:"step"`
}

func (q *Queries) GetWorkflowNameForStep(ctx context.Context, step string) (GetWorkflowNameForStepRow, error) {
	row := q.db.QueryRow(ctx, getWorkflowNameForStep, step)
	var i GetWorkflowNameForStepRow
	err := row.Scan(&i.Workflow, &i.Step)
	return i, err
}