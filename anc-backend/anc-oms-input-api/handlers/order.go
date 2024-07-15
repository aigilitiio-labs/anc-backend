package handlers

import (
	"anc-oms-input-api/common"
	"anc-oms-input-api/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(c *gin.Context) {
    var orderRequest model.OrderRequest
    if err := c.ShouldBindJSON(&orderRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order := orderRequest.ToOrder()
    order.ID = primitive.NewObjectID()

    collection := common.GetCollection("orders")
    _, err := collection.InsertOne(context.Background(), order)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func CancelOrder(c *gin.Context) {
    var orderRequest struct {
        ClientOrderID string `json:"client_order_id"`
    }
    if err := c.ShouldBindJSON(&orderRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    collection := common.GetCollection("orders")
    fmt.Println("Got Collection ", collection)

    filter := bson.M{"client_order_id": orderRequest.ClientOrderID}
    update := bson.M{"$set": bson.M{"status": "Cancelled"}}

    _, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        fmt.Println("Error..", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order cancelled"})
}