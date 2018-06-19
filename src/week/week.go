package week

var logWeekend = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func GetStringByNumber(idx int) string {
	idx = idx % 7
	return logWeekend[idx]
}
