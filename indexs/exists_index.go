package indexs

import (
	"context"
	"es-demo/global"
)

// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := global.EsClient.IndexExists(index).Do(context.Background())
	return exists
}
