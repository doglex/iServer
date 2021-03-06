package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./static/fav.ico")
	r.StaticFile("/wg", "./static/wg.html")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello work",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/reset_wg", func(c *gin.Context) {
		command := "/opt/wireguard.sh -n && /opt/wireguard.sh -s"
		fmt.Println(command)
		cmd := exec.Command("/bin/bash", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			c.JSON(200, gin.H{
				"message": "error",
			})
			return
		}
		fmt.Printf("Execute Shell:%s finished")
		c.JSON(200, gin.H{
			"message": "reset",
		})
	})
	r.GET("/show_wg", func(c *gin.Context) {
		command := "cat /etc/wireguard/wg0_client "
		fmt.Println(command)
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			c.JSON(200, gin.H{
				"message": "error",
			})
			return
		}
		fmt.Printf("Execute Shell:%s finished")
		c.JSON(200, gin.H{
			"message": output,
		})
	})

	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
