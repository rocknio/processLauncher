package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	FlagsInit()

	flag.Parse()

	if H {
		flag.Usage()
		return
	}

	for {
		RunProgress()
		log.Printf("Process [%s] Exited, restart it ......", Cmd)
		time.Sleep(3 * time.Second)
	}
}
