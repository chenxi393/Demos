package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"net/http"

	//"go.opentelemetry.io/otel"
	//"go.opentelemetry.io/otel/trace"

	"github.com/gofiber/fiber/v2"
)

func AppInit(app *fiber.App) {

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	}).Name("api")

	///路由命名
	data, _ := json.MarshalIndent(app.GetRoute("api"), "", "  ")
	fmt.Print(string(data))

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	//和上面get会有点冲突 估计会有优先级的
	///静态文件服务
	app.Static("/", "./public")
	// => http://localhost:3000/js/script.js
	// => http://localhost:3000/css/style.css

	//使用左边路径映射到右边
	app.Static("/prefix", "./public")
	// => http://localhost:3000/prefix/js/script.js
	// => http://localhost:3000/prefix/css/style.css

	//任意路径访问都能映射到后面1.html
	app.Static("*", "./public/js/1.html")
	// => http://localhost:3000/any/path/shows/index/html

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	///中间件和 Next
	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥈 Second handler")
		return c.Next()
	})

	// GET /api/register
	app.Get("/api/list", func(c *fiber.Ctx) error {
		fmt.Println("🥉 Last handler")
		return c.SendString("Hello, World 👋!")
	})
}

// 可以有很大trace 但是最好一个应用一个全局的trace
//var tracer = otel.Tracer("app_or_package_name")

// span 存储在标准库的 context.Context里
func handler(w http.ResponseWriter, req *http.Request) {
	// Get the context.
	ctx := req.Context()

	// Set the context. This is usually done by instrumentations, for example, otelhttp.
	req = req.WithContext(ctx)
	//anotherHandler(w, req)
}
/*
func insertUser(ctx context.Context, user *User) error {
	//第二步 开启一个Trace
	ctx, span := tracer.Start(ctx, "insert-user", trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	//第一步写数据库的插入操作
	if _, err := db.NewInsert().Model(user).Exec(ctx); err != nil {
		//第三步 记录错误 和设置状态码
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	//第四步 记录上下文信息
	if span.IsRecording() {
		span.SetAttributes(
			attribute.Int64("enduser.id", user.ID),
			attribute.String("enduser.email", user.Email),
		)
	}

	return nil
}
*/
