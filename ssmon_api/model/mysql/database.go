package mysql

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"ssmon-api/config"

	// "github.com/bdwilliams/go-jsonify/jsonify"
	_ "github.com/go-sql-driver/mysql"
)


var (
	DBConn map[string]*sql.DB
)


func init() {
	DBConn = make(map[string]*sql.DB)
}


// https://pkg.go.dev/database/sql#example-DB.Query-MultipleResultSets
// https://dev.mysql.com/doc/internals/en/multi-resultset.html
func Init(db_config map[string]string) {
	var err error
	// retina:alice.2012@tcp(cooper.pickup9.com:3306)/tokendb
	DSN := db_config["USERNAME"] + ":" + db_config["PASSWORD"] + "@tcp(" + db_config["HOST"] + ":" + db_config["PORT"] + ")/" + db_config["DBNAME"]
	// fmt.Println("DSN: " + DSN)
	DBConn[db_config["CONFIG_NAME"]], err = sql.Open("mysql", DSN)
	// defer conn.Close()
	// check for error
	if err != nil {
		panic("failed to connect database")
	}

	// Connection Pool
	MAX_IDLE_CONNS, _ := strconv.Atoi(db_config["MAX_IDLE_CONNS"])
	MAX_OPEN_CONNS, _ := strconv.Atoi(db_config["MAX_OPEN_CONNS"])
	DBConn[db_config["CONFIG_NAME"]].SetMaxIdleConns(MAX_IDLE_CONNS)
	DBConn[db_config["CONFIG_NAME"]].SetMaxOpenConns(MAX_OPEN_CONNS)
	DBConn[db_config["CONFIG_NAME"]].SetConnMaxLifetime(time.Hour)
	// fmt.Println("Connection Opened to MySQL  Database...: " + config["CONFIG_NAME"] + "\n")
}


func Exec_procedure(db_config map[string]string, data_json string, b64 bool) string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	vec := []string{}
	sp := make_query(data_json, b64)
	if config.IS_LOGGING == true {
		fmt.Print("SP: ")
		fmt.Println(sp)
	}
	value := make_value(data_json)
	if config.IS_LOGGING == true {
		fmt.Print("value: ")
		fmt.Println(value)
	}
	// rows, _ := DBConn[config["CONFIG_NAME"]].Query(sp, "34121", "dfgasf3g")
	rows, err := DBConn[db_config["CONFIG_NAME"]].Query(sp, value...)
	if config.IS_LOGGING == true {
		fmt.Print("rows: ")
		fmt.Println(rows)
		fmt.Print("err: ")
		fmt.Println(err)
	}
	if(err != nil || rows == nil) {
		fmt.Println(`{"RESULT":"`+fmt.Sprint(err)+`"}`)
		return `{"RESULT":"`+fmt.Sprint(err)+`"}`
	}
	if rows != nil {
		defer rows.Close()
	}
	json := Jsonify(rows)
	vec = append(vec, fmt.Sprint(json))

	for rows.NextResultSet() {
		json := Jsonify(rows)
		vec = append(vec, fmt.Sprint(json))
	}
	rows.Close()
	rows = nil

	result := []string{}
	// result := fmt.Sprintln("{")
	result = append(result, "{")
	for i := 0 ; i < len(vec) ; i++ {
		// result += fmt.Sprint("\"out"+strconv.Itoa(i+1)+"\": ")
		result = append(result, `"out`+strconv.Itoa(i+1)+`": `)

		// result += fmt.Sprint(vec[i])
		result = append(result, fmt.Sprint(vec[i]))
		if len(vec) != i+1 {
			// result += fmt.Sprintln(",")
			result = append(result, ",")
		}
		// result += fmt.Sprintln("")
	}
	vec = nil
	
	// result += fmt.Sprint("}")
	result = append(result, "}")
	json_str := strings.Join(result, "")
	result = nil

	if config.IS_LOGGING == true {
		fmt.Print("JSON_STR: ")
		fmt.Println(json_str)
	}

	return json_str
}


// JSON 관련 참고 : https://brownbears.tistory.com/298
// "CALL PKG_USER.SP_L_ADMIN_MEMBER()"
func make_query(data_json string, b64 bool) string {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()

	// var data map[string]interface{}
	var data map[string]any // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
	json.Unmarshal([]byte(data_json), &data)

	var p_nm = ""
	if b64 == true {
		p_nm_byte, _ := base64.StdEncoding.DecodeString(data["p_nm"].(string))
		p_nm = string(p_nm_byte)
	} else {
		p_nm = data["p_nm"].(string)
	}

	sp := "CALL " + p_nm + "("
	for i := 1; i < len(data); i++ {
		_, exists := data["in"+strconv.Itoa(i)]
		if !exists {
			break
		}
		if fmt.Sprint(reflect.TypeOf(data["in"+strconv.Itoa(i)])) == "string" && data["in"+strconv.Itoa(i)].(string) == "" {
			if i == 1 {
				sp += "NULL"
			} else {
				sp += ",NULL"
			}
		} else {
			if i == 1 {
				sp += "?"
			} else {
				sp += ",?"
			}
		}
	}
	sp += ");"

	return sp
}


func make_value(data_json string) []interface{} {
	defer func() {
		r := recover() // 복구 및 에러 메시지 초기화
		if r != nil {
			fmt.Println(r) // 에러 메시지 출력 
		}
	}()
	
	// var data map[string]interface{}
	// var data map[string]string // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
	var data map[string]any // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
	json.Unmarshal([]byte(data_json), &data)

	var result = []interface{}{}

	for i := 1; i < len(data); i++ {
		val, exists := data["in"+strconv.Itoa(i)]
		if !exists {
			break
		}
		if fmt.Sprint(reflect.TypeOf(data["in"+strconv.Itoa(i)])) == "string" && data["in"+strconv.Itoa(i)].(string) == "" {
		} else {
			result = append(result, val)
		}
	}

	// fmt.Println("make_value complete ~!!!")
	// fmt.Print("Result: ")
	// fmt.Println(result)
	return result
}


func Jsonify(rows *sql.Rows) ([]string) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	c := 0
	results := make(map[string]interface{})
	data := []string{}

	for rows.Next() {
		if c > 0 {
			data = append(data, ",")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, value := range values {
			switch value.(type) {
				case nil:
					results[columns[i]] = nil

				case []byte:
					s := string(value.([]byte))
					// x, err := strconv.Atoi(s)
					results[columns[i]] = s

					// if err != nil {
					// 	results[columns[i]] = s
					// } else {
					// 	results[columns[i]] = x
					// }

				default:
					// results[columns[i]] = value
					results[columns[i]] = fmt.Sprint(value)
			}
		}

		b, _ := json.Marshal(results)
		data = append(data, strings.TrimSpace(string(b)))
		c++
	}

	return data
}