package action_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"

	"go-get-it/route"
	"go-get-it/action"
)


const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	route.Routes()
}

func TestSaveAction(t * testing.T){
	resource := "/save"
	statusCode := 200

	t.Log("Given the /save resource.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", resource, statusCode)
		{
			mcPostBody := action.SaveRequest{}
			data, _ := json.Marshal(mcPostBody)

			req, err := http.NewRequest("POST", resource, bytes.NewReader(data))

			if err != nil {
				t.Fatalf("\t\tShould be able to make the Get call.\t %s", err, ballotX)
			}
			t.Logf("\t\tShould be able to make the Get call. %s", checkMark)

			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)

			if rw.Code != 200 {
				t.Fatal("\t\tShould receive \"200\"", ballotX, rw.Code)
			}
			t.Log("\t\tShould receive \"200\"", checkMark)
		}
	}
}
