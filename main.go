package main

import (
	"fmt"
	"requestbingo/models"
	_ "requestbingo/models"
	_ "requestbingo/routers"
	//"github.com/astaxie/beego"
)

//func main() {
//	beego.Run()
//}

func main() {
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
	fmt.Println(models.TinyId4())
}
