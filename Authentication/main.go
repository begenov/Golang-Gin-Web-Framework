package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Username  string
	Email     string
	pswdHash  string
	CreatedAt string
	Active    string
	verHash   string
	timeout   string
}

var db *sql.DB

var store = sessions.NewCookieStore([]byte("super-secret"))

func init() {
	store.Options.HttpOnly = true
	store.Options.Secure = true
	gob.Register(&User{})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/home/orazali/Golang-Gin-Web-Framework/Authentication/templates/*html")
	log.Println("-----------------")
	var err error
	db, err = sql.Open("mysql", "task:pass@/task?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	authRouter := router.Group("/user", auth)
	router.GET("/", indexHandler)
	router.GET("/login", loginHandler)
	router.POST("/login", loginPOSThadnler)

	authRouter.GET("/profile", profileHandler)

	err = router.Run()
	if err != nil {
		panic(err)
	}

}

func auth(ctx *gin.Context) {
	session, _ := store.Get(ctx.Request, "session")
	fmt.Println("session ", session)
	_, ok := session.Values["user"]
	if !ok {
		ctx.HTML(http.StatusForbidden, "login.html", nil)
	}
	ctx.Next()

}
func indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)

}
func loginHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)

}
func loginPOSThadnler(ctx *gin.Context) {
	var user User
	user.Username = ctx.PostForm("username")
	password := ctx.PostForm("password")
	err := user.getUserByUsername()
	if err != nil {
		fmt.Println("error selecting pswd_hash in db by username, err:", err)
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "check username and password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.pswdHash), []byte(password))
	log.Println("err from bycrypt:", err)
	if err == nil {
		session, _ := store.Get(ctx.Request, "session")
		session.Values["user"] = user
		session.Values["user"] = user
		session.Save(ctx.Request, ctx.Writer)
		ctx.HTML(http.StatusOK, "loggedin.html", gin.H{
			"message": "check user and passwrod",
		})
		return
	}
	ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{
		"message": "Check username and password"})
}
func profileHandler(ctx *gin.Context) {
	session, _ := store.Get(ctx.Request, "session")
	var user = &User{}
	val := session.Values["user"]
	var ok bool
	if user, ok = val.(*User); !ok {
		ctx.HTML(http.StatusForbidden, "login.html", nil)
		return
	}
	ctx.HTML(http.StatusOK, "profile.html", gin.H{"user": user})
}

func (u *User) getUserByUsername() error {
	stmt := "SELECT * FROM user WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Email, &u.Username, u.pswdHash)
	if err != nil {
		log.Println("error get user by username")
		return err
	}
	return nil
}
