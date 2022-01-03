package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	// print string
	print("Hello World\n")
}

func splitSkuName(name string) string {
	var skuName string
	splitSkuName := strings.Split(name, "_")
	if len(splitSkuName) >= 3 {
		skuName = fmt.Sprintf("%s %s", splitSkuName[1], splitSkuName[2])
	} else {
		skuName = fmt.Sprintf(splitSkuName[1])
	}

	return strings.TrimSpace(skuName)
}
