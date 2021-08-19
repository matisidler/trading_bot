package calls

import (
	"binance/models"
	"binance/storage"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var (
	apiKey        = "yourApiKey"
	secretKey     = "yourSecretKey"
	LastOrderID   int
	client        = binance.NewClient(apiKey, secretKey)
	futuresClient = binance.NewFuturesClient(apiKey, secretKey)
)

//Creo una nueva orden a market price.
func BuyAtMarketPrice(symbol string, quantity float64) (*binance.CreateOrderResponse, error) {
	//recibe symbol, ej: "BTCUSDT" y cantidad, "0.00450"
	price, err := ConvertPrices(quantity, symbol)
	if err != nil {
		log.Fatalf("%v", err)
	}
	order, err := client.NewCreateOrderService().Symbol(symbol).Side(binance.SideTypeBuy).Type(binance.OrderTypeMarket).Quantity(price).Do(context.Background())
	if err != nil {
		return nil, err
	}
	mod := models.CreateOrderResponse{
		Symbol:                   order.Symbol,
		OrderID:                  int(order.OrderID),
		Side:                     string(order.Side),
		Type:                     string(order.Type),
		Price:                    order.Price,
		ExecutedQuantity:         order.ExecutedQuantity,
		CummulativeQuoteQuantity: order.CummulativeQuoteQuantity,
	}

	db := storage.NewConnection(string(storage.MySQL))
	db.Create(&mod)

	return order, nil
}

func SellAtMarketPrice(symbol string, quantity float64) (*binance.CreateOrderResponse, error) {
	//recibe symbol, ej: "BTCUSDT" y cantidad, "0.00450"
	price, err := ConvertPrices(quantity, symbol)
	if err != nil {
		log.Fatalf("%v", err)
	}
	order, err := client.NewCreateOrderService().Symbol(symbol).Side(binance.SideTypeSell).Type(binance.OrderTypeMarket).Quantity(price).Do(context.Background())
	if err != nil {
		return nil, err
	}
	mod := models.CreateOrderResponse{
		Symbol:                   order.Symbol,
		OrderID:                  int(order.OrderID),
		Side:                     string(order.Side),
		Type:                     string(order.Type),
		Price:                    order.Price,
		ExecutedQuantity:         order.ExecutedQuantity,
		CummulativeQuoteQuantity: order.CummulativeQuoteQuantity,
	}

	db := storage.NewConnection(string(storage.MySQL))
	db.Create(&mod)
	return order, nil
}

//obtener información sobre la orden
func GetOrderById(symbol string, id int64) (*binance.Order, error) {
	order, err := client.NewGetOrderService().Symbol(symbol).
		OrderID(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

//Setear TP y SL
func SellAtLimitPrice(symbol, price, quantity string) (*binance.CreateOrderResponse, error) {
	order, err := client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideTypeSell).Type(binance.OrderTypeLimit).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {

		return nil, err
	}
	return order, nil
}

//Comprar a x precio
func BuyAtLimitPrice(symbol, price, quantity string) (*binance.CreateOrderResponse, error) {
	order, err := client.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {

		return nil, err
	}
	return order, nil
}

func FuturesBuyAtMarketPrice(symbol string, quantity float64) (*futures.CreateOrderResponse, error) {
	//recibe symbol, ej: "BTCUSDT" y cantidad, "0.00450"
	price, err := ConvertPrices(quantity, symbol)
	fmt.Println(price)
	if err != nil {
		log.Fatalf("%v", err)
	}
	order, err := futuresClient.NewCreateOrderService().Symbol(symbol).Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).Quantity(price).PositionSide(futures.PositionSideTypeLong).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

//Convertidor USDT/CRIPTO
func ConvertPrices(usdt float64, symbol string) (string, error) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		return "", err
	}

	for _, p := range prices {
		if p.Symbol == symbol {
			price, err := strconv.ParseFloat(p.Price, 32)
			if err != nil {
				return "", err
			}
			price = usdt / price
			strprice := fmt.Sprintf("%f", price)
			return strprice, nil
		}
	}
	return "", err
}
