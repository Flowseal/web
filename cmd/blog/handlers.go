package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"encoding/base64"
	"os"
	"time"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
)

type featuredPostData struct {
	PostId      string `db:"post_id"`
	Tag    			string `db:"tag"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"img_modifier"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
	URLTitle    string
}

type mostRecentPostData struct {
	PostId      string `db:"post_id"`
	Img         string `db:"img"`
	ImgAlt      string `db:"img_alt"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
	URLTitle    string
}

type indexPage struct {
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type contentPage struct {
	PostId     string `db:"post_id"`
	Img        string `db:"img"`
	ImgAlt     string `db:"img_alt"`
	Title      string `db:"title"`
	Subtitle   string `db:"subtitle"`
	Content    string `db:"content"`
	Paragraphs []string
}

type publishPostRequest struct {
	Img           string `json:"card-image"`
	ImgName       string `json:"card-image-file-name"`
	Title         string `json:"title"`
	Subtitle      string `json:"description"`
	Author        string `json:"author-name"`
	AuthorImg     string `json:"author-photo"`
	AuthorImgName string `json:"author-photo-file-name"`
	PublishDate   string `json:"date"`
	Content       string `json:"content"`
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
			author_img,
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
			img,
			img_alt,
			title,
			subtitle,
			author,
			author_img,
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

	for index, post := range featuredPosts {
		featuredPosts[index].URLTitle = strings.ReplaceAll(post.Title, " ", "-")
	}
	
	mostRecentPosts, err := mostRecentPosts(db)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	for index, post := range mostRecentPosts {
		mostRecentPosts[index].URLTitle = strings.ReplaceAll(post.Title, " ", "-")
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

func newPost(db *sqlx.DB, req publishPostRequest) error {
	const query = `
       INSERT INTO
			 post
       (
					 img,
           img_alt,
					 title,
					 subtitle,
					 author,
					 author_img,
					 publish_date,
					 content
       )
       VALUES
       (
           ?,
           ?,
           ?,
           ?,
           ?,
           ?,
           ?,
           ?
       )
   `
	card_img_decoded, err := base64.StdEncoding.DecodeString(req.Img)
	if err != nil {
		return err
	}
	card_img, err := os.Create("static/img/" + req.ImgName)
	if err != nil {
		return err
	}
	_, err = card_img.Write(card_img_decoded)
	if err != nil {
		return err
	}
	author_img_decoded, err := base64.StdEncoding.DecodeString(req.AuthorImg)
	if err != nil {
		return err
	}
	author_img, err := os.Create("static/img/" + req.AuthorImgName)
	if err != nil {
		return err
	}
	_, err = author_img.Write(author_img_decoded)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, "static/img/"+req.ImgName, req.Title+"-preview", req.Title, req.Subtitle, req.Author, "static/img/"+req.AuthorImgName, req.PublishDate, req.Content)
	return err
}

func adminHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func loginHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func postHandler(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			post_id,
			title,
			subtitle,
			img,
			img_alt,
			content
		FROM
			post
		WHERE post_id = 
	`
	query += mux.Vars(r)["postId"]

	var content contentPage
	err := db.Get(&content, query)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	ts, err := template.ParseFiles("pages/post.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	content.Paragraphs = strings.Split(content.Content, "\n")

	err = ts.Execute(w, content)
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
		} else if strings.Contains(r.URL.Path, "/admin") {
			adminHandler(db, w, r)
		} else if strings.Contains(r.URL.Path, "/login") {
			loginHandler(db, w, r)
		} else if strings.Contains(r.URL.Path, "/post") {
			postHandler(db, w, r)
		} else {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
	}
}