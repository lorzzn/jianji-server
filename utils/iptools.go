package utils

import (
	"errors"
	"jianji-server/config"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"resenje.org/mmdb"
)

func ClientPublicIP(c *gin.Context) (*geoip2.City, error) {
	mmdbPath := "sources/GeoLite2-City.mmdb"
	mmdb.UpdateGeoLite2City(c, mmdbPath, config.MaxMind.LicenseKey)
	if FileExists(mmdbPath) {
		return nil, errors.New("GeoLite2-City.mmdb 下载失败")
	}
	geoDb, err := geoip2.Open("sources/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer func(db *geoip2.Reader) {
		err := db.Close()
		if err != nil {

		}
	}(geoDb)
	record, _ := geoDb.City(net.ParseIP(c.ClientIP()))
	return record, nil
}
