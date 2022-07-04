package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"imdbv2/ent"
	"imdbv2/ent/user"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
)

var PasswordReset struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `gorm:"unique"`
}

func Forgot(c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		buf, err1 := io.ReadAll(r.Body)
		fmt.Println(err1, string(buf))

		var Forgot struct {
			Email string `json:"email"`
		}

		er := json.Unmarshal(buf, &Forgot)
		if er != nil {
			log.Fatal(er)
		}

		fmt.Println("email given", Forgot.Email)

		// need to check if there is a user with this mail
		userMail, err2 := c.User.Query().Where(user.Email(Forgot.Email)).All(r.Context())
		if err2 != nil {
			log.Fatal("error while querying user by email", err2)
		}

		if userMail == nil {
			fmt.Println("email not found")
			return
		}

		token := RandStringRunes(16)

		from := "Admin@imdb_clone.com"
		password := "123456789"

		to := []string{
			Forgot.Email,
		}

		host := "smtp.gmail.com"
		port := "587"

		url := "http://localhost:3000/reset/" + token

		message := []byte("click <a href=\"" + url + "\">here</a> to reset your password!")

		err3 := smtp.SendMail(host+":"+port, password, from, to, message)

		if err3 != nil {
			log.Println(err3)
			return
		}

		// writing the response for successful sending email process
		res2, e := w.Write([]byte("Check your email"))
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("res : ", res2)

	})
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzACDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func resetHandler(c *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		buf, err := io.ReadAll(r.Body)
		fmt.Println(err, string(buf))

		var userData struct {
			GivenPassword        string `json:"GivenPassword"`
			GivenPasswordConfirm string `json:"GivenPasswordConfirm"`
			GivenID              int    `json:"GivenID"`
		}

		err = json.Unmarshal(buf, &userData)
		if err != nil {
			log.Fatal(err)
		}

		if userData.GivenPassword != userData.GivenPasswordConfirm {
			fmt.Println("Passwords do not match!")
			return
		}

		bcrypedPassword, _ := bcrypt.GenerateFromPassword([]byte(userData.GivenPassword), 14)

		newUser := c.User.
			Update().Where(user.ID(userData.GivenID)).
			SetPassword(string(bcrypedPassword)).
			SaveX(r.Context())
		fmt.Println("new user added:", newUser)

	})
}