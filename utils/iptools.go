package utils

import (
	"errors"
	"jianji-server/config"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"golang.org/x/net/context"
	"resenje.org/mmdb"
)

var (
	mmdbPath = "sources/GeoLite2-City.mmdb"
	ctx      = context.Background()
)

func SetupMmdb() bool {
	log.Println("更新 geoip 数据库...")
	success, err := mmdb.UpdateGeoLite2City(ctx, mmdbPath, config.MaxMind.LicenseKey)
	if err != nil {
		log.Println("更新失败: ", err)
	} else {
		log.Println("更新成功")
	}
	return success
}

func GetClientIP(c *gin.Context) net.IP {
	return net.ParseIP(c.ClientIP())
}

func GetIPGeoRecord(ip net.IP) (*geoip2.City, error) {
	if FileExists(mmdbPath) {
		return nil, errors.New("GeoLite2-City.mmdb 不存在")
	}
	geoDb, err := geoip2.Open("sources/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer func(db *geoip2.Reader) {
		err := db.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(geoDb)
	record, _ := geoDb.City(ip)
	return record, nil
}
