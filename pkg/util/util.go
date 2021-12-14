/*
@File    :   util/util.go
@Time    :   2021/12/10 11:39:22
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package util

import (
	"errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// PanicErr panic error
func PanicErr(err error) {
	if err != nil {
		panic(color.RedString(err.Error()))
	}
}

// CobraErr panic err from Cobra ERR func
func CobraErr(msg string) {
	msg = color.RedString(msg) + "\ncan use v-go -h see more information"
	cobra.CheckErr(errors.New(msg))
}
