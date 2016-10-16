# jsonmap
An easy way to handle json file in your web application developed with go

### I will create API to extract existence JSON data ASAP
### I will provide docs ASAP


###1. EXAMPLE
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
		"string0",
		"string1",
		"string2",
		"string3",
		"string4",
		"string5",
	})

	mappa.Set("lino", false)

	mappa.Set("gino dino", []string{
		"string0",
		"string1",
		"string2",
		"string3",
		"string4",
		"string5",
	})

	mappa.Delete("gino dino")

	mappa.PrintSortOrder()

	mappa.IndentJsonPrint()
}
~~~

###2. EXAMPLE
~~~ go
package main

import (
	"github.com/Gexkill/jsonmap"
	"log"
)

func main() {
	mappazza := jsonmap.CreateJson(jsonmap.Maps{
		"Name":          "service1a",
		"Type":          "service1a",
		"Image":         "pino/service1a",
		"Remote":        "services",
		"DiscoveryPort": "50100",
		"Analytics": []string{
			"resp_time_ratio",
		},
		"Constraints": jsonmap.Maps{
			"MAX_RESP_TIME": 1500,
		},
		"Configuration": jsonmap.Maps{
			"Env": jsonmap.Maps{
				"ADDRESS":   "",
				"Host":      "",
				"USER": "",
				"PASSROD":  "",
				"ADDR": "",
			},
			"Cmd": []string{
				"string0",
				"string1",
				"string2",
				"string3",
				"string4",
				"string5",
			},
			"cpunumber": 1,
			"Ports": jsonmap.Maps{
				"50100": "50100-50104",
			},
			"StopTimeout": 30,
		},
	})

	mappazza.Search("Configuration", "j")
}
~~~
