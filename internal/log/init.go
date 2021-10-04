package log

import (
	"github.com/beckbikang/flg"
	"go.uber.org/zap"
)

//you can define you log var outer
var (
	Sl   *zap.Logger
	gflg *flg.Logger
)

// init log
func InitLog(filePath string) {
	gflg = &flg.Logger{}

	err := gflg.LoadFromFile(filePath)
	if err != nil {
		panic("get file faild")
	}

	Sl, err = gflg.GetLogByKey("test")
	if err != nil {
		panic(err)
	}
}
