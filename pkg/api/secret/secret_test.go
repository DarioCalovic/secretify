package secret_test

import (
	"testing"
	"time"

	"github.com/DarioCalovic/secretify"
	"github.com/DarioCalovic/secretify/pkg/api/secret"
	utildb "github.com/DarioCalovic/secretify/pkg/util/db"
	"github.com/DarioCalovic/secretify/pkg/util/mock"
	"github.com/DarioCalovic/secretify/pkg/util/mock/mockdb"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		email      string
		secret     string
		passphrase string
		expiresAt  time.Time
		revealOnce bool
	}
	cases := []struct {
		name     string
		args     args
		wantErr  bool
		wantData *secretify.Secret
		sdb      *mockdb.Secret
	}{{
		name:    "Fail on is lower role",
		wantErr: false,
		args: args{
			email:      "",
			secret:     "",
			passphrase: "",
			expiresAt:  mock.TestTime(2021),
			revealOnce: true,
		},
		sdb: &mockdb.Secret{
			CreateFn: func(db utildb.DB, secret secretify.Secret) (secretify.Secret, error) {
				var s secretify.Secret
				s.ID = 1
				s.Identifier = "generated-uuid"
				s.CreatedAt = mock.TestTime(2021)
				s.UpdatedAt = mock.TestTime(2021)
				return s, nil
			},
		},
	},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := secret.New(nil, tt.sdb, nil, nil)
			usr, err := s.Create(tt.args.email, tt.args.secret, tt.args.passphrase, tt.args.expiresAt, tt.args.revealOnce)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData != nil {
				assert.Equal(t, tt.wantData, usr)
			}
		})
	}
}
