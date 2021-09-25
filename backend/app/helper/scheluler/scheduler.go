package scheduler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"app/helper/rabbitmq"
	"time"

	cron "github.com/robfig/cron/v3"
	"app/helper/mail"
	"app/models/entity"
)
/*
Goal: Every 1 minute
	- Scan order from database -> array EmailContent
	- Add host infor, clients infor, content ... to array EmailContent
	- Put this array EmailContent into channel of sending
*/
const (
	DefaultThankyouSubject 		= 	"Thank you for purchasing from mystore.com"
	DefaultThankyouBodyPlain	=	"Thank you for purchasing from our store. Here's your order details:"
	DefaultThankyouBodyHtml		=	"<strong>Thank you for purchasing from our store. Here's your order details:</strong>"
	DefaultFromName				=	"My Store Owner"
	DefaultFromEmail			=	"support@mystore.com"
)



type Scheduler struct{
	db 			*gorm.DB
	cr      *cron.Cron
	channel *rabbitmq.ChannelMQ
	ctx     context.Context
}

type ScanDB struct{
	ID int
	CustomerName string
	CustomerEmail string
}
//Create new Scheduler
func NewScheduler(db *sql.DB, MQchan *rabbitmq.ChannelMQ, ctx context.Context) *Scheduler{
	return &Scheduler{
		db:      db,
		cr:      cron.New(cron.WithSeconds()),
		channel: MQchan,
		ctx:     ctx,
	}
}


func (sched *Scheduler) Start() {
	sched.cr.AddFunc("0 * * * * *", sched.ScheduleJob)
	fmt.Println("Scheduler start at ", time.Now().Format("2021-September-23 09:12:10"))
	sched.cr.Start()
}

func (sched *Scheduler) Stop() {
	sched.Stop()
	fmt.Println("Scheduler stop at ", time.Now().Format("2021-September-23 09:12:10"))
}

//func (Scheduler) ScheduleJob put email's clients to channel of sending
func (sched *Scheduler) ScheduleJob() {
	resp,err := sched.SetHostEmailForSending()
	if err != nil {
		return
	}
	fmt.Println("Scheduling send ", len(resp)," emails at ", time.Now().Format("2021-September-23 09:12:10"))

	for _, email := range resp {
		body := email.String()
		err = sched.channel.PublishMQ(body)
		if err != nil {
			fmt.Println("Can't channel mes")
		}
	}
}

//func (Scheduler) GetHostEmailForSending set address email of host for sending email's clients scaned from database
func (sched *Scheduler) SetHostEmailForSending() ([]*mail.EmailContent, error){
	resp,err := sched.ScanFromDB()
	if err != nil {
		return resp,err
	}
	for _,emailContent := range resp {
		emailContent.FromUser = &mail.EmailUser{
			Name: 	DefaultFromName,
			Email: 	DefaultFromEmail,
		}
	}
	return resp,nil
}

func (sched *Scheduler) ScanFromDB() ([]*mail.EmailContent, error){
	var resp []*mail.EmailContent
	fromTime := time.Now().Add(-time.Minute*2)
	fmt.Println("Scan from database from ",fromTime, " to ",time.Now())

	sttm := sched.db.Model(&entity.Order).Limit(10).Find(&ScanDB{}).Where("creat_at >= @fromTime AND thankyou_email_sent=0 OR thankyou_email_sent is null")


	for i:=0;i<len(sttm);i++ {

		resp = append(resp, &mail.EmailContent{
			Id: sttm[i].ID,
			Subject: DefaultThankyouSubject,
			ToUser: &mail.EmailUser{
				Name: sttm[i].CustomerName,
				Email: sttm[i].CustomerEmail,
			},
			PlainContent: DefaultThankyouBodyPlain,
			HtmlContent: DefaultThankyouBodyHtml,
		})
	}
	return resp,nil
}