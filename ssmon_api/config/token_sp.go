package config

// PASS : 해당SP에 대해서는 토큰 체크하지 않고 통과시킨다. (기존의 IGNORE 와 같은 개념)
// SEND : 해당 SP만 실행 (토큰체크는 한다)

var ORACLE_PASS_LIST = map[string]int{
}

var MYSQL_PASS_LIST = map[string]int{
}

var ORACLE_SEND_LIST = map[string]int{
}

var MYSQL_SEND_LIST = map[string]int{
}
