package repositories

import (
	"database/sql"
	"social-network/src/models"
)

type posts struct {
	db *sql.DB
}

func NewRepositoryPosts(db *sql.DB) *posts {
	return &posts{db}
}

func (repositoryPosts posts) Create(post models.Post) (uint64, error) {
	statement, error := repositoryPosts.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(post.Title, post.Content, post.AuthorID)
	if error != nil {
		return 0, error
	}
	lastId, error := result.LastInsertId()
	if error != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

func (repositoryPosts posts) GetPost(postId uint64) (models.Post, error) {
	line, error := repositoryPosts.db.Query(`
		select p.*, u.id from posts p
			inner join users u on u.id = p.author_id
		where p.id = ?
		`,
		postId,
	)
	if error != nil {
		return models.Post{}, error
	}

	var post models.Post
	if line.Next() {
		if error := line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return models.Post{}, error
		}
	}
	return post, nil
}

func (repositoryPosts posts) ListPosts(userID uint64) ([]models.Post, error) {
	lines, error := repositoryPosts.db.Query(`
			select distinct p.*, u.nick from posts p
				inner join users u on u.id = p.author_id
				inner join followers f on f.user_id = p.author_id
			where u.id = ? or f.follower_id = ?
			order by 1 desc
		`,
		userID,
		userID,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		var post models.Post
		if error := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repositoryPosts posts) UpdatePost(postId uint64, post models.Post) error {
	statement, error := repositoryPosts.db.Prepare("update posts set title = ?, content = ? where author_id = ? and id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(post.Title, post.Content, post.AuthorID, postId); error != nil {
		return error
	}

	return nil
}

func (repositoryPosts posts) DeletePost(postId uint64) error {
	statement, error := repositoryPosts.db.Prepare("delete from posts where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(postId); error != nil {
		return error
	}

	return nil
}

func (repositoryPosts posts) GetPostsPerUser(userId uint64) ([]models.Post, error) {
	lines, error := repositoryPosts.db.Query(`
		select p.*, u.nick from posts p
			inner join users u on u.id = p.author_id
		where p.author_id = ?
		`,
		userId,
	)
	if error != nil {
		return nil, error
	}

	var posts []models.Post
	for lines.Next() {
		var post models.Post
		if error := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repositoryPosts posts) LikePost(postId uint64) error {
	statement, error := repositoryPosts.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(postId); error != nil {
		return error
	}

	return nil
}

func (repositoryPosts posts) UnlikePost(postId uint64) error {
	statement, error := repositoryPosts.db.Prepare(`
		update posts set likes = 
		case
			when likes > 0 then likes - 1 
			else 0
		end
		where id = ?
	`)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(postId); error != nil {
		return error
	}

	return nil
}
