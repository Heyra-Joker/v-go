/*
@author:Joker
@file: judgement.go
@time: 2021/12/13/10:30 上午
@blog: https://github.com/joker-heyra
@description: --

🤡
code is far away from bugs with the god animal protecting
    I love animals. They taste delicious.
🤡
*/

package judgement

type Judge interface {
	verify() error
}

// New verify some flag value
func New(judge Judge) error {
	return judge.verify()
}
