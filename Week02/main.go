package main

import (
	"database/sql"
	"fmt"
	 er "errors"
	"github.com/pkg/errors"
)

//作业 - 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//dao 层中直接抛出 sql.ErrNoRows 不 Wrap 这个 error 
//上层 接受到 sql.ErrNoRows 错误，Wrap 这个错误，包含 堆栈信息 继续返回
func main() {
	err := Service()
	fmt.Printf("%+v",err)
}

func Service() error {
	err := Dao()
	return  errors.Wrap(err,"数据为空")

	fmt.Errorf()

	er.Unwrap()
}

func Dao() error {
	return sql.ErrNoRows
}


