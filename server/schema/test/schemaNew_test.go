package schema_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/remiges-tech/alya/wscutils"
	"github.com/remiges-tech/crux/server/schema"
	"github.com/remiges-tech/crux/testutils"
	"github.com/stretchr/testify/require"
)

func TestSchemaNew(t *testing.T) {
	testCases := schemaNewTestcase()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setting up buffer
			payload := bytes.NewBuffer(testutils.MarshalJson(tc.RequestPayload))

			res := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/wfschemaNew", payload)
			require.NoError(t, err)

			r.ServeHTTP(res, req)

			require.Equal(t, tc.ExpectedHttpCode, res.Code)
			if tc.ExpectedResult != nil {
				jsonData := testutils.MarshalJson(tc.ExpectedResult)
				require.JSONEq(t, string(jsonData), res.Body.String())
			} else {
				jsonData, err := testutils.ReadJsonFromFile(tc.TestJsonFile)
				require.NoError(t, err)
				require.JSONEq(t, string(jsonData), res.Body.String())
			}
		})
	}

}

func schemaNewTestcase() []testutils.TestCasesStruct {
	valTestJson, err := testutils.ReadJsonFromFile("./data/schema_new_validation_payload.json")
	if err != nil {
		log.Fatalln("Error reading JSON file:", err)
	}
	var valPayload schema.SchemaNewReq
	if err := json.Unmarshal(valTestJson, &valPayload); err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	cusValTestJson, err := testutils.ReadJsonFromFile("./data/schema_new_custom_validation_payload.json")
	if err != nil {
		log.Fatalln("Error reading JSON file:", err)
	}
	var cusValPayload schema.SchemaNewReq
	if err := json.Unmarshal(cusValTestJson, &cusValPayload); err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	successTestJson, err := testutils.ReadJsonFromFile("./data/schema_new_success_payload.json")
	if err != nil {
		log.Fatalln("Error reading JSON file:", err)
	}
	var successPayload schema.SchemaNewReq
	if err := json.Unmarshal(successTestJson, &successPayload); err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}

	schemaNewTestcase := []testutils.TestCasesStruct{
		{
			Name: "err- binding_json_error",
			RequestPayload: wscutils.Request{
				Data: nil,
			},

			ExpectedHttpCode: http.StatusBadRequest,
			ExpectedResult: &wscutils.Response{
				Status: wscutils.ErrorStatus,
				Data:   nil,
				Messages: []wscutils.ErrorMessage{
					{
						MsgID:   1001,
						ErrCode: wscutils.ErrcodeInvalidJson,
					},
				},
			},
		},
		{
			Name: "err- standard validation",
			RequestPayload: wscutils.Request{
				Data: valPayload,
			},

			ExpectedHttpCode: http.StatusBadRequest,
			TestJsonFile:     "./data/schema_new_validation_error.json",
		},
		{
			Name: "err- custom validation",
			RequestPayload: wscutils.Request{
				Data: cusValPayload,
			},

			ExpectedHttpCode: http.StatusBadRequest,
			TestJsonFile:     "./data/schema_new_custom_validation_error.json",
		},
		{
			Name: "Success- create schema ",
			RequestPayload: wscutils.Request{
				Data: successPayload,
			},

			ExpectedHttpCode: http.StatusOK,
			ExpectedResult: &wscutils.Response{
				Status:   wscutils.SuccessStatus,
				Data:     nil,
				Messages: nil,
			},
		},
	}
	return schemaNewTestcase
}
