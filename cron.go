package main

import (
	"github.com/robfig/cron"
	"github.com/utf6/go-blog/models"
	"log"
	"time"
)

func main()  {
	log.Println("Starting....")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle....")

		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)

	for  {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}