package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/xorwise/golang-todo-api/bootstrap"
	"github.com/xorwise/golang-todo-api/domain"
	"gorm.io/gorm"
)

func CheckDeadlines(env *bootstrap.Env, db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Checking deadlines")
		rows, err := db.Model(&domain.Task{}).Where("completed = ?", false).Rows()
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var task domain.Task
			if err = db.ScanRows(rows, &task); err != nil {
				log.Fatal(err)
			}

			if task.Deadline.Before(time.Now()) {
				task.Completed = true
				if err = db.Save(&task).Error; err != nil {
					log.Fatal(err)
				}
				env.Mu.Lock()
				if ch, ok := env.ClientChannels[task.UserID]; ok {
					ch <- fmt.Sprintf("Task %d is completed", task.ID)
				}
				env.Mu.Unlock()
			}
		}
		rows.Close()
	}
}
