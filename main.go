package main

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHostname() (host string) {
	host, err := os.Hostname()
	if err != nil {
		return
	}
	return
}

func GetIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok {
			if ipNet.IP.IsLoopback() {
				continue
			}
			if ipNet.IP.To4() == nil {
				continue
			}
			ip = ipNet.IP.String()
			return
		}
	}
	return
}

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hostname": GetHostname(),
			"ip":       GetIP(),
		})
	})

	r.Run(":80")
}
