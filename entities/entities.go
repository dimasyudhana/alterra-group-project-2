package entities

type (
	User struct {
		Id       uint   `json:"-" gorm:"type:int;primaryKey;autoIncrement"`
		Name     string `json:"name" gorm:"type:varchar(50);not null"`
		Email    string `json:"email" gorm:"type:varchar(50);not null"`
		Password string `json:"password" gorm:"type:varchar(80);not null"`
		Image    string `json:"image" gorm:"type:varchar(30);not null"`
		Address  string `json:"address" gorm:"type:varchar(255);not null"`
	}
	UserReqLogin struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	UserReqRegister struct {
		Name     string `form:"name" validate:"required"`
		Email    string `form:"email" validate:"required"`
		Password string `form:"password" validate:"required"`
		Image    string
		Address  string `form:"address" validate:"required"`
	}
	UserReqUpdate struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Image    string `json:"image" `
		Address  string `json:"address" `
	}
	WebResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}
)
