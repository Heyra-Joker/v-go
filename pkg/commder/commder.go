/*
@File    :   commder/commder.go
@Time    :   2021/12/08 11:39:00
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   bash -c command something
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package commder

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
)

type Result struct {
	Success bool
	Stdout  string
	StdErr  string
}

// Command Run Command at input cmd.
func Command(cmd string) (Result, error) {
	res := Result{}
	stdOutString := ""

	command := exec.Command("bash", "-c", cmd)

	// define std out
	stdout, err := command.StdoutPipe()
	if err != nil {
		return Result{}, err
	}

	// define std error
	errBuff := bytes.NewBuffer(nil)
	command.Stderr = errBuff

	// start run cmd
	err = command.Start()
	if err != nil {
		return Result{}, err
	}

	// define std content
	reader := bufio.NewReader(stdout)

	// reading...
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		stdOutString += line
	}

	_ = command.Wait()

	res.StdErr = errBuff.String()
	res.Stdout = stdOutString
	res.Success = command.ProcessState.Success()
	return res, nil
}
