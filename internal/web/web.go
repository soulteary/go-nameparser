package web

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/go-nameparser/internal/bridge"
	"github.com/soulteary/go-nameparser/internal/define"
)

func Launch() {
	gin.SetMode(gin.ReleaseMode)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	const ProjectInfo = `project: <a href="https://github.com/soulteary/go-nameparser">soulteary/go-nameparser</a>`

	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(ProjectInfo))
	})

	type Data struct {
		Name string `json:"name"`
	}

	route.POST("/api/convert", func(c *gin.Context) {
		var data Data
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := bridge.Parse(data.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	})

	srv := &http.Server{
		Addr:              define.WEB_PORT,
		Handler:           route,
		ReadHeaderTimeout: time.Second * 10,
		ReadTimeout:       time.Second * 10,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Program start error: %s\n", err)
		}
	}()
	log.Println("soulteary/go-nameparser has started 🚀")

	<-ctx.Done()

	stop()
	log.Println("The program is closing, if you want to end it immediately, please press `CTRL+C`")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Program was forced to close: %s\n", err)
	}

	log.Println("Look forward to meeting you again ❤️")
}
