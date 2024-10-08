package rmtpsetup

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
)

// This struct contains information about the RTMP Ip Address
// In future it can consist of things like where to save stream
type RTMPModule struct {
	IpAddr net.IP
}

func (rtmp *RTMPModule) setIpAddress() {
	conn, err := net.Dial("udp", "8.8.8.8:80") // Dial the udp dns to get the outbound address
	if err != nil {
		log.Fatalf("ERROR: cannot get local Ip Address: %v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("LOG: Streaming IP address : %s", localAddr.IP.String())
	rtmp.IpAddr = localAddr.IP
}

// Do all sorts of scripting stuff from installing and all that shi
func (rtmp *RTMPModule) Run(wg *sync.WaitGroup) error {
	defer wg.Done()
	rtmp.setIpAddress()
	cmd := exec.Command("docker", "compose", "-f", "./rtmp_setup/docker-compose.yml", "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ERROR: Cannot Run rtmp server: %v", err)
	}
	ouput, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("ERROR: Cannot get ouput of command %v", err)
	}
	fmt.Printf("ouput: %v", ouput)
	return nil
}
