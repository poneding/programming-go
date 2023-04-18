package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Format("20060102150405"))
	// 如果要格式化到毫秒，微秒，纳秒，这里需要跟在小数点后面
	// 其他格式显示0，例如：20060102150405-000000000，将格式化为 yyyyMMddHHmmss-000000000
	p(now.Format("20060102150405.000000000"))
	now.Nanosecond()

	then := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
