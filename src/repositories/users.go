package repositories

import (
	"database/sql"
	"fmt"
	"social-network/src/models"
)

type users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *users {
	return &users{db}
}

func (repositoryUser users) Create(user models.User) (uint64, error) {
	statement, error := repositoryUser.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastIDInserted, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastIDInserted), nil
}

func (repositoryUser users) SearchUsers(filter string) ([]models.User, error) {
	filter = fmt.Sprintf("%%%s%%", filter)

	lines, error := repositoryUser.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		filter,
		filter,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if error := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}

func (repositoryUser users) GetUser(ID uint64) (models.User, error) {
	line, error := repositoryUser.db.Query(
		"select id, name, nick, email, created_at from users where ID = ?",
		ID,
	)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error := line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}
	return user, nil
}

func (repositoryUser users) UpdateUser(ID uint64, user models.User) error {
	statement, error := repositoryUser.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}

func (repositoryUser users) DeleteUser(ID uint64) error {
	statement, error := repositoryUser.db.Prepare("delete from users where id =?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

func (repositoryUser users) GetUserForEmail(email string) (models.User, error) {
	line, error := repositoryUser.db.Query("select id, password from users where email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error := line.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}
	return user, nil
}

func (repositoryUser users) FollowUser(userId, followerId uint64) error {
	statement, error := repositoryUser.db.Prepare("insert ignore into followers(user_id, follower_id) values(?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(userId, followerId); error != nil {
		return error
	}

	return nil
}

func (repositoryUser users) UnfollowUser(userId, followerId uint64) error {
	statement, error := repositoryUser.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(userId, followerId); error != nil {
		return error
	}

	return nil
}

func (repositoryUser users) GetFollowers(userId uint64) ([]models.User, error) {
	lines, error := repositoryUser.db.Query(`
		select u.id, u.name, u.nick, u.email from followers f
			inner join users u on u.id = f.follower_id
		where user_id = ?`,
		userId,
	)
	if error != nil {
		return nil, error
	}

	var followers []models.User
	for lines.Next() {
		var follower models.User
		if error := lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
		); error != nil {
			return nil, error
		}
		followers = append(followers, follower)
	}
	return followers, nil
}

func (repositoryUser users) GetFollowing(userId uint64) ([]models.User, error) {
	lines, error := repositoryUser.db.Query(`
			select u.id, u.name, u.nick, u.email from followers f
				inner join users u on u.id = f.user_id
			where f.follower_id = ?
		`,
		userId,
	)
	if error != nil {
		return nil, error
	}

	var users []models.User
	for lines.Next() {
		var user models.User
		if error := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}
	return users, nil
}

func (repositoryUser users) GetPassword(userId uint64) (string, error) {
	line, error := repositoryUser.db.Query("select password from users where id = ?", userId)
	if error != nil {
		return "", error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error := line.Scan(&user.Password); error != nil {
			return "", error
		}
	}
	return user.Password, nil
}

func (repositoryUser users) UpdatePassword(userId uint64, newPassword string) error {
	statement, error := repositoryUser.db.Prepare("update users set password = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(newPassword, userId); error != nil {
		return error
	}

	return nil
}
