package dbapi

import (
	"database/sql"
	"log"
	"os"

	"github.com/VladRomanciuc/Go-classes/api/models"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct{}

func NewSQLiteDb() models.DbOps {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable := `
	create table posts (id integer not null primary key, title text, txt text);
	delete from posts;
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Printf("%q: %s\n", err, createTable)
	}
	return &sqlite{}
}

func (*sqlite) AddPost(post *models.Post) (*models.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	prep, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	addpost, err := prep.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer addpost.Close()

	_, err = addpost.Exec(post.Id, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	prep.Commit()
	return post, nil
}

func (*sqlite) GetAll() ([]models.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	entry, err := db.Query("select id, title, txt from posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer entry.Close()

	var posts []models.Post
	for entry.Next() {
		var id string
		var title string
		var text string
		err = entry.Scan(&id, &title, &text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		post := models.Post{
			Id:    id,
			Title: title,
			Text:  text,
		}
		posts = append(posts, post)
	}
	err = entry.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}

func (*sqlite) DeleteById(id string) (*models.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	prep, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	delete, err := prep.Prepare("delete from posts where id = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer delete.Close()

	_, err = delete.Exec(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	prep.Commit()

	return nil, nil
}


func (*sqlite) GetById(id string) (*models.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	row := db.QueryRow("select id, title, txt from posts where id = ?", id)

	var post models.Post
	if row != nil {
		var id string
		var title string
		var text string
		err := row.Scan(&id, &title, &text)
		if err != nil {
			return nil, err
		} else {
			post = models.Post{
				Id:    id,
				Title: title,
				Text:  text,
			}
		}
	}

	return &post, nil

}