# Cthulu Framework
---


![](__static/1.png)


# Intro

CTHULU is a distributed scanner, which spreads the pool of ip addresses to scan over a group of relatively medium capacity devices (horizontal scaling botnet-style), in order to reduce scan time.

Demo: https://youtu.be/mk22sYc7R1o

# Setup


## I. Backend


1. Start `server`, `operator-cli`, `operator-dashboard` services

```sh
$  docker compose up
```

- `server` listens on port `5000` of your host
- `operator-cli` listens on port `5001` of your host
- `operator-dashboard` listend on port `3000` (https) of your host

2. Install Grafana Infinity Plugin, and create default infinity DataSource

-  Search for Infinity Plugin

![Pasted image 20221102032217.png](__static/2.png)


- Install Infinity Plugin

![Pasted image 20221102032336.png](__static/Pasted%20image%2020221102032336.png)

- Create Infinity DataSource

![Pasted image 20221102032451.png](__static/Pasted%20image%2020221102032451.png)


3. Import the json dashboard `./operator/BotNet Master.json` 

![Pasted image 20221102033331.png](__static/Pasted%20image%2020221102033331.png)

After import you should obtain this dashboard. Not the most creative but i bet you can customize it and make pull request )).

![Pasted image 20221102033550.png](__static/Pasted%20image%2020221102033550.png)



## II. Bot client

  
1. Change ip address of server in `./bot/main.go`

```go
func main() {
    sys := &bot.LinuxSystem{}
    sys.Init()
    tasker := bot.TaskerProxy{Url: "http://server:5000"}
    tracker := bot.TrackerProxy{Url: "http://server:5000"}
```

2. Build the bot client binary
  

```sh
$ ls
bot  go.mod  main.go
$ go mod tidy
$ go build -o botclient .
$ ls
bot  botclient  go.mod  main.go
$

```

  

2. Copy and Run the `botclient` binary to workstations you want to use as `WorkerBots` within your network.

![Pasted image 20221102041624.png](__static/Pasted%20image%2020221102041624.png)

![Pasted image 20221102042433.png](__static/Pasted%20image%2020221102042433.png)





## III. Monitor bots added to pool

![Pasted image 20221102044024.png](__static/Pasted%20image%2020221102044024.png)


## IV. Provide ip subnets to scan and monitor

To connect as an operator and schedule scan jobs, connect to the operator cli on port `5001`

![Pasted image 20221102050201.png](__static/Pasted%20image%2020221102050201.png)

![Pasted image 20221102050701.png](__static/Pasted%20image%2020221102050701.png)


# Architecture (at least the intended  on :) )

![Pasted image 20221102051110.png](__static/Pasted%20image%2020221102051110.png)

# Uml Modelling
![uml](./UML%20Model.png)

  

# TODO:

  
- Add nginx reverse-proxy to `docker-compose.yml`
