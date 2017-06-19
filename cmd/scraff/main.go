package main

import (
	"flag"
	"runtime"
	"time"

	"github.com/golang/glog"
	. "github.com/realbot/scraff"
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

	if *mjPublicKey == "" {
		glog.Fatal("Mailjet Public Key ID required")
	}

	if *mjPrivateKey == "" {
		glog.Fatal("Mailjet Private Key ID required")
	}

	providers := []AdProvider{
		NewImmobiliareAdProvider(
			"Immobiliare Assago",
			"https://www.immobiliare.it/Milano/affitti_appartamenti-Assago.html?criterio=rilevanza"),
		NewImmobiliareAdProvider(
			"Immobiliare Buccinasco",
			"https://www.immobiliare.it/Milano/affitti_appartamenti-Buccinasco.html?criterio=rilevanza"),
		NewIdealistaAdProvider(
			"Idealista Assago",
			"https://www.idealista.it/affitto-case/assago-milano/"),
		NewIdealistaAdProvider(
			"Idealista Buccinasco",
			"https://www.idealista.it/affitto-case/buccinasco-milano/"),
	}

	ap := NewAdProcessor(
		providers,
		NewRedisAdStore(*redisURL),
		NewMailjetAdSender(*mjPublicKey, *mjPrivateKey),
		3*time.Hour)
	ap.Run()
}
