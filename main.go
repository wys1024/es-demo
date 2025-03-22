package main

import (
	"es-demo/core"
	"es-demo/global"
	"log"
)

func main() {
	core.EsConnect()
	log.Println(global.EsClient)
}
