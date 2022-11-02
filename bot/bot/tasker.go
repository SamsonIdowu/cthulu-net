package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Task struct {
	// we call them tags, and serve as hints
	// to json.UnMarshal  function for deserializing
	// the json
	// This fields should be exported (Have First letter in Caps)
	// so that the json package can access them for modification
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port int  `json:"port"`
	Type string `json:"type"`
}
type Result struct {
	TaskId      string `json:"TaskId"`
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
	Ip_status   string `json:"ip_status"`
	Port_status string `json:"port_status"`
}

type Tasker interface {
	Next() Task
	Report(r Result, b Bot)
}

type MockTasker struct {
	JsonObj string
}

func (mt MockTasker) Next() Task {
	t := Task{}
	json.Unmarshal([]byte(mt.JsonObj), &t)
	log.Printf("#Tasker fetched new task: %v", t)
	return t
}

func (mt MockTasker) Report(r Result, b Bot) {
	log.Printf("#Tasker Reporting result: %v", r)
}

type TaskerProxy struct {
	Url string
}

func (proxy TaskerProxy) Next() Task {
	url := fmt.Sprintf("%s/bot/tasks", proxy.Url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//log.Println(err)
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t := Task{}
	json.Unmarshal(body, &t)
	log.Printf("#Tasker fetched new task: %v", t)
	return t
}

func (proxy TaskerProxy) Report(r Result, b Bot) {
	url := fmt.Sprintf("%s/bot/%s", proxy.Url, b.uuid)
	jsonstr := fmt.Sprintf(`{"TaskId": "%s","Ip_status": "%s","Port_status": "%s"}`,
		r.TaskId, r.Ip_status, r.Port_status)

	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(jsonstr))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
