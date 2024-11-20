package generate

import (
	"L0/internal/models"
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/rand"
	"strconv"
	"time"
)


func GetOrder() *models.OrderJSON {
	orderId := hash32()
	var orderCount = 1 + rand.Intn(2)
	items := make([]models.Item, orderCount)
	for i := 0; i < orderCount; i++ {
		items[i] = models.Item{
			ChrtId:      0,
			TrackNumber: "",
			Price:       0,
			Rid:         "",
			Name:        "",
			Sale:        0,
			Size:        "",
			TotalPrice:  0,
			NmId:        0,
			Brand:       "",
			Status:      0,
		}
	}
	order := &models.OrderJSON{
		OrderUid:          orderId[:len(orderId)-15],
		TrackNumber:       "SOMETRACK",
		Entry:             "SOMEIL",
		Delivery:          models.Delivery{Name: "test", Phone: "+79999999999", Zip: "0", City: "moscow", Adress: "adress", Region: "msk", Email: "email"},
		Payments:          models.Payment{Transaction: "", RequestId: "", Currency: "", Provider: "", Amount: 0, PaymentDt: 0, Bank: "", DeliveryCost: 0, GoodsTotal: 0, CustomFee: 0},
		Items:             items,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "some service",
		Shardkey:          "9",
		SmId:              0,
		DateCreated:       time.Now().Format(time.RFC3339),
		OOFShard:          "0",
	}
	return order
}


func hash32() string {
	sum := md5.Sum([]byte(strconv.Itoa(rand.Intn(150000))))
	return hex.EncodeToString(sum[:])
}
