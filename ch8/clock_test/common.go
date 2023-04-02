package main

const (
	Beijing = "Beijing"
	NewYork = "NewYork"
	London  = "London"
)

var portToArea = map[string]string{
	"8010": Beijing,
	"8020": NewYork,
	"8030": London,
}

var areaToTimeZoo = map[string]string{
	Beijing: "Asia/Shanghai",
	NewYork: "America/New_York",
	London:  "Europe/London",
}
