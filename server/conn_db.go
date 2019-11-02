package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	RetryWeight int = 2 // DBの接続に失敗した時の再接続時間の重み(2なら2, 4, 6, 8...と待つ)
	RetryTimes  int = 5 // DBの接続の最大施行回数
)

// DBへの接続を行い、インスタンスを取得する
func ConnectToDB(dbSrc string) (db *gorm.DB, err error) {
	// 接続する
	// 指定回まで接続を再施行する
	for i := 1; i < RetryTimes; i++ {
		db, err := gorm.Open("mysql", dbSrc)
		if err != nil {
			waiting := time.Duration(i*RetryWeight) * time.Second
			fmt.Printf("Can't initialize db. Retrying in %d secs... : %v\n", int(waiting.Seconds()), err)
			time.Sleep(waiting)
			continue
		}

		// 接続に成功した
		return db, nil
	}

	errMsg := fmt.Sprintf("Can't initialize db.")
	return nil, errors.New(errMsg)
}
