package main

import (
	"es-demo/core"
	"es-demo/docs"
	"es-demo/global"
	"es-demo/indexs"
	"log"
)

func main() {
	// 连接es
	core.EsConnect()
	log.Println(global.EsClient)

	// 创建索引
	indexs.CreateIndex()

	docs.DoCreate()
}
