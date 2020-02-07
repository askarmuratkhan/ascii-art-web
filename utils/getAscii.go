package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetASCII(f, s string) string {
	s = strings.ReplaceAll(s, string([]byte{13, 10}), string([]byte{92, 110}))
	fmt.Println(s, []byte(s))
	out, err := exec.Command("./utils/main.exe", s, f).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(out)
	// sout := strings.ReplaceAll(string(out), "\n", "<br>")
	return string(out)
}
