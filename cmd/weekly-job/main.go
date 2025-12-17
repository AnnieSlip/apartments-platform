package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ani-javakhishvili/apartments-platform/domain/storage/cassandra"
)

func main() {
	ctx := context.Background()

	if err := cassandra.Connect(); err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

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
	currentWeek := getCurrentWeek()

	// fetch all user matches for current week
	userMatches, err := cassandra.NewRepository(cassandra.Session).GetAllMatchesForWeek(ctx, currentWeek)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for userID, aptIDs := range userMatches {
		wg.Add(1)
		go func(uID int, aIDs []int) {
			defer wg.Done()
			sendNotification(uID, aIDs)
		}(userID, aptIDs)
	}

	wg.Wait()
	return nil
}

func getCurrentWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}
