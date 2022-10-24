# Go Server 만들면서 배우기

## gvm 설치

nvm같은 역할을 한다. 하지만 gvm을 설치하기 위해서 go가 설치되어있어야한다고 한다.(정말인가?) 확인은 안해봐서 잘 모르겠다. 왜냐하면 나는 이미 homebrew로 go를 설치했기 때문이다.
[gvm](https://github.com/moovweb/gvm)

```shell
// 설치 가능한 go package
$ gvm listall

// go 설치
$ gvm install go1.19.2

// 설치된 go 목록보기
$ gvm list

// 설치된 go 중 사용하고 싶은 거 사용하기
$ gvm use go1.19.2 --default
```

## 프로젝트 세팅

### go mod 및 패키지 설치

```shell
$ mkdir go-server
$ cd go-server
$ go mod init
$ go get "github.com/gofiber/fiber/v2"
$ go get "go.mongodb.org/mongo-driver"
```

## 클린 아키텍쳐로 예제로 배우는 Fiber로 서버 만들기

Fiber [Clean Architecture Recipe](https://github.com/gofiber/recipes/tree/master/clean-architecture)

맛있는 요리인줄 알고 시작했는데 조금 이해가 어려웠다. 보다보니까 책을 사서 읽어보고 싶어졌다. 코드가 매우 아름답다. 각 역할에 대해서 [아티클](https://techblog.woowahan.com/2647/)을 보고 나름 분석을 해보았다.

### entities

entities는 DB와 닿아있다. DB가 어떤 모양이 될 것인지 추상화한 것들이 들어간다.

```go
type Notice struct{
    ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
```

`json`은 json 형식이라는 것을 알려주는 것이다. bson은 json 패키지인듯.

### handlers

핸들러는 가장 고레벨의 함수들이 모여있다. MVC 패턴에서 Controller와 비슷한 역할을 한다.

### pkg

패키지 안에는 repository와 service가 있다.

#### repository

entities를 바라보고 있다. 데이터를 받아서 직접 DB와 소통해서 데이터를 처리하고 결과를 retrun한다.

#### service

가장 추상적이다. entities를 받아 repository를 반환한다.

#### notice

### presenter

### router

### main.go
