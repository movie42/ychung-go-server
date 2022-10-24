# Q1.어덯게하면 router를 묶을 수 있을까요?

안녕하세요. go로 만들어보면서 배우고 있는 go린이입니다.
golang fiber로 API 서버를 만들고 있는데요.
질문 : express처럼 router를 나누는 방법이 감이 안잡혀서 질문드립니다. go fiber에서 router를 설계할 때 어떻게 코드를 작성하시나요?
express에서는 아래 코드를 작성하면 router를 분리할 수 있었습니다.
(url : localhost:3000/api/notices로 호출할 수 있다.)
// app.js
const app = express();

app.use("/api", apiRouter);

// api.router.js

const api = express.Router();

api.use("/notices", NoticeRouter);

// notice.router.js
const notice = express.Router();

notice.get("/", getNotices);
go fiber에서는 아래와 같이 작성했습니다.
(url : localhost:3000/api/notices로 호출할 수 있다.)
// main.go
func main(){
app := fiber.New()
router.ApiRouter(app)
}

// api.router.go
func ApiRouter(app \*fiber.App){
api := fiber.New()
NoticeRouter(api.Group("/notices"))
app.Mount("/api", api)
}

// notice.router.go
func NoticeRouter(router fiber.router){
router.Get("/", func(c \*fiber.Ctx)error{
return c.SendString("동작")
})
}
동작은 합니다. 그런데 제가 궁금한 부분은 위의 코드가 조금 직관적으로 느껴지지 않아서 혹시 다른 방법으로 Router를 나눌 수 있는 방법이 있을까 하여 질문드립니다.
router를 설계할 때 어떻게 코드를 작성하시나요? (edited)

# Q2. 혹시 repository를 묶을 수 있는 방법이 있을까요?
