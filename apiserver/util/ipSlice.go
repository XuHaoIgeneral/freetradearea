package util

import "strings"

func IpSlice(ipStr string) (respStr string) {
	respStr = strings.Split(ipStr, ":")[0]
	return
}

func HttpSlice(url string) (respStr string) {
	respStr = strings.Split(url, "//")[1]
	return
}
