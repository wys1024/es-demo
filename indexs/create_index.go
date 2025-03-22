package indexs

import (
	"context"
	"es-demo/global"
	"es-demo/models"
	"fmt"
)

func CreateIndex() {
	index := "user_index"
	if ExistsIndex(index) {
		// 先删除再创建
		DeleteIndex(index)
	}
	createIndex, err := global.EsClient.CreateIndex(index).BodyString(models.UserModel{}.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createIndex.Index, "索引创建成功")

}
