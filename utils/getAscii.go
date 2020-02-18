package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsAlphanumerical(s string) bool {
	s = strings.ToLower(s)
	ok := true
	for _, r := range s {
		ok = ok && ((int(r) > 32 && int(r) < 132) || int(r) == 134 || r == '\n' || r == ' ' || r == '\r')
	}
	return ok
}
func GetASCII(f, s string) (string, error) {
	//s = strings.Replace(s, string([]byte{13, 10}), string([]byte{92, 110}), -1)
	s = strings.ReplaceAll(s, string([]byte{13, 10}), string([]byte{92, 110}))
	fmt.Println(s, []byte(s))
	//out, err := exec.Command("./utils/ascii-art", s, f).Output() //linux
	out, err := exec.Command("./utils/main.exe", s, f).Output() //windows
	//out, err := exec.Command("./utils/main", s, f).Output() //macos
	if err != nil {
		return "", err
	}
	fmt.Println("no exec err")
	return string(out), nil
}
