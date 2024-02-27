package adapters

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"user-service/pkg/entities"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func NewPostgresAdapter(connStr string) (*PostgresAdapter, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresAdapter{DB: db}, nil
}

func (p *PostgresAdapter) Placeholder(ctx context.Context, placeholder string) error {
	//TODO implement me
	panic("implement me")
}
func (p *PostgresAdapter) CreateUser(ctx context.Context, id string, email string, name string, admin bool) error {
	query := `INSERT INTO users (user_id,user_email,user_name,admin,disabled) VALUES ($1, $2, $3, $4, false)`

	_, err := p.DB.ExecContext(ctx, query, id, email, name, admin)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresAdapter) GetUser(id string) (entities.User, error) {
	user := entities.User{}
	query := `SELECT * FROM users WHERE user_id = $1`
	row := p.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Admin, &user.Disabled)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *PostgresAdapter) GetUsers() ([]entities.User, error) {
	users := []entities.User{}
	query := `SELECT * FROM users`
	rows, err := p.DB.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := entities.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Admin, &user.Disabled)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (p *PostgresAdapter) AssignAdmin(ctx context.Context, id string) error {
	query := `UPDATE users SET admin = true WHERE user_id = $1`
	_, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (p *PostgresAdapter) UpdateUser(ctx context.Context, id string, email string, name string, admin bool, disabled bool) error {
	query := `UPDATE users SET user_email = $2, user_name = $3, admin = $4, disabled = $5 WHERE user_id = $1`
	_, err := p.DB.ExecContext(ctx, query, id, email, name, admin, disabled)
	if err != nil {
		return err
	}
	return nil
}
