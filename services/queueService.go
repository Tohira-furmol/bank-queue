package services

import (
	"errors"
	"time"

	"bank-queue-system/models"
)

var queues []models.Queue

func BookQueue(queue models.Queue) (models.Queue, error) {
	if queue.TimeSlot.Before(time.Now()) {
		return models.Queue{}, errors.New("time slot must be in the future")
	}

	queue.ID = uint(len(queues) + 1)
	queue.TicketNo = generateTicketNo()
	queue.CreatedAt = time.Now()

	queues = append(queues, queue)

	return queue, nil
}

func GetQueue(id uint) (models.Queue, error) {
	for _, queue := range queues {
		if queue.ID == id {
			return queue, nil
		}
	}
	return models.Queue{}, errors.New("queue not found")
}

func CancelQueue(id uint) error {
	for i, queue := range queues {
		if queue.ID == id {
			// Remove the queue from the slice
			queues = append(queues[:i], queues[i+1:]...)
			return nil
		}
	}
	return errors.New("queue not found")
}

func UpdateQueueTime(id uint, newTimeSlot time.Time) (models.Queue, error) {
	for i, queue := range queues {
		if queue.ID == id {
			if newTimeSlot.Before(time.Now()) {
				return models.Queue{}, errors.New("time slot must be in the future")
			}

			queues[i].TimeSlot = newTimeSlot
			return queues[i], nil
		}
	}
	return models.Queue{}, errors.New("queue not found")
}

func generateTicketNo() string {
	return "Q" + time.Now().Format("150405") // Format as "QHHMMSS"
}
