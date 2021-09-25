package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"app/helper//mail"
	"app/helper/rabbitmq"
)

//struct Worker define a worker
type Worker struct {
	emailer 	mail.Emailer
	channel 	*rabbitmq.ChannelMQ
	db      	*sql.DB
	ctx			context.Context
}

//func NewWorker create a new Worker
func NewWorker(emailer mail.Emailer, MQchan *rabbitmq.ChannelMQ, db *sql.DB, ctx context.Context) *Worker {
	return &Worker{
		emailer: emailer,
		channel: MQchan,
		db:      db,
		ctx:     ctx,
	}
}

/*
func (Worker) Start starts Worker
process logic:
	1. Wait message in channel inChan
	2. Send email with each Emailer
	3. Update database (thankyou_email_sent = true)
*/
func (w *Worker) Start() {
	if w.emailer == nil || w.db == nil {
		fmt.Println("Can't start Worker with emailer nil (or db nil)")
		return
	}
	sttm,err := w.db.Prepare("UPDATE `db_order` SET thankyou_email_sent = ? WHERE id =?;")
	if err != nil {
		fmt.Println("Can not update thankyou_email_sent")
	}
	msgs, err := w.channel.ConsumeMQ()
	fmt.Println("Worker start sending ...")
	endChan := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Have email to send ",d.Body)
			var ec mail.EmailContent
			err := json.Unmarshal(d.Body,&ec)
			if err != nil {
				fmt.Println("Can't unMarshal body msgs:",err)
				continue
			}
			err = w.emailer.Send(&ec)
			if err != nil {
				fmt.Println("Can't send msgs :",err)
				continue
			}
			_,err = sttm.Query(1, ec.Id)
			if err != nil {
				fmt.Println("Can't update thankyou_email_sent:",err)
			}
		}
	}()
	<-endChan
}