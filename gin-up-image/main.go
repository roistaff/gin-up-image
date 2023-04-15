package main
import(
	"os"
	"github.com/gin-gonic/gin"
)
type File struct{
	Base64Text	string	`json:"base64Text"`
}
func post (c *gin.Context){
	var image File
	c.BindJSON(&image)
	textlog(image.Base64Text)
}
func textlog(text string){
	bytes := []byte(text)
	if err := os.WriteFile("out.txt",bytes,0666); err != nil {
		panic(err)
	}
}
func main(){
	app := gin.Default()
	app.POST("/post",post)
	app.Run()
}
