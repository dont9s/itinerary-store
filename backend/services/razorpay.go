package services

import (
    "os"
    "github.com/razorpay/razorpay-go"
)

var Client *razorpay.Client

func InitRazorpay() {
    Client = razorpay.NewClient(
        os.Getenv("RAZORPAY_KEY_ID"),
        os.Getenv("RAZORPAY_KEY_SECRET"),
    )
}
