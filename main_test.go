package mogilefs

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	var exitCode = 0
	defer os.Exit(exitCode)

	go exec.Command("mogilefsd", "-c", "mogilefs/mogilefsd.conf").Run()
	go exec.Command("mogstored", "-c", "mogilefs/mogstored.conf").Run()

	defer exec.Command("pkill", "-f", "mogilefsd").Run()
	defer exec.Command("pkill", "-f", "mogstored").Run()

	exitCode = m.Run()
}
