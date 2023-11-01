package token_db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"ssmon-api/config"
	"ssmon-api/model/mysql"

	"github.com/buger/jsonparser"
)

// 토큰 생성 (64자리)
func Make_token() string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

    // return crypto.randomBytes(16).toString('hex') + new Date().getTime().toString(); // 45자리
	rand.Seed(time.Now().UnixNano())
	rand_num := fmt.Sprintf("%d", rand.Intn(1000))
	now := time.Now().Unix()
	hash_val := sha256.Sum256([]byte(fmt.Sprint(now)))
	hash_val_str := fmt.Sprintf("%x", hash_val)
	hash_val_str = hash_val_str + rand_num

	return hash_val_str
}


// 신규 토큰 등록
func Insert_token(token string, token_no string) string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	return Sp_token("SP_I_TOKEN_V2", token, token_no)
}


// 토큰 체크
func Check_token(token string, token_no string) string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	return Sp_token("SP_U_TOKEN_V2", token, token_no)
}


// 토큰을 위한 SP 실행 함수
func Sp_token(p_nm string, token string, token_no string) string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	param := make(map[string]string)
	param["p_nm"] = p_nm
	param["in1"]  = token_no
	param["in2"]  = token
	data_json, _ := json.Marshal(param)
	if config.IS_LOGGING == true {
		fmt.Print("data_json: ")
		fmt.Println(string(data_json))
	}

	result := mysql.Exec_procedure(config.TOKEN_DB_CONFIG, string(data_json), false)
	if result == "" {
		fmt.Println("Exec_procedure 에러")
		return "0" // 에러가 발생한 경우, "0" 을 리턴
	}
	if config.IS_LOGGING == true {
		fmt.Print("result: ")
		fmt.Println(result)
	}

	data := [][]map[string]interface{}{}
	err := json.Unmarshal([]byte(result), &data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unmarshal 에러")
		return "0" // 에러가 발생한 경우, "0" 을 리턴
	}
	// fmt.Print("Data: ")
	// fmt.Println(data)
	// fmt.Print("O_ROWCOUNT: ")
	// fmt.Println(data[0][0]["O_ROWCOUNT"])

	// 성공인 경우 "1"을 리턴
	return fmt.Sprint(data[0][0]["O_ROWCOUNT"])
}


// PASS_SP 체크하기
// PASS : 해당SP에 대해서는 토큰 체크하지 않고 통과시킨다. (기존의 IGNORE 와 같은 개념)
func Check_pass_sp(pass_list map[string]int, p_nm string) int {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()
	
	val, exists := pass_list[p_nm]
	if exists == true {
		return val
	} else {
		return -1
	}
}


// SEND_SP 체크하기
// SEND : 해당 SP만 실행 (토큰체크는 한다)
func Check_send_sp(send_list map[string]int, p_nm string) int {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()
	
	val, exists := send_list[p_nm]
	if exists == true {
		return val
	} else {
		return -1
	}
}


// PASS/SEND SP 목록 가져오기
func Set_token_sp() {
	param := make(map[string]string)
	param["p_nm"] = "SP_L_TOKEN_SP"
	data_json, _ := json.Marshal(param)

	result := mysql.Exec_procedure(config.TOKEN_DB_CONFIG, string(data_json), false)
	if result == "" {
		fmt.Println("Exec_procedure 에러")
		return
	}
	data_array, _, _, _ := jsonparser.Get([]byte(result), "[0]")

	jsonparser.ArrayEach(data_array, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dbms, _    := jsonparser.GetString(value, "DBMS")
		sp_name, _ := jsonparser.GetString(value, "SP_NAME")
		sp_type, _ := jsonparser.GetString(value, "SP_TYPE")

		if string(dbms) == "oracle" {
			if string(sp_type) == "pass" {
				config.ORACLE_PASS_LIST[sp_name] = 1
			} else if string(sp_type) == "send" {
				config.ORACLE_SEND_LIST[sp_name] = 1
			}
		} else if string(dbms) == "mysql" {
			if string(sp_type) == "pass" {
				config.MYSQL_PASS_LIST[sp_name] = 1
			} else if string(sp_type) == "send" {
				config.MYSQL_SEND_LIST[sp_name] = 1
			}
		}
	})
}
