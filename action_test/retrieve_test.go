package action_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

)



func TestRetrieveAction(t * testing.T){
	resource := "/retrieve"
	statusCode := 200

	t.Log("Given the /retrieve resource.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", resource, statusCode)
		{
			req, err := http.NewRequest("GET", resource, nil)
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
