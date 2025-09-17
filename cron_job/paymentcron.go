package cronjob

import (
	"feebee/db"
	"feebee/model"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// sending notification

func PaymentCron() {
	c := cron.New()

	c.AddFunc("1 * * * *", func() {

		var details []uint

		db.DB.Model(&model.Payment{}).Distinct().Pluck("student_id", &details)

		for _, id := range details {
			var firtpayment model.Payment

			db.DB.Where("student_id = ?", id).
				Order("created_at ASC").First(&firtpayment)

			monthpassed := int(time.Since(firtpayment.CreatedAt).Hours() / (24 * 30))
			nextDue := firtpayment.CreatedAt.AddDate(0, monthpassed+1, 0)

			if nextDue.Format("2025-01-02") == time.Now().Format("2025-06-02") {
				fmt.Printf("payemnt is due today %d", id)
			}

		}
	})

	c.Start()
}
