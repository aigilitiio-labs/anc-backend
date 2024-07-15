package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrderRequest is used to receive order creation requests
type OrderRequest struct {
    ClientOrderID   string `json:"client_order_id"`
    ProductID       string `json:"product_id"`
    Side            string `json:"side"`
    OrderConfiguration struct {
        MarketMarketIOC struct {
            BaseSize  string `json:"base_size,omitempty"`
            QuoteSize string `json:"quote_size,omitempty"`
        } `json:"market_market_ioc,omitempty"`
        LimitLimitGTC struct {
            BaseSize   string `json:"base_size,omitempty"`
            LimitPrice string `json:"limit_price,omitempty"`
        } `json:"limit_limit_gtc,omitempty"`
    } `json:"order_configuration"`
}

// Order is used to store orders in the database
type Order struct {
    ID                primitive.ObjectID `bson:"_id,omitempty"`
    ClientOrderID     string             `bson:"client_order_id"`
    ProductID         string             `bson:"product_id"`
    Side              string             `bson:"side"`
    OrderConfiguration struct {
        MarketMarketIOC struct {
            BaseSize  string `bson:"base_size,omitempty"`
            QuoteSize string `bson:"quote_size,omitempty"`
        } `bson:"market_market_ioc,omitempty"`
        LimitLimitGTC struct {
            BaseSize   string `bson:"base_size,omitempty"`
            LimitPrice string `bson:"limit_price,omitempty"`
        } `bson:"limit_limit_gtc,omitempty"`
    } `bson:"order_configuration"`
    Status string `bson:"status"`
}

// ToOrder converts an OrderRequest to an Order
func (or OrderRequest) ToOrder() Order {
    return Order{
        ClientOrderID: or.ClientOrderID,
        ProductID:     or.ProductID,
        Side:          or.Side,
        OrderConfiguration: struct {
            MarketMarketIOC struct {
                BaseSize  string `bson:"base_size,omitempty"`
                QuoteSize string `bson:"quote_size,omitempty"`
            } `bson:"market_market_ioc,omitempty"`
            LimitLimitGTC struct {
                BaseSize   string `bson:"base_size,omitempty"`
                LimitPrice string `bson:"limit_price,omitempty"`
            } `bson:"limit_limit_gtc,omitempty"`
        }{
            MarketMarketIOC: struct {
                BaseSize  string `bson:"base_size,omitempty"`
                QuoteSize string `bson:"quote_size,omitempty"`
            }{
                BaseSize:  or.OrderConfiguration.MarketMarketIOC.BaseSize,
                QuoteSize: or.OrderConfiguration.MarketMarketIOC.QuoteSize,
            },
            LimitLimitGTC: struct {
                BaseSize   string `bson:"base_size,omitempty"`
                LimitPrice string `bson:"limit_price,omitempty"`
            }{
                BaseSize:   or.OrderConfiguration.LimitLimitGTC.BaseSize,
                LimitPrice: or.OrderConfiguration.LimitLimitGTC.LimitPrice,
            },
        },
        Status: "New",
    }
}