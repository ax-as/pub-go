package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-vgo/robotgo"
	"github.com/ncruces/zenity"
)

var cmdMap = map[string]bool{
	"-b": true,
	"-p": true,
}

func main() {

	selText, err := ex("xclip", "-r", "-o", "-selection", "primary")
	robotgo.MilliSleep(10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	if selText == "" {
		return
	}

	params := []string{
		"-t", "pt+en",
	}

	for _, p := range os.Args {
		if cmdMap[p] {
			params = append(params, p)
		}
	}
	params = append(params, selText)

	text, err := ex("crow", params...)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}

	zenity.Info(text, zenity.Title("Tradução"), zenity.InfoIcon)
}

func ex(cmd string, params ...string) (string, error) {
	fmt.Println(params)
	out, err := exec.Command(cmd, params...).Output()
	return string(out), err
}
