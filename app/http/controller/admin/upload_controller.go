package admin

import (
	"firstProject/app/http/result"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	returnData := result.NewResult(c)
	// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
	// upload是表单字段名字
	file, _ := c.FormFile("upload")
	// 打印上传的文件名
	fmt.Println(file.Filename)

	timeObj := time.Now()
	var str = timeObj.Format("2006-01-02")
	imgPath := GetRandomString(10)
	os.Mkdir("storage/"+str, 0666)
	imgUrl := "storage/" + str + "/" + imgPath + ".jpg"
	// 将上传的文件，保存到./data/1111.jpg 文件中
	c.SaveUploadedFile(file, imgUrl)
	var data struct {
		Url      string `json:"url"`
		Uploaded bool   `json:"uploaded"`
	}
	data.Url = imgUrl
	data.Uploaded = true
	returnData.Success(data)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
