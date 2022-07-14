package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var (
	Proto string
	Port  string
)

func init() {
	flag.StringVar(&Proto, "proto", "tcp", "The network must be tcp, tcp4, tcp6, unix or unixpacket")
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		println("Usage: kill-port <port>")
		os.Exit(1)
	}
	Port = os.Args[1]

	if Port == "" {
		println("Please specify port")
		os.Exit(1)
	}

	if PortIsUsed(Proto, Port) == true {
		FreePort(Proto, Port)
	} else {
		PortNotUsed()
	}

}

func PortNotUsed() {
	println("Port is not used")
	os.Exit(1)
}

func PortIsUsed(proto, port string) bool {
	_, err := net.Listen(proto, ":"+port)
	if err != nil {
		return true
	}

	return false
}

// Execute command and return exited code.
func execCmd(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		println("failed to release port error:", err.Error())
		os.Exit(1)
	}
}

func FreePort(proto, port string) {
	command := fmt.Sprintf("lsof -i %s:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", proto, port)
	execCmd(exec.Command("bash", "-c", command))
	os.Exit(0)
}
