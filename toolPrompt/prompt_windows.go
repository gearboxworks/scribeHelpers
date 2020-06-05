// +build windows

package toolPrompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
