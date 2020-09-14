package main

import (
	"github.com/utf6/go-blog/models"
	"log"
)

func main()  {
	log.Println("Starting....")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle....")

		models.CleanAllArticle()
	})

	c.Start()
}
