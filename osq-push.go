// example from infludb client

package main

import (
	"fmt"
	"github.com/influxdb/influxdb/client"
	"log"
	"net/url"
	"os"
	"os/exec"
)

const (
	MyHost        = "localhost"
	MyPort        = 8086
	MyDB          = "serverdata"
	MyMeasurement = "processes"
)

func main() {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", MyHost, MyPort))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *u,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	dur, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Happy as a Hippo! %v, %s", dur, ver)

	sql := "\"SELECT pid, parent, name, resident_size, phys_footprint, user_time, system_time FROM processes;\""
	data := runQuery(sql)
	// fmt.Println(data)

	/*
		s := []*influxdb.Series{{
						Name: "api_access",
						Columns: []string{
							"status", "latency", "value", "query", "app_id",
						},
						Points: [][]interface{}{
							{status, latency, cCopy.Request.URL.Path, cCopy.Request.URL.RawQuery, appId},
						},
					}}

					// time.Sleep(time.Second * 10)
					err := influxdbC.WriteSeries(s)

	*/

}

func runQuery(sql string) string {
	app := "osqueryi"
	//	sql =  +  sql
	fmt.Println(sql)
	//cmd := exec.Command(app," --json ", sql)
	cmd := app + " --json " + sql

	out, err := exec.Command("sh", "-c", cmd).Output()
	//out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "error"
	}
	return string(out)

}
