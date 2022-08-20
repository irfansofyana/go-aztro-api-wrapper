# Aztro API Wrapper in Go

[Aztro API](https://github.com/sameerkumar18/aztro) is API get daily horoscope. This project is aim to create a simple wrapper for the API using Go.

## Installation

Open your terminal and run:
```go
go get github.com/irfansofyana/go-aztro-api-wrapper
```


## Example Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/irfansofyana/go-aztro-api-wrapper/aztro"
)

func main() {
	aztroClient, err := aztro.NewAztroClient()
	if err != nil {
		log.Fatal(err)
	}

	aztroParam := aztro.NewAztroRequestParam(aztro.Taurus)
	todayHoroscope, aztroErr := aztroClient.GetHoroscope(aztroParam)
	if aztroErr != nil {
		log.Fatal(aztroErr)
	}
	fmt.Println(todayHoroscope) // Get today's horoscope

	tmrrowParam := aztro.NewAztroRequestParam(
		aztro.Taurus,
		aztro.WithDay(aztro.Tomorrow),
	)
	tmrrwHoroscope, aztroErr := aztroClient.GetHoroscope(tmrrowParam)
	if aztroErr != nil {
		log.Fatal(aztroErr)
	}
	fmt.Println(tmrrwHoroscope) // Get tomorrow's horoscope
}

```