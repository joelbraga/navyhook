package base
import (
	"os/exec"
	"bytes"
	"fmt"
	"errors"
)


func RunUnixCommand(path string) (string, error)  {
	cmd := exec.Command("bash",  path)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}
	return "Result: " + out.String(), nil
}