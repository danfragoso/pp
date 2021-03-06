package main

import (
	"database/sql"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	http "github.com/valyala/fasthttp"
)

const ppPort = "3000"

func main() {
	var rootDir string = "."
	if len(os.Args) >= 2 {
		rootDir = os.Args[1]
	}

	db := LoadDatabase(rootDir)
	app := fiber.New()

	app.Get("/ack", func(c *fiber.Ctx) error {
		saveHit(c.Request(), db)
		return c.SendFile(rootDir+"/assets/thumbs_up.svg", true)
	})

	app.Listen(":" + ppPort)
}

func saveHit(req *http.Request, db *sql.DB) {
	db.Exec(`
		INSERT INTO hits (method, user_agent, host) 
		VALUES (?, ?, ?)`,

		string(req.Header.Method()),
		string(req.Header.UserAgent()),
		string(req.Host()),
	)
}
