package ip_proxy_service

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-resty/resty/v2"
)

type IpProxy struct {
	IP          string
	Port        string
	Protocol    string
	CountryCode string
}

var pool []IpProxy

func GetProxyPool() []IpProxy {
	if len(pool) < 1 || pool == nil {
		pool = []IpProxy{}
		client := resty.New()
		if "dev" == os.Getenv("RUN_ENV") {
			client.SetProxy(`http://127.0.0.1:1081`)
		}

		request := client.R().EnableTrace()
		resp, err := request.Get("https://raw.githubusercontent.com/clarketm/proxy-list/master/proxy-list.txt")
		if err != nil {
			log.Println("Get ip-list.txt failed:", err.Error())
			return pool
		}
		log.Println(resp.String())
		lines := strings.Split(resp.String(), "\n")

		regex := regexp.MustCompile(`.*:[0-9]+ [A-Z]+-[A-Z].*`)

		for _, line := range lines {
			if "" != regex.FindString(line) {
				pool = append(pool, FormatProxy(line))
			} else {
				log.Println(line)
			}
		}
	}

	return pool
}

func FormatProxy(str string) IpProxy {
	proxyParts := strings.Split(str, " ")
	ip := strings.Split(proxyParts[0], ":")[0]
	port := strings.Split(proxyParts[0], ":")[1]
	countryCode := strings.Split(proxyParts[1], "-")[0]
	proxy := IpProxy{
		IP:          ip,
		Port:        port,
		Protocol:    "HTTP",
		CountryCode: countryCode,
	}
	return proxy
}
