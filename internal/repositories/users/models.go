package users_repository

import (
	"fmt"
	"strings"
	"time"
)

const UsersTable = "kaspi.users"

type CreateUserModel struct {
	IIN        string `json:"iin" db:"iin"`
	Sex        int8   `json:"sex" db:"s"`
	BirthYear  int    `json:"b_year" db:"b_year"`
	BirthDay   int    `json:"b_day" db:"b_day"`
	BirthMonth int    `json:"b_month" db:"b_month"`
	Phone      string `json:"phone" db:"phone"`
	Name       string `json:"name" db:"name"`
}

func (m *CreateUserModel) GetNames() (first, last, patro string) {
	parts := strings.Split(m.Name, " ")
	if len(parts) >= 1 {
		first = parts[0]
	}
	if len(parts) >= 2 {
		last = parts[len(parts)-1]
	}
	if len(parts) >= 3 {
		patro = parts[len(parts)-2]
	}
	return first, last, patro
}

type UsersModel struct {
	ID         int64     `json:"id" db:"id"`
	IIN        string    `json:"iin" db:"iin"`
	Sex        int8      `json:"sex" db:"s"`
	BirthYear  int       `json:"b_year" db:"b_year"`
	BirthDay   int       `json:"b_day" db:"b_day"`
	BirthMonth int       `json:"b_month" db:"b_month"`
	Phone      string    `json:"phone" db:"phone"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (m *UsersModel) Name() string {
	return fmt.Sprintf("%s %s %s", m.LastName, m.FirstName, m.Patronymic)
}

type IUserRepository interface {
	CreateIfNotExists(data *CreateUserModel) (*UsersModel, error)
	FindByIIN(iin string) (*UsersModel, error)
	FindByKey(key string) ([]UsersModel, error)
}
