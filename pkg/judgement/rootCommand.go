/*
@author:Joker
@file: rootCommand.go
@time: 2021/12/13/10:10 上午
@blog: https://github.com/joker-heyra
@description: root command judgement.

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package judgement

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/Heyra-Joker/v-go/pkg/util"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type Root struct {
	Path       string
	ShowFormat string
	Output     string
	Coverage   bool
	Name       string
}

var allowShowFormatFlag = []string{"csv", "flat", "ini", "json", "xml"}

// verify root command flags
// include path, show_format
// path: whether exists or empty
// show_format: whether legal of allowShowFormatFlag
func (r Root) verify() error {
	// check path
	if r.Path == "" {
		return errors.New(color.RedString("path empty"))
	}

	if !util.PathExists(r.Path) {
		return errors.New(color.RedString("path: %s not exists", r.Path))
	}

	// check output
	output := r.setOutputDefault()

	if output != "" {
		if util.PathExists(output) && !r.Coverage {
			return errors.New(color.RedString("output: %s is exists, you can use -c[--coverage] to override", output))
		}
	}

	// check show_format
	pass := util.ItemContainer(r.ShowFormat, allowShowFormatFlag)
	if !pass {
		return errors.New(color.RedString("show_format just accept %s", strings.Join(allowShowFormatFlag, ", ")))
	}

	return nil
}

// setOutputDefault, whether set output default value when it empties
func (r Root) setOutputDefault() string {
	output := r.Output

	if output != "" {
		return output
	}

	setDefault := false
	// when arg is compressing
	if util.ItemContainer(r.Name, [1]string{"compress"}) {
		setDefault = true
	}

	if setDefault {
		fileName := "vgo-" + filepath.Base(r.Path)
		output = filepath.Join(filepath.Dir(r.Path), fileName)
		viper.Set("output", output)
	}

	return output
}
