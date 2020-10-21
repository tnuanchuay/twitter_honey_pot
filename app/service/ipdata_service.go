package service

import (
	"errors"
	"fmt"
	"github.com/ipdata/go"
	"log"
	"net"
	"os"
	"strings"
)

func GetIpData(ips string) (*ipdata.IP, error) {
	if strings.Contains(ips, ":") {
		ips = strings.Split(ips, ":")[0]
	}

	ip := net.ParseIP(ips)
	if ip.To4() == nil {
		return nil, errors.New(fmt.Sprintf("%s is not ipv4", ips))
	}

	client, err := ipdata.NewClient(os.Getenv("IPDATA_TOKEN"))
	if err != nil {
		return nil, err
	}
	log.Println(ips)
	ipLookup, err := client.Lookup(ips)

	return &ipLookup, err
}
