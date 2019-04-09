package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func main() {

	var inputip string
	inputip = os.Args[1]

	db, err := geoip2.Open("GeoLite2-City_20190409/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(inputip)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP  : %v\n", inputip)
	fmt.Printf("国家: %v %v %v\n", record.Country.Names["zh-CN"], record.Country.Names["en"], record.Country.IsoCode)

	if len(record.Subdivisions) >= 1 {
		fmt.Printf("城市: %v %v\n", record.Subdivisions[0].Names["zh-CN"], record.Subdivisions[0].Names["en"])
		fmt.Printf("时区: %v\n", record.Location.TimeZone)
	}

	fmt.Printf("坐标: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

}
