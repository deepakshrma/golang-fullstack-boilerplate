package repo

import (
	"context"
	"log/slog"
	"webapp/pkg/model"
)

func (m *PostgresDBRepo) AllTodos() ([]*model.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select t.id, t.title, t.description, t.done,  t.created_at, t.updated_at, u.id
		from  todos t left join users u on (t.user_id = u.id)
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.CreatedBy.ID,
		)
		if err != nil {
			return nil, err
		}

		userQuery := `select id, email, first_name, last_name from users where id = $1`
		var user model.User
		err = m.DB.QueryRowContext(ctx, userQuery, todo.CreatedBy.ID).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			slog.Warn("error while scanning user", "userid", user.ID)
		}
		todo.CreatedBy = user
		todos = append(todos, &todo)
	}

	return todos, nil
}
