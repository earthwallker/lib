
package mathlib

import (
    "fmt"
)

// 导出一个函数用于加法操作
func Add(a, b int) int {
    c:=a + b
    fmt.Println(c)
    return c
}