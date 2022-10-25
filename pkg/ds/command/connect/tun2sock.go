package connect

import (
	"fmt"
	"time"

	"github.com/wwqdrh/logger"
	"github.com/wwqdrh/nettool/common"
	"github.com/wwqdrh/nettool/device/tun"
	"github.com/wwqdrh/nettool/server/https"
)

func UpdateRoute() {
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()
	for range timer.C {
		ips, err := getAllServiceIp()
		if err != nil {
			logger.DefaultLogger.Error(err.Error())
		}
		setroute(ips)
	}
}

func getAllServiceIp() ([]string, error) {
	res, err := https.DoReq(&https.HTTPOpt{
		Method: "GET",
		Url:    "/service/dnsmap",
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Body)
	return nil, nil
}

func setroute(ips []string) {
	routes := map[string]struct{}{}
	for _, item := range ips {
		netpart, err := common.IpNetPart(item)
		if err != nil {
			logger.DefaultLogger.Warn(err.Error())
		}
		routes[netpart] = struct{}{}
	}

	routesArr := []string{}
	for item := range routes {
		routesArr = append(routesArr, item)
	}
	tun.Ins().SetRoute(routesArr, nil)
}
