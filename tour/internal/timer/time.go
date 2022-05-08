package timer

import "time"

// GetNowTime 获取时间
func GetNowTime(timezone string) time.Time {
	//return time.Now()
	if timezone == "" {
		timezone = "Asia/Shanghai"
	}
	location, _ := time.LoadLocation(timezone)
	return time.Now().In(location)
}

// GetCalculateTime 推算时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
