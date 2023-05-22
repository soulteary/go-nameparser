package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	python3 "github.com/datadog/go-python3"
	"github.com/gin-gonic/gin"
)

func LoadModule(dir string) *python3.PyObject {
	path := python3.PyImport_ImportModule("sys").GetAttrString("path")
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	return python3.PyImport_ImportModule(filepath.Base(dir))
}

func Convert(input string) string {
	module := LoadModule("./convert")
	function := module.GetAttrString("Convert")
	args := python3.PyTuple_New(1)
	python3.PyTuple_SetItem(args, 0, python3.PyUnicode_FromString(input))
	return python3.PyUnicode_AsUTF8(function.Call(args, python3.Py_None))
}

type HumanName struct {
	Text   string `json:"text"`
	Detail struct {
		Title    string `json:"title"`
		First    string `json:"first"`
		Middle   string `json:"middle"`
		Last     string `json:"last"`
		Suffix   string `json:"suffix"`
		Nickname string `json:"nickname"`
	} `json:"detail"`
}

func Parse(input string) (ret HumanName, err error) {
	var name HumanName
	err = json.Unmarshal([]byte(Convert(input)), &name)
	if err != nil {
		return ret, fmt.Errorf("Parsing JSON failed: %v", err)
	}
	return name, nil
}

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		log.Fatalln("Failed to initialize Python environment")
	}

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

		result, err := Parse(data.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	})

	srv := &http.Server{
		Addr:              ":8080",
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
