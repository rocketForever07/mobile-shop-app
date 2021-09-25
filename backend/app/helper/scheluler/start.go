package scheduler

import (
	"context"
	"database/sql"
	"fmt"
	"app/helper/rabbitmq"
	"app/helper//scheduler"
	"time"

	"app/helper/mail"
	"lecture12_rabbitmq/native/worker"
	"gorm.io/gorm"

)

func start(db *gorm.DB) {

	ctx := context.Context(context.Background())
	rmq := rabbitmq.NewChannelMQ("shop_mobile_app")

	go func() {
		sched := scheduler.NewScheduler(db, rmq, ctx)
		sched.Start()
	}()

	mySendgrid := mail.NewSendgrid("apiKey")
	emailer := mail.Emailer(mySendgrid)
	testWork := worker.NewWorker(emailer, rmq, db, ctx)
	testWork.Start()
	fmt.Println("Exit at ",time.Now())
}