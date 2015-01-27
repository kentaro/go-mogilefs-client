package mogilefs

import (
	"github.com/lestrrat/go-tcputil"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	var exitCode = 0
	defer os.Exit(exitCode)

	go exec.Command("mogilefsd", "-c", "mogilefs/mogilefsd.conf").Run()
	go exec.Command("mogstored", "-c", "mogilefs/mogstored.conf").Run()

	defer exec.Command("pkill", "-f", "mogilefsd").Run()
	defer exec.Command("pkill", "-f", "mogstored").Run()

	var err error
	err = tcputil.WaitLocalPort(7001, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	err = tcputil.WaitLocalPort(7500, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	err = tcputil.WaitLocalPort(7501, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	exitCode = m.Run()
}
