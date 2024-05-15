package users_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/wildegor/kaspi-rest/internal/db/postgres"
	"time"
)

var _ IUserRepository = (*UsersRepository)(nil)

type UsersRepository struct {
	conn *postgres.PostgresConnection
}

func NewUsersRepository(conn *postgres.PostgresConnection) *UsersRepository {
	return &UsersRepository{conn: conn}
}

func (u UsersRepository) CreateIfNotExists(data *CreateUserModel) (*UsersModel, error) {
	query := fmt.Sprintf(`INSERT INTO %s (iin, phone, first_name, last_name, patronymic) VALUES (@iin, @phone, @first_name, @last_name, @patronymic) ON CONFLICT (iin) WHERE @iin DO NOTHING RETURNING id;`, UsersTable)

	first, last, patro := data.GetNames()

	entity := &UsersModel{
		IIN:        data.IIN,
		Phone:      data.Phone,
		FirstName:  first,
		LastName:   last,
		Patronymic: patro,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := u.conn.DB.QueryRow(context.TODO(), query, pgx.NamedArgs{
		"iin":        entity.IIN,
		"phone":      entity.Phone,
		"first_name": entity.FirstName,
		"last_name":  entity.LastName,
		"patronymic": entity.Patronymic,
	}).Scan(&entity.ID)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u UsersRepository) FindByIIN(iin string) (*UsersModel, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE iin LIKE '@iin%' LIMIT 1`, UsersTable)

	rows, err := u.conn.DB.Query(context.TODO(), query, pgx.NamedArgs{
		"iin": iin,
	})
	if err != nil {
		return nil, err
	}

	collectRows, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UsersModel])
	if err != nil {
		return nil, err
	}

	if len(collectRows) == 0 {
		return nil, pgx.ErrNoRows
	}

	return &collectRows[0], nil
}

func (u UsersRepository) FindByKey(key string) ([]UsersModel, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE fullname LIKE '%@key%' LIMIT 100`, UsersTable)

	rows, err := u.conn.DB.Query(context.TODO(), query, pgx.NamedArgs{
		"key": key,
	})
	if err != nil {
		return nil, err
	}

	collectRows, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UsersModel])
	if err != nil {
		return nil, err
	}

	if len(collectRows) == 0 {
		return nil, pgx.ErrNoRows
	}

	return collectRows, nil
}
