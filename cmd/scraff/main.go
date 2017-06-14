package main

import (
	"flag"
	"os"
	"runtime"

	scraff "github.com/realbot/scraff"
)

func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func main() {
	redisURL := flag.String("redis", "localhost:6379", "Redis address")
	flag.Parse()

	extractors := []scraff.AdProvider{
		scraff.ImmobiliareAd{
			Name: "Immobiliare.it",
			Retriever: scraff.AdRetriever{
				Url: "https://www.immobiliare.it/Milano/affitti_appartamenti-Assago.html?criterio=rilevanza",
			},
		},
	}

	ap := scraff.AdProcessor{
		Extractors: extractors,
		Store:      scraff.NewRedisAdStore(*redisURL),
	}
	exitCode := ap.Run()
	os.Exit(exitCode)
}
