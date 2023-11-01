package main

import (
	"fmt"
	"runtime"

	"ssmon-api/config"
	"ssmon-api/model/mysql"
	"ssmon-api/routes"
)

func main() {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	num_cpu := runtime.NumCPU()
	if num_cpu > 1 {
		num_cpu = num_cpu - 1
	}
	runtime.GOMAXPROCS(num_cpu)
	app := routes.Router()

	if config.START_ACTIVE["SSMON"] == true {
		mysql.Init(config.TOKEN_DB_CONFIG)	// TOKEN DB 커넥션풀 설정
		mysql.Init(config.SSMON_DB_CONFIG)	// SSMON DB 커넥션풀 설정
	}

	// token_db.Set_token_sp() // PASS/SEND SP 목록 가져오기
	
	app.Run(":" + config.APP_PORT)
}

