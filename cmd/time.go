package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"xgocli/internal/timer"
)

var (
	calculateTime string
	duration      string
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format conversion",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		fmt.Printf("result: %s, %d\n", nowTime.Format("2006-01-0215:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "calculate time in need",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2016-01-0215:04:05"
		location, _ := time.LoadLocation("Asia/Taipei")

		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")

			if space == 0 {
				layout = "2016-01-02"
			} else if space == 1 {
				layout = "2006-01-0215:04"
			}

			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0).In(location)
			}
		}

		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("error! %v\n", err)
		}

		fmt.Printf("result:\n %s\n %s\n %d\n", t.Format(time.RFC3339Nano), t.Format(time.ANSIC), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "The time need to be calculated, should be a timestamp or formatted time.")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `Duration, using the following units: "ns", "us", "ms", "s", "m", "h"`)
}
