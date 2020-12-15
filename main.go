package main

import (
	"GoGinBlog/routes"
)

func main() {
	routes.InitRoutes()
}


//
//func CheckError(err error) bool {
//	if err != nil {
//		_, file, line, ok := runtime.Caller(1)
//		if ok {
//			emsg := fmt.Sprintf("file:%s, line:%d, error:%s", file, line, err.Error())
//			log.Info("Upload sitemap success, %s", emsg)
//		} else {
//			log.Info("Upload sitemap success, %s", err)
//		}
//		return true
//	}
//	return false
//}