package realmSliceManagement

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/remiges-tech/alya/service"
	"github.com/remiges-tech/alya/wscutils"
	"github.com/remiges-tech/crux/db"
	"github.com/remiges-tech/crux/db/sqlc-gen"
	"github.com/remiges-tech/crux/server"
	"github.com/remiges-tech/crux/types"
)

// var (
// 	userID    = "1234"
// 	capForNew = []string{"root"}
// 	realmName = "NSE"
// 	// realmID   = int32(11)
// )

type RealmSliceDeactivateReq struct {
	// Id is refer to `realmslice_id` in db
	Id int32 `json:"id" validate:"required,gt=0"`
	//`from` is an optional timestamp parameter specifying
	// from when the slice will be activated. This timestamp
	// must be in the future.
	From *time.Time `json:"from,omitempty"`
}

func RealmSliceDeactivate(c *gin.Context, s *service.Service) {
	l := s.LogHarbour
	l.Debug0().Log("starting execution of RealmSliceDeactivate()")

	isCapable, _ := server.Authz_check(types.OpReq{
		User:      userID,
		CapNeeded: capForNew,
	}, false)

	if !isCapable {
		l.Info().LogActivity("unauthorized user:", userID)
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Unauthorized, server.ErrCode_Unauthorized))
		return
	}

	var (
		isActive bool
		fromt    time.Time
		req      RealmSliceDeactivateReq
	)

	err := wscutils.BindJSON(c, &req)
	if err != nil {
		l.Error(err).Log("error unmarshalling query parameters to struct:")
		return
	}

	// Validate request
	validationErrors := wscutils.WscValidate(req, func(err validator.FieldError) []string { return []string{} })
	if len(validationErrors) > 0 {
		l.Debug0().LogDebug("standard validation errors", validationErrors)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}

	if req.From != nil {
		if req.From.Before(time.Now()) {
			l.Debug0().LogDebug("givenfrom time is too early:", req.From)
			wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_Unauthorized, server.ErrCode_TooEarly))
			return
		} else {
			isActive = true
		}
		fromt = *req.From
	}

	query, ok := s.Dependencies["queries"].(*sqlc.Queries)
	if !ok {
		l.Info().Log("error while getting query instance from service dependencies")
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(server.MsgId_InternalErr, server.ErrCode_Internal))
		return
	}

	newSliceID, err := query.RealmSliceDeactivate(c, sqlc.RealmSliceDeactivateParams{
		ID:           req.Id,
		Isactive:     isActive,
		Deactivateat: pgtype.Timestamp{Time: fromt, Valid: req.From != nil},
	})
	if err != nil {
		l.Info().Error(err).Log("error while changing active status with func RealmSliceDeactivate")
		errmsg := db.HandleDatabaseError(err)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{errmsg}))
		return
	}
	wscutils.SendSuccessResponse(c, &wscutils.Response{Status: wscutils.SuccessStatus, Data: newSliceID, Messages: nil})
	l.Debug0().Log("exiting from RealmSliceDeactivate()")
	return
}
