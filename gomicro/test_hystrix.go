package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func GetProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 6 {
		time.Sleep(time.Second * 3)
	}
	return Product{
		ID:    1,
		Title: "radial",
		Price: 13,
	}, nil
}

func RecProduct() (Product, error) {
	return Product{
		ID:    100,
		Title: "radial_hks",
		Price: 13000,
	}, nil
}

// v1.0
func main_old() {
	rand.Seed(time.Now().UnixNano())
	for {
		p, _ := GetProduct()
		fmt.Println(p)
		time.Sleep(time.Second * 1)
	}
}

// v2.0
func main_() {
	rand.Seed(time.Now().UnixNano())
	for {
		err := hystrix.Do("get_product", func() error {
			p, _ := GetProduct()
			fmt.Println(p)
			return nil
		}, nil)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}

// v2.0
func main_2() {
	rand.Seed(time.Now().UnixNano())
	// config
	confidA := hystrix.CommandConfig{
		Timeout: 2000,
	}
	hystrix.ConfigureCommand("get_product", confidA)

	for {
		err := hystrix.Do("get_product",
			func() error {
				p, _ := GetProduct()
				fmt.Println(p)
				return nil
			},
			nil)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}

// v3.0
func main_fallback() {
	rand.Seed(time.Now().UnixNano())
	// config
	confidA := hystrix.CommandConfig{
		Timeout: 2000,
	}
	hystrix.ConfigureCommand("get_product", confidA)

	for {
		err := hystrix.Do("get_product",
			func() error {
				p, _ := GetProduct()
				fmt.Println(p)
				return nil
			}, func(e error) error {
				p, _ := RecProduct()
				fmt.Println(p)
				//return nil
				return errors.New("time out")
			})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}

// v4.0
func main() {
	rand.Seed(time.Now().UnixNano())
	// config
	confidA := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  5,
		RequestVolumeThreshold: 3,
		ErrorPercentThreshold:  20,
		SleepWindow:            int(time.Second * 15),
	}
	hystrix.ConfigureCommand("get_product", confidA)
	resChan := make(chan Product, 1)
	c, _, _ := hystrix.GetCircuit("get_product")
	for {
		errs := hystrix.Go("get_product",
			func() error {
				p, _ := GetProduct()
				resChan <- p
				return nil
			}, func(e error) error {
				p, _ := RecProduct()
				resChan <- p
				//return nil
				return errors.New("time out")
			})
		//if err != nil {
		//	fmt.Println(err)
		//}
		select {
		case product := <-resChan:
			fmt.Println(product)
		case err := <-errs:
			fmt.Println(err)
		}
		fmt.Println(c.IsOpen())
		time.Sleep(time.Second * 1)
	}
}
