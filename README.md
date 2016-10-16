# jsonmap
An easy way for manipulating json file in your web application developed with go

### I will create API to extract existence JSON data ASAP
### I will provide docs ASAP


simple examplte
~~~ go
package main

import (
	"github.com/Gexkill/jsonmap"
	"log"
)

func main() {
	mappa := jsonmap.CreateJson(jsonmap.Maps{})

	mappa.Set("gino dino pino", []string{
		"start",
		"service1a",
		"-p",
		"50100",
		"-w",
		"low",
	})

	mappa.Set("tino", jsonmap.Maps{
		"ETCD_ADDR":   "",
		"HostIP":      "",
		"INFLUX_USER": "",
		"INFLUX_PWD":  "",
		"INFLUX_ADDR": "",
	})

	mappa.Set("gino lino", []string{
		"start",
		"service1a",
		"-p",
		"50100",
		"-w",
		"low",
	})

	mappa.Set("lino", false)

	mappa.Set("gino dino", []string{
		"start",
		"service1a",
		"-p",
		"50100",
		"-w",
		"low",
	})

	// mappa.Delete("gino dino")

	mappa.PrintSortOrder()
	// mappa.PrintReverseOrder()

	mappa.IndentJsonPrint()

	mappazza := jsonmap.CreateJson(jsonmap.Maps{
		"Name":          "service1a",
		"Type":          "service1a",
		"Image":         "elleflorio/service1a",
		"Remote":        "/gru/pippo/services/service1a",
		"DiscoveryPort": "50100",
		"Analytics": []string{
			"resp_time_ratio",
		},
		"Constraints": jsonmap.Maps{
			"MAX_RESP_TIME": 1500,
		},
		"Configuration": jsonmap.Maps{
			"Env": jsonmap.Maps{
				"ETCD_ADDR":   "",
				"HostIP":      "",
				"INFLUX_USER": "",
				"INFLUX_PWD":  "",
				"INFLUX_ADDR": "",
			},
			"Cmd": []string{
				"start",
				"service1a",
				"-p",
				"50100",
				"-w",
				"low",
			},
			"cpunumber": 1,
			"Ports": jsonmap.Maps{
				"50100": "50100-50104",
			},
			"StopTimeout": 30,
		},
	})

	log.Printf("*****************************************")
	log.Printf("Ricerca valori in Json - BEGIN")
	log.Printf("*****************************************")

	// log.Printf("%v", mappazza.Exist("Configuration"))
	// mappazza.Delete("Constraints")
	// mappazza.IndentJsonPrint()
	mappazza.Search("Configuration", "j")

	log.Printf("*****************************************")
	log.Printf("Ricerca valori in Json - END")
	log.Printf("*****************************************")
}
~~~
