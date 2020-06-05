// +build darwin

package toolPrompt

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)


func (p *TypePrompt) UserPrompt() string {
	var ret string

	for range onlyOnce {
		fmt.Printf("%s", p.string)

		r := bufio.NewReader(os.Stdin)

		var err error
		ret, err = r.ReadString('\n')
		fmt.Printf("\n")
		if err != nil {
			break
		}

		ret = strings.TrimSuffix(ret, "\n")
	}

	return ret
}


func (p *TypePrompt) UserPromptHidden() string {
	var ret string

	for range onlyOnce {
		fmt.Printf("%s", p.string)

		hidden, err := terminal.ReadPassword(syscall.Stdin)
		fmt.Printf("\n")
		if err != nil {
			break
		}

		ret = strings.TrimSuffix(string(hidden), "\n")
	}

	return ret
}
