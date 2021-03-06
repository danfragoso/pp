package main

import (
	"database/sql"

	fiber "github.com/gofiber/fiber/v2"
	http "github.com/valyala/fasthttp"
)

const ppPort = "3000"

func main() {
	db := LoadDatabase()

	app := fiber.New()

	app.Get("/ack", func(c *fiber.Ctx) error {
		saveHit(c.Request(), db)
		return c.SendFile("./assets/thumbs_up.svg", true)
	})

	app.Listen(":" + ppPort)
}

func saveHit(req *http.Request, db *sql.DB) {
	insert, _ := db.Prepare(`
		INSERT INTO hits (method, user_agent, host) 
		VALUES (?, ?, ?)
	`)

	insert.Exec(
		string(req.Header.Method()),
		string(req.Header.UserAgent()),
		string(req.Host()),
	)
}
