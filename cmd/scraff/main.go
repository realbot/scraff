package main

import (
	"flag"
	"runtime"
	"time"

	"github.com/realbot/scraff"
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
	mjPublicKey := flag.String("mjPublicKey", "", "Mailjet Public Key")
	mjPrivateKey := flag.String("mjPrivateKey", "", "Mailjet Private Key")

	flag.Parse()

	providers := []scraff.AdProvider{
		scraff.NewImmobiliareAdProvider(
			"Immobiliare Assago",
			"https://www.immobiliare.it/Milano/affitti_appartamenti-Assago.html?criterio=rilevanza"),
		scraff.NewImmobiliareAdProvider(
			"Immobiliare Buccinasco",
			"https://www.immobiliare.it/Milano/affitti_appartamenti-Buccinasco.html?criterio=rilevanza"),
	}

	ap := scraff.NewAdProcessor(
		providers,
		scraff.NewRedisAdStore(*redisURL),
		scraff.NewMailjetAdSender(*mjPublicKey, *mjPrivateKey),
		3*time.Hour)
	ap.Run()
}
