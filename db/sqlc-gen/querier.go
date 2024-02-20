// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddWFNewInstances(ctx context.Context, arg AddWFNewInstancesParams) ([]AddWFNewInstancesRow, error)
	GetApp(ctx context.Context, arg GetAppParams) (string, error)
	GetClass(ctx context.Context, arg GetClassParams) (string, error)
	GetSchemaWithLock(ctx context.Context, arg GetSchemaWithLockParams) (GetSchemaWithLockRow, error)
	GetWFActiveStatus(ctx context.Context, arg GetWFActiveStatusParams) (pgtype.Bool, error)
	GetWFINstance(ctx context.Context, arg GetWFINstanceParams) (int64, error)
	GetWFInternalStatus(ctx context.Context, arg GetWFInternalStatusParams) (bool, error)
	GetWorkflow(ctx context.Context, step string) ([]GetWorkflowRow, error)
	RulesetRowLock(ctx context.Context, arg RulesetRowLockParams) (Ruleset, error)
	SchemaDelete(ctx context.Context, id int32) (int32, error)
	SchemaGet(ctx context.Context, arg SchemaGetParams) ([]SchemaGetRow, error)
	SchemaList(ctx context.Context) ([]SchemaListRow, error)
	SchemaListByApp(ctx context.Context, app string) ([]SchemaListByAppRow, error)
	SchemaListByAppAndClass(ctx context.Context, arg SchemaListByAppAndClassParams) ([]SchemaListByAppAndClassRow, error)
	SchemaListByAppAndSlice(ctx context.Context, arg SchemaListByAppAndSliceParams) ([]SchemaListByAppAndSliceRow, error)
	SchemaListByClass(ctx context.Context, class string) ([]SchemaListByClassRow, error)
	SchemaListByClassAndSlice(ctx context.Context, arg SchemaListByClassAndSliceParams) ([]SchemaListByClassAndSliceRow, error)
	SchemaListBySlice(ctx context.Context, slice int32) ([]SchemaListBySliceRow, error)
	SchemaNew(ctx context.Context, arg SchemaNewParams) (int32, error)
	SchemaUpdate(ctx context.Context, arg SchemaUpdateParams) error
	WfPatternSchemaGet(ctx context.Context, arg WfPatternSchemaGetParams) ([]byte, error)
	WfSchemaGet(ctx context.Context, arg WfSchemaGetParams) (Schema, error)
	WfSchemaList(ctx context.Context, arg WfSchemaListParams) ([]WfSchemaListRow, error)
	Wfschemadelete(ctx context.Context, arg WfschemadeleteParams) error
	Wfschemaget(ctx context.Context, arg WfschemagetParams) (WfschemagetRow, error)
	WorkFlowNew(ctx context.Context, arg WorkFlowNewParams) error
	WorkFlowUpdate(ctx context.Context, arg WorkFlowUpdateParams) (pgconn.CommandTag, error)
	WorkflowDelete(ctx context.Context, arg WorkflowDeleteParams) (pgconn.CommandTag, error)
	WorkflowList(ctx context.Context, arg WorkflowListParams) ([]WorkflowListRow, error)
	Workflowget(ctx context.Context, arg WorkflowgetParams) (WorkflowgetRow, error)
}

var _ Querier = (*Queries)(nil)
