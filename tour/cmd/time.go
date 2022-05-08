package cmd

import (
	"github.com/d-xingxing/go-programming-tour/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string
var timezone string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime(timezone)
		log.Printf("输出时间: %s,(时间戳)秒:%d,毫秒:%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix(), nowTime.UnixMilli())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		//Jan 2 15:04:05 2006 MST
		//1   2  3  4  5    6  -7
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime("")
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出时间: %s,(时间戳)秒:%d,毫秒:%d", t.Format("2006-01-02 15:04:05"), t.Unix(), t.UnixMilli())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	nowTimeCmd.Flags().StringVarP(&timezone, "timezone", "z", "", `时区，默认为"Asia/Shanghai"`)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间,有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间,有效时间单位为"ns","us","ms","s","m","h"`)
}

/**
$ go run main.go time now
输出时间: 2022-05-08 17:34:57,(时间戳)秒:1652002497,毫秒:1652002497726

$ go run main.go time now -z="Asia/Shanghai"
$ go run main.go time now -z="America/New_York"
$ go run main.go time now -z="America/Los_Angeles"

$ go run main.go time calc -c="2022-05-08 17:34:57" -d=5m
输出时间: 2022-05-08 17:39:57,(时间戳)秒:1652031597,毫秒:1652031597000

$ go run main.go time calc -c="2022-05-08 17:34:57" -d=-2h
输出时间: 2022-05-08 15:34:57,(时间戳)秒:1652024097,毫秒:1652024097000
*/
