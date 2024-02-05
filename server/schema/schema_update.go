package schema

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/remiges-tech/alya/service"
	"github.com/remiges-tech/alya/wscutils"
	"github.com/remiges-tech/crux/db/sqlc-gen"
	"github.com/remiges-tech/logharbour/logharbour"
)

func SchemaUpdate(c *gin.Context, s *service.Service) {
	l := s.LogHarbour
	l.Log("Starting execution of SchemaUpdate()")
	editedBy := "admin"

	var sh Schema

	// check the capgrant table to see if the calling user has the capability to perform the
	// operation
	// isCapable, _ := utils.Authz_check(types.OpReq{
	// 	User:      username,
	// 	CapNeeded: []string{"schema"},
	// }, false)

	// if !isCapable {
	// 	l.Log("Unauthorized user:")
	// 	wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(utils.ErrUnauthorized))
	// 	return
	// }

	// The system will check whether there are any rulesets in the ruleset table whose
	// (slice,app,class) match this record. If this is true, then the call will fail.
	// In other words, updating a schema is not allowed once rulesets referring to it are defined.

	err := wscutils.BindJSON(c, &sh)
	if err != nil {
		l.LogActivity("Error Unmarshalling Query paramaeters to struct:", err.Error())
		return
	}

	// Validate request
	validationErrors := wscutils.WscValidate(sh, func(err validator.FieldError) []string { return []string{} })
	customValidationErrors := customValidationErrors(sh)
	validationErrors = append(validationErrors, customValidationErrors...)
	if len(validationErrors) > 0 {
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, validationErrors))
		return
	}

	query, ok := s.Database.(*sqlc.Queries)
	if !ok {
		l.Log("Error while getting query instance from service")
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(wscutils.ErrcodeDatabaseError))
		return
	}
	patternSchema, err := json.Marshal(sh.PatternSchema)
	if err != nil {
		patternSchema := "patternSchema"
		l.LogDebug("Error while marshaling patternSchema", err)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(wscutils.ErrcodeInvalidJson, &patternSchema)}))
		return
	}

	actionSchema, err := json.Marshal(sh.ActionSchema)
	if err != nil {
		actionSchema := "actionSchema"
		l.LogDebug("Error while marshaling actionSchema", err)
		wscutils.SendErrorResponse(c, wscutils.NewResponse(wscutils.ErrorStatus, nil, []wscutils.ErrorMessage{wscutils.BuildErrorMessage(wscutils.ErrcodeInvalidJson, &actionSchema)}))
		return
	}
	_, err = query.SchemaUpdate(c, sqlc.SchemaUpdateParams{Slice: sh.Slice, Class: sh.Class, App: sh.App, Patternschema: patternSchema, Actionschema: actionSchema, Editedby: editedBy})
	if err != nil {
		l.LogActivity("Error while Updating schema", err.Error())
		wscutils.SendErrorResponse(c, wscutils.NewErrorResponse(wscutils.ErrcodeDatabaseError))
		return
	}
	l.LogDataChange("Updated schema", logharbour.ChangeInfo{
		Entity:    "schema",
		Operation: "Update",
		Changes: []logharbour.ChangeDetail{
			{
				Field:    "patternSchema",
				OldValue: "",
				NewValue: patternSchema},
			{
				Field:    "actionSchema",
				OldValue: "",
				NewValue: actionSchema},
		},
	})
	wscutils.SendSuccessResponse(c, &wscutils.Response{Status: wscutils.SuccessStatus, Data: "updated successfully", Messages: nil})
	l.Log("Starting execution of SchemaUpdate()")
}
