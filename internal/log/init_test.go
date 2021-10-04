package log

import (
	"testing"
)

/**
测试初始化
*/
func TestInitLog(t *testing.T) {

	InitLog("../../test/test_log.toml")
	Sl.Info("test")
}
