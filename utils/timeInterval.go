package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// 打印函数执行时间
func TimeInterval(f func(), times int) {
	fnName := GetFunctionName(f)
	start := time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	end := time.Now()
	interval := end.Sub(start)
	fmt.Printf("Function [%v] Execution Duration: %v\n", fnName, interval)
}

// 获取函数名称
func GetFunctionName(i interface{}) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return fn
}
