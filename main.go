package main

import "github.com/scienceMuseum/content-service/app"

func main() {
   // 애플리케이션 초기화
	application := app.InitializeApp()

	// 서버 실행
	application.StartServer()
}
