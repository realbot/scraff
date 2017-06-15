package main

import (
	"flag"
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

	providers := []scraff.AdProvider{
		scraff.ImmobiliareAd{
			Retriever: scraff.AdRetriever{
				Url: "https://www.immobiliare.it/Milano/affitti_appartamenti-Assago.html?criterio=rilevanza",
			},
		},
	}

	ap := scraff.AdProcessor{
		Providers: providers,
		Store:     scraff.NewRedisAdStore(*redisURL),
	}
	ap.Run()
}
