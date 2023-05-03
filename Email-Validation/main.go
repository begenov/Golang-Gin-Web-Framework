package main

import (
	"fmt"
	"log"
	"net/http"

	emailverifer "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
)

var (
	verifier = emailverifer.NewVerifier()
)

func main() {

	verifier = verifier.EnableDomainSuggest()
	verifier = verifier.AddDisposableDomains([]string{"tractorjj.com"})

	router := gin.Default()

	router.LoadHTMLGlob("/home/orazali/Golang-Gin-Web-Framework/Email-Validation/templates/*.html")

	router.GET("/verifyemail", verEmailGetHandler)
	router.POST("/verifyemail", verEmailPostHandler)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

}

func verEmailGetHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ver-email.html", nil)

}

func verEmailPostHandler(ctx *gin.Context) {
	fmt.Println("verEemailPostHandler running")

	email := ctx.PostForm("email")
	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		ctx.HTML(http.StatusBadRequest, "ver-email.html", gin.H{
			"message": "unable to register email address, please try again",
		})
		return
	}

	fmt.Println("email validation result", ret)
	fmt.Println("Email:", ret.Email, "\nReachable:", ret.Reachable, "\nSyntax:", ret.Syntax, "\nSMTP", ret.SMTP, "\nGravatar:", ret.Gravatar)

	if ret.Syntax.Valid {
		fmt.Println("sorry we do not accept disposable email addresses")
		ctx.HTML(http.StatusBadRequest, "ver-email.html", gin.H{
			"message": "email address syntax is invalid",
		})
		return
	}
	if ret.Disposable {
		fmt.Println("sorry we do not accept disposable email addresses")
		ctx.HTML(http.StatusBadRequest, "ver-email.html", gin.H{
			"message": "email is not reachable, looking for " + ret.Suggestion + " ins",
		})
		return
	}
	if ret.Reachable == "no" {
		fmt.Println("email address is not reachable")
		ctx.HTML(http.StatusBadRequest, "ver-email.html", gin.H{
			"message": "email address was unreachable",
		})
		return
	}

	if !ret.HasMxRecords {
		fmt.Println("domain entered not properly setup to recieve emails, MX record not found")
		ctx.HTML(http.StatusBadRequest, "ver-email.html", gin.H{
			"message": "entered not properly setup to recieve emails, MX record not found",
		})
	}
	ctx.HTML(http.StatusOK, "ver-email.html", gin.H{
		"email": email,
	})

}
