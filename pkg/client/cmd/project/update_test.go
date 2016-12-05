package project_test

import (
	"encoding/json"
	"github.com/lastbackend/lastbackend/libs/db"
	h "github.com/lastbackend/lastbackend/libs/http"
	"github.com/lastbackend/lastbackend/libs/model"
	"github.com/lastbackend/lastbackend/pkg/client/cmd/project"
	"github.com/lastbackend/lastbackend/pkg/client/context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	var (
		err error
		ctx = context.Mock()
	)

	const token = "mocktoken"

	ctx.Token = token

	//------------------------------------------------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tk := r.Header.Get("Authorization")
		assert.NotEmpty(t, tk, "token should be not empty")
		assert.Equal(t, tk, "Bearer "+token, "they should be equal")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		assert.NotEmpty(t, body, "body should not be empty")
		db_project := new(model.Project)
		reqw_project := new(model.Project)
		err = json.Unmarshal(body, &db_project)

		if err != nil {
			t.Error(err)
			return
		}

		data, err := db.Init()

		if err != nil {
			t.Error(err)
			return
		}

		defer data.Close()

		err = data.Get("project", &db_project)

		if err != nil {
			t.Error(err)
			return
		}

		if reqw_project.Name != "" {
			db_project.Name = reqw_project.Name
		}

		db_project.Description = reqw_project.Description
		db_project.Updated = time.Now()
		err = data.Set("project", db_project)

		if err != nil {
			t.Error(err)
			return
		}

		w.WriteHeader(200)
		_, err = w.Write(nil)

		if err != nil {
			t.Error(err)
			return
		}
	}))
	defer server.Close()
	//------------------------------------------------------------------------------------------

	ctx.HTTP = h.New(server.URL)
	err = project.Update("mock_name", "mock desc")

	if err != nil {
		t.Error(err)
		return
	}

	return

}