package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sendEmail(c *gin.Context) {
	var event struct {
		Event struct {
			Data struct {
				New struct {
					Email string `json:"email"`
					Name  string `json:"name"`
				} `json:"new"`
			} `json:"data"`
		} `json:"event"`
	}

	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	to := event.Event.Data.New.Email
	subject := "Welcome to Our Service"
	body := "Hi " + event.Event.Data.New.Name + `,

Welcome to ABC_cinema!

We're excited to have you with us. Start exploring movie schedules, bookmark your favorites, and never miss a show.

If you need any assistance, reach out to our support team at abc.cinema.com@gmail.com.

Enjoy the movies!

Best,
The ABC_cinema Team`

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

func main() {
	r := gin.Default()
	r.POST("/send-email", sendEmail)

	if err := r.Run(":3045"); err != nil {
		log.Fatal(err)
	}
}
