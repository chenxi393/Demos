package main

import (
	"context"
	"fmt"
	_ "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("eazy http request")

func insertUser() {

}

func main() {
	///基础路由
	app := fiber.New()
	ctx := context.Background()

	ctx, span := tracer.Start(ctx, "Get Request", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	app.Get("/nih", func(c *fiber.Ctx) error {
		if span.IsRecording() {
			span.SetAttributes(
				attribute.Int64("id", 123),
				attribute.String("Get", "hello world"),
			)
		}
		fmt.Println(ctx)
		fmt.Println(span)
		return c.SendString("Hello, World!")
	})
	app.Static("hello", "./public/js/1.html")

	log.Fatal(app.Listen(":3333"))

}

func wait() {
	panic("unimplemented")
}

// func handler(w http.ResponseWriter, req *http.Request) {
// 	// Get the context.
// 	ctx := req.Context()

// 	// Get the active span from the context.
// 	span := trace.SpanFromContext(ctx)

// 	// Save the active span in the context.
// 	ctx = trace.ContextWithSpan(ctx, span)

// 	// Add some custom attributes to the span.
// 	span.SetAttributes([]attribute.KeyValue{
// 		attribute.String("customKey", "customValue"),
// 	})

// 	// Print the span information.
// 	fmt.Printf("Span Name: %s\n", span.Name())
// 	fmt.Printf("Span Kind: %s\n", span.SpanKind())

// 	// Continue processing the request.
// 	w.Write([]byte("Hello, World!"))
// }
