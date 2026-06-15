package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Latihan 1: Sistem Validasi Kupon Belanja dengan Map & Sentinel Error
*/

var (
	ErrCouponNotFound = errors.New("Coupon do not match")
	ErrRateInvalid    = errors.New("Discount rate invalid")
)

type DiscountRate float64

type DiscountManager struct {
	Items map[string]DiscountRate
}

func (manager *DiscountManager) RegisterCoupon(
	code string,
	rate DiscountRate,
) error {
	if rate < 0.0 || rate > 1.0 {
		return ErrRateInvalid
	}
	codeUpperCase := strings.ToUpper(code)
	manager.Items[codeUpperCase] = rate
	return nil
}

func (manager *DiscountManager) CalculateDiscount(
	code string,
	originalPrice float64,
) (float64, error) {
	discountRate, exist := manager.Items[code]

	if !exist {
		return 0.0, ErrCouponNotFound
	}
	price := originalPrice * (1.0 - float64(discountRate))
	return price, nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("Sistem Validasi Kupon Belanja dengan Map & Sentinel Error")
	fmt.Println("==================================================")

	manager := DiscountManager{Items: map[string]DiscountRate{}}
	// codes := []string{"SAVE20", "MEGA80", "NIL150"}
	codes := make(map[string]DiscountRate)
	codes["SAVE20"] = 0.2
	codes["MEGA80"] = 0.8
	codes["NIL150"] = 1.5

	prices := []float64{1_000_000.0, 5_000_000}

	for code, discountRate := range codes {

		if err := manager.RegisterCoupon(code, discountRate); err != nil {
			errWrap := fmt.Errorf("Code [%s] DiscountRate [%.2f] %w", code, discountRate, err)
			fmt.Println(errWrap)
		}
	}
	fmt.Println(prices)
	fmt.Println(codes)
	fmt.Println(manager.Items)

	for i, originalPrice := range prices {
		fmt.Printf("%d. Original Price [%.2f]: \n", i+1, originalPrice)
		for code, discountRate := range codes {
			fmt.Printf("  Coupon Code [%s]", code)
			if price, err := manager.CalculateDiscount(code, originalPrice); err == nil {
				fmt.Printf("    Dicounted Price: Rp %.2f\n", price)
			} else {
				errWrap := fmt.Errorf(" Discount Rate [%.2f] %w", discountRate, err)
				fmt.Println(errWrap)
			}
		}

	}
}
