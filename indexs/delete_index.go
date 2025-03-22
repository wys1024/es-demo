package indexs

import (
	"context"
	"es-demo/global"
	"fmt"
)

func DeleteIndex(index string) {
	_, err := global.EsClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index, "索引删除成功")
}
