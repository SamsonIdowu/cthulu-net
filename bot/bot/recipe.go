package bot
import (
	"net"
	"fmt"
	"time"
	"log"
)

type Recipe interface {
	Do(t Task) Result
}

type MockRecipe struct{}

func (m MockRecipe) Do(t Task) Result {
	return Result{TaskId: t.Id, Ip: t.Ip, Port: int(t.Port), Ip_status: "up", Port_status: "filtered"}
}

// I could then build several types of recipes here, like
// ScanRecipe
// SSHBruteRecipe
// FTPBruteRecipe
// SMBBruteRecipe

type ScanRecipe struct{}

func (recipe ScanRecipe) Do(t Task) Result {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.Ip, int(t.Port)), 5*time.Second)
	if err != nil {
		log.Println(err)
		return Result{TaskId: t.Id, Ip: t.Ip, Port: int(t.Port), Ip_status: "down", Port_status: "down"}
	}
	defer conn.Close()
	return Result{TaskId: t.Id, Ip: t.Ip, Port: int(t.Port), Ip_status: "up", Port_status: "up"}
}
