package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/lkdevelopment/hetzner-cloud-api-mock"
	"github.com/lkdevelopment/hetzner-cloud-api-mock/apiblueprint"
	"github.com/lkdevelopment/hetzner-cloud-api-mock/model"
	"github.com/lkdevelopment/hetzner-cloud-api-mock/server"
)

var (
	flagPath = flag.String(
		"path",
		envString("PATH", "/var/mock-server/"),
		"Path where the mock server will store temporary files",
	)
	flagPort = flag.String(
		"port",
		envString("PORT", "8080"),
		"Port where the mock server will launch on.",
	)
	flagSource = flag.String(
		"source",
		envString("source", "https://docs.hetzner.cloud/docs.apib"),
		"Port where the server will launch on.",
	)
	flagDebug = flag.Bool(
		"debug",
		false,
		"Add a log line for every request",
	)
	flagLocal = flag.Bool(
		"local",
		false,
		"Use local version of as source instead of downloading a new one. (Need a docs.apib in the directory. When you want to run local version of the docs use --path . and this flag.)",
	)
)

func main() {
	flag.Parse()
	path := *flagPath
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	log.Printf("Welcome to the API Mock Server %s\n", api_mock.Version)
	log.Printf("I will store all files locally at %v\n", path)
	if *flagLocal {
		log.Printf("I will get my data from %v\n", path+"docs.apib")
	} else {
		log.Printf("I will get my data from %v\n", *flagSource)
		if err := DownloadFile(path+"docs.apib", *flagSource); err != nil {
			panic(err)
		}
	}

	parser := apiblueprint.Transformer{
		Path:                 path,
		ApiBlueprintFileName: "docs.apib",
		JSONFileName:         "result.json",
	}
	err := parser.TransformAPIBlueprintFile()
	if err != nil {
		panic(err)
	}
	if len(parser.Resources) > 0 {
		transformer := model.Transformer{}
		apiHandlers := transformer.Transform(parser.Resources)
		server := server.Server{
			Handler:     apiHandlers,
			Port:        *flagPort,
			APIFilePath: path + "docs.apib",
			DebugMode:   *flagDebug,
		}
		if *flagDebug {
			log.Println("[Running in Debug Mode]")
		}
		server.Start()
	}

}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// envString get a ENV Var with fallback
// copied from DAS
func envString(name, fallback string) string {
	if s := os.Getenv("API_MOCK_" + name); s != "" {
		return s
	}
	return fallback
}
