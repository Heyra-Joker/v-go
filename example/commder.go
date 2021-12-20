/*
@File    :   example/commder.go
@Time    :   2021/12/08 11:59:26
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package main

import (
	"log"

	"github.com/Heyra-Joker/v-go/pkg/commder"
)

func main() {
	cmd := "ls -la"
	result, err := commder.Command(cmd)
	if err != nil {
		panic(err)
	}

	log.Printf("success: %v, out: %s, err: %s", result.Success, result.Stdout, result.StdErr)
}
