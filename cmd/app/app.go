package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"example-oauth/pkg/lib/oauth"
)

func main() {
	ctx := context.Background()
	tokenSource, err := oauth.TokenSource(ctx, "client_secrets.json", "request.token", calendar.CalendarScope)
	if err != nil {
		log.Fatal("token: ", err)
	}

	calendarService, err := calendar.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatal("service: ", err)
	}

	calendars, err := calendarService.CalendarList.List().Do()
	if err != nil {
		log.Fatal("get: ", err)
	}

	// print user calendars
	for _, item := range calendars.Items {
		fmt.Println(item.Summary)
	}
}
