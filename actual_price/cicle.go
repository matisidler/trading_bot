package actual_price

import (
	"binance/calls"
	"binance/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go"
)

func GetPrice() {
	var status bool
	for {

		cfg := finnhub.NewConfiguration()
		cfg.AddDefaultHeader("X-Finnhub-Token", "c4d4lvqad3icnt8r9eng")
		finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

		ahora := time.Now()
		btc10days, _, err := finnhubClient.CryptoCandles(context.Background(), "BINANCE:BTCUSDT", "1", ahora.Add(10*time.Minute*-1).Unix(), time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		}

		closePrice := btc10days.C
		var mediaMovil10 float32
		for _, valor := range closePrice {
			mediaMovil10 = mediaMovil10 + valor
		}
		mediaMovil10 = mediaMovil10 / float32(len(closePrice))
		fmt.Printf("simple moving average 10 periods: %v\n", mediaMovil10)

		btc20days, _, err := finnhubClient.CryptoCandles(context.Background(), "BINANCE:BTCUSDT", "1", ahora.Add(20*time.Minute*-1).Unix(), time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		}
		closePrice = btc20days.C
		var mediaMovil20 float32
		for _, valor := range closePrice {
			mediaMovil20 = mediaMovil20 + valor
		}
		mediaMovil20 = mediaMovil20 / float32(len(closePrice))
		fmt.Printf("simple moving average 20 periods: %v\n", mediaMovil20)

		if mediaMovil10 > mediaMovil20 {
			if status != true {
				fmt.Println("BUY")
				order, err := calls.BuyAtMarketPrice("BTCUSDT", 10)
				if err != nil {
					panic(err)
				}
				fmt.Println(order)
				status = true
			} else if status == true {
				fmt.Println("YOU'RE LONG")
				calls.BuyAtMarketPrice("BTCUSDT", 10)
			}

		}

		if mediaMovil10 < mediaMovil20 {
			if status != false {
				fmt.Println("SELL")
				order, err := calls.SellAtMarketPrice("BTCUSDT", 10)
				if err != nil {
					panic(err)
				}
				fmt.Println(order)
				status = false
			} else if status == false {
				fmt.Println("YOU'RE SHORT")
			}

		}

		btcvolumeInfo, _, err := finnhubClient.CryptoCandles(context.Background(), "BINANCE:BTCUSDT", "D", ahora.Add(24*time.Hour*-1).Unix(), time.Now().Unix())

		if err != nil {
			fmt.Println(err)
		}
		btcvolume := btcvolumeInfo.V
		res, err := http.Get("http://finnhub.io/api/v1/quote?symbol=BINANCE:BTCUSDT&token=c4d4lvqad3icnt8r9eng")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		data := models.FinnhubResponse{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("----------BINANCE:BTCUSDT-----------\n, current price: %s\n, change: %s\n, percent change: %s\n, highest price of day: %s\n, lowest price of day: %s\n, open price of day: %s\n, previous price close: %s\n, volume: %v\n \n---------------------------\n", data.CurrentPrice, data.Change, data.PercentChange, data.HighPriceOfDay, data.LowPriceOfDay, data.OpenPriceOfDay, data.PreviousPriceClose, btcvolume)

		time.Sleep(5 * time.Second)
	}

}
