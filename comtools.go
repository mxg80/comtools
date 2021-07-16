package comtools

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}
func GetHostName() string {
	name, _ := os.Hostname()
	return name + "yzmb" + "v1.0.1"
}
func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v\n", err)
		return ""
	} else {
		localIP := strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
		return localIP
	}

}
