package bot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

type System interface {
	//
	//@notice returns the current system configurations
	//@notice Hostname string;
	//@notice Username string;
	//@notice Ncpu int;
	//@notice Ram	int;
	//@dev Ip {}string; an device could have many interfaces
	Init()
	GetHostName() string
	GetUserName() string
	GetNcpu() int
	GetRam() int
	GetIp() string
}

// @notice test if a Bot instance parametized with system attributes get created
type MockSystem struct {
	Hostname string
	Username string
	Ncpu     int
	Ram      int
	Ip       []map[string]string
}

// @notice: *MockSystem is different from MockSystem
// @notice: it is *MockSystem which satisfies the System interface
// @notice: and not MockSystem itself
// @notice: so we pass &MockSystem to the want:=Bot{}
func (m *MockSystem) Init() {
	m.Hostname = "kinetic-kudu"
	m.Username = "ubuntu"
	m.Ncpu = 2
	m.Ram = 2048
	m.Ip = make([]map[string]string, 10)
	m.Ip[0] = map[string]string{"eth1": "10.1.1.53"}
	m.Ip[1] = map[string]string{"eth2": "192.168.6.9"}
	m.Ip[2] = map[string]string{"eth3": "172.14.8.299"}
	m.Ip[3] = map[string]string{"eth4": "188.25.61.53"}
	m.Ip[4] = map[string]string{"eth5": "192.168.17.10"}
}

func (m *MockSystem) GetHostName() string {
	return m.Hostname
}
func (m *MockSystem) GetUserName() string {
	return m.Username
}
func (m *MockSystem) GetNcpu() int {
	return m.Ncpu
}
func (m *MockSystem) GetRam() int {
	return m.Ram
}
func (m *MockSystem) GetIp() string {
	return m.Ip[0]["eth1"]
}

// @notice
type LinuxSystem struct {
	Hostname string
	Username string
	Ncpu     int
	Ram      int
	Ip       string
}

func getLinuxRam() int {
	file, err := os.OpenFile("/proc/meminfo", os.O_RDONLY, 0400)
	if err != nil {
		log.Println(err)
	}
	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		log.Println(err)
	}
	_, token, err := bufio.ScanLines(buf, true)
	if err != nil {
		log.Println(err)
	}
	line := string(token)
	ram := 0
	fmt.Sscanf(line, "MemTotal:        %d kB", &ram)

	log.Printf("ram: %d GB", ram/1024)
	return int(ram / 1024)
}

func (m *LinuxSystem) Init() {
	m.Hostname, _ = os.Hostname()
	m.Username = os.Getenv("USER")
	m.Ncpu = runtime.NumCPU()
	m.Ram = getLinuxRam()
}

func (m *LinuxSystem) GetHostName() string {
	return m.Hostname
}
func (m *LinuxSystem) GetUserName() string {
	return m.Username
}
func (m *LinuxSystem) GetNcpu() int {
	return m.Ncpu
}
func (m *LinuxSystem) GetRam() int {
	return m.Ram
}
func (m *LinuxSystem) GetIp() string {
	return ""
}
