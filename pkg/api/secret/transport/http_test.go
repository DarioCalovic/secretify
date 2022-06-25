package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	"github.com/DarioCalovic/secretify/pkg/api/setting"
	utilconfig "github.com/DarioCalovic/secretify/pkg/util/config"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/mock"
	"github.com/DarioCalovic/secretify/pkg/util/mock/mockdb"
	"github.com/DarioCalovic/secretify/pkg/util/server"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *secretify.HTTPOKResponse
		sdb        *mockdb.Secret
		cfg        *utilconfig.Configuration
	}{
		{
			name:       "Fail on cipher not passed",
			req:        `{"no": "cipher"}`,
			wantStatus: http.StatusNotAcceptable,
			cfg:        utilconfig.NewConfiguration(),
		},
		{
			name:       "Fail on expires_at not passed",
			req:        `{"cipher": "cipher"}`,
			wantStatus: http.StatusNotAcceptable,
			cfg:        utilconfig.NewConfiguration(),
		},
		{
			name:       "Fail on wrongly expires_at passed",
			req:        `{"cipher": "cipher","expires_at": "10d"}`,
			wantStatus: http.StatusNotAcceptable,
			cfg:        utilconfig.NewConfiguration(),
		},
		{
			name:       "Fail on passphrase not passed",
			req:        `{"cipher": "cipher","expires_at": "10h"}`,
			wantStatus: http.StatusNotAcceptable,
			cfg: &utilconfig.Configuration{
				Policy: &utilconfig.Policy{
					Passphrase: struct {
						Required bool "json:\"required\""
					}{true},
				},
			},
		},
		{
			name:       "Success",
			req:        `{"cipher": "cipher","expires_at": "10h"}`,
			wantStatus: http.StatusOK,
			cfg: &utilconfig.Configuration{
				Policy: &utilconfig.Policy{
					Identifier: struct {
						Size int `json:"size"`
					}{18},
				},
			},
			sdb: &mockdb.Secret{
				CreateFn: func(db utildb.DB, secret secretify.Secret) (secretify.Secret, error) {
					var s secretify.Secret
					s.ID = 1
					s.Identifier = "generated-identifier"
					s.CreatedAt = mock.TestTime(2021)
					s.UpdatedAt = mock.TestTime(2021)
					return s, nil
				},
			},
			wantResp: &secretify.HTTPOKResponse{
				Data: createRes{
					Identifier: "generated-identifier",
					CreatedAt:  mock.TestTime(2021),
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			NewHTTP(secret.New(nil, tt.sdb, setting.New(tt.cfg), nil), r)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/secret"
			res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := &secretify.HTTPOKResponse{
					Data: createRes{},
				}
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				dbByte, _ := json.Marshal(response.Data)
				var data createRes
				_ = json.Unmarshal(dbByte, &data)
				assert.Equal(t, tt.wantResp.Data, data)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}
