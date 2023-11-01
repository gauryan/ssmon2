# SSMON_API

Go언어의 gin 프레임워크를 이용하였고, MySQL Stored Procedure 를 사용하는 범용적인 API 서버이다.

## Air 설치
[Air](https://github.com/cosmtrek/air) : Live reload for Go apps
```bash
$ go install github.com/cosmtrek/air@latest
```


## 배포 (테스트)
```bash
# go 컴파일러 설치
$ cd
$ vi .bashrc
export PATH="$PATH:$HOME/bin/go/bin"
export GOROOT="$HOME/bin/go"
export GOPATH="$HOME/bin/go/gopath"
$ mkdir bin; cd bin
$ wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
$ tar xvfz go1.20.1.linux-amd64.tar.gz
$ mv go1.20.1.linux-amd64 go
$ cd ; source .bashrc
$ go version


# 소스 받아오기 / 빌드
$ git clone https://gitlab.com/hebees_dev/retina_rapi_gin.git -b develop
$ cd retina_rapi_gin/config
$ cp config.go.test config.go
$ cd ..
$ chmod +x *.sh
$ crontab -e
0 1 * * * sh /home/retina/retina_rapi_gin/log.sh

$ go get
$ go build
$ ./start.sh
```


## 배포 (운영)
```bash
# go 컴파일러 설치
$ cd
$ vi .bashrc
export PATH="$PATH:$HOME/bin/go/bin"
export GOROOT="$HOME/bin/go"
export GOPATH="$HOME/bin/go/gopath"
$ mkdir bin; cd bin
$ wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
$ tar xvfz go1.20.1.linux-amd64.tar.gz
$ mv go1.20.1.linux-amd64 go
$ cd ; source .bashrc
$ go version


# 소스 받아오기 / 빌드
$ git clone https://gitlab.com/hebees_dev/retina_rapi_gin.git -b release
$ cd retina_rapi_gin/config
$ cp config.go.real config.go
$ cd ..
$ chmod +x *.sh
$ crontab -e
0 1 * * * sh /home/retina/retina_rapi_gin/log.sh

$ go get
$ go build
$ ./start.sh
```


## 개발서버 실행
```bash
> .\run.bat

or 

$ ./run.sh
```


## 운영서버 실행
```bash
$ ./start.sh
```


## Frontend (fetch) 예시
```javascript
async function PKG_ABC_SP_L_DEF() {
    const req_obj = {
        p_nm: "PKG_ABC.SP_L_DEF",
        in1: "Korea",
        in2: "010-1111-2222",
        in3: 25,
        out1: "cursor",
        out2: "int"
    };

    const result = await exec_request(req_obj);
    console.log("result:", result);
}

//
// Data를 Token I/O에 요청하는 함수
//
export async function exec_request(req_obj) {
    let token = (sessionStorage.getItem("token"))?sessionStorage.getItem("token"):"";
    let token_no = (sessionStorage.getItem("token_no"))?sessionStorage.getItem("token_no"):"";

    let form_data = new FormData();
    form_data.append("data", JSON.stringify(req_obj));
    form_data.append("token", token);
    form_data.append("token_no", token_no);

    try {
        let response = await fetch(TOKEN_IO_URL + "request", {
            method: "POST",
            body: form_data
        });
        let result = await response.json();
        return result.data;
    } catch (xhr) {
        console.error("request 에러:", xhr);
        console.error(req_obj.p_nm + " 에러");
        return null;
    }
}

```