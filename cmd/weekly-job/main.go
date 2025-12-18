package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	log.Println("Weekly job started")

	if err := runWeeklyJob(ctx); err != nil {
		log.Fatalf("Weekly job failed: %v", err)
	}

	log.Println("Weekly job finished")
}

func sendNotification(userID int, apartmentIDs []int) {
	// here goes email sending logic -  with email, push, etc.
	log.Printf("Sending notification to user %d: matched apartments %v\n", userID, apartmentIDs)
}

func runWeeklyJob(ctx context.Context) error {

	return nil
}

func getCurrentWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}
