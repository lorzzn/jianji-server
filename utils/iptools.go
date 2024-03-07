package utils

import (
	"errors"
	"jianji-server/config"
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
	success, _ := mmdb.UpdateGeoLite2City(ctx, mmdbPath, config.MaxMind.LicenseKey)
	return success
}

func ClientPublicIP(c *gin.Context) (*geoip2.City, error) {
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

		}
	}(geoDb)
	record, _ := geoDb.City(net.ParseIP(c.ClientIP()))
	return record, nil
}
