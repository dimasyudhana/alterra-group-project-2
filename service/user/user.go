package user

import (
	"context"

	"github.com/dimasyudhana/alterra-group-project-2/entities"
)

type UserServiceInterface interface {
	Login(ctx context.Context, req entities.UserReqLogin) (error, int)
}
