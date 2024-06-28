package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//// Load the Shared AWS Configuration (~/.aws/config)
	//cfg, err := config.LoadDefaultConfig(context.TODO())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("AWS Config loaded successfully")
	//fmt.Println(cfg)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Up",
		})
	})
	r.Run(":8080")
}
