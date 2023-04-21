package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"github.com/jmoiron/sqlx"
)

type featuredPostData struct {
	PostId      string `db:"post_id"`
	Tag    			string `db:"tag"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"img_modifier"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

type mostRecentPostData struct {
	PostId      string `db:"post_id"`
	Img         string `db:"img_url"`
	ImgAlt      string `db:"img_alt"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

type indexPage struct {
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			post_id,
			tag,
			title,
			subtitle,
			img_modifier,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 1
	`
	var posts []featuredPostData
	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
			post_id,
			img_url,
			img_alt,
			title,
			subtitle,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
	`
	var posts []mostRecentPostData
	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func homeHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	featuredPosts, err := featuredPosts(db)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	mostRecentPosts, err := mostRecentPosts(db)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := indexPage{
		FeaturedPosts:   featuredPosts,
		MostRecentPosts: mostRecentPosts,
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func catchAllHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/home" {
			homeHandler(db, w, r)
		} else if strings.Contains(r.URL.Path, "/post") {
			// postHandler(w, r)
			http.Redirect(w, r, "/home", http.StatusNotFound)
		} else {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
	}
}