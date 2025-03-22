package core

import (
	"es-demo/global"
	"github.com/olivere/elastic/v7"
)

// EsConnect 用于建立与Elasticsearch的连接。
// 该函数不接收任何参数，也不返回任何值。
// 它通常用于初始化Elasticsearch客户端，以便后续的搜索或索引操作。
func EsConnect() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		panic(err)
		return
	}

	// 将全局变量 global.EsClient 设置为传入的 Elasticsearch 客户端实例，确保全局范围内可以访问该客户端。
	global.EsClient = client
}
