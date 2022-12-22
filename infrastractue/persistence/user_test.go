package persistence

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/domain/entity"
)

func TestUserRepo_GetUser(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		want    *entity.User
		wantErr error
	}{
		{
			name:   "ok",
			userID: "1",
			want: &entity.User{
				ID:   "1",
				Name: "user1",
			},
			wantErr: nil,
		},
		{
			name:   "ok",
			userID: "2",
			want: &entity.User{
				ID:   "2",
				Name: "user2",
			},
			wantErr: nil,
		},
		{
			name:    "notExistUser",
			userID:  "999",
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userRepo.GetUser(context.Background(), tt.userID)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

		})
	}

}
