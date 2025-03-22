package docs

import (
	"context"
	"es-demo/global"
	"es-demo/models"
	"fmt"
	"time"
)

func DoCreate() {
	user := models.UserModel{
		UserName:  "admin",
		NickName:  "管理员",
		ID:        1,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	indexResponse, err := global.EsClient.Index().Index(user.IndexName()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", indexResponse.Result)
}
