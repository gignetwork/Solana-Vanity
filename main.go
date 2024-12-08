package main

import (
	"fmt"
	"strings"
	"time"
        "os"


	"github.com/gagliardetto/solana-go"
)

var generatedCount = 0
var numThreads = 10
var startTime = time.Now()
var searchTerm = ""
var shouldStopThreads = false

func generateWallet() {
	for {
		if shouldStopThreads {
			return
		}
		newWallet := solana.NewWallet()
		if strings.HasPrefix(newWallet.PublicKey().String(), searchTerm) && !shouldStopThreads {
			firstCharAfterSearchTerm := strings.Split(newWallet.PublicKey().String(), searchTerm)[1][0:1]
			if firstCharAfterSearchTerm == strings.ToUpper(firstCharAfterSearchTerm) {
                                fmt.Printf("%s", newWallet.PublicKey())
				fmt.Printf(":%v\n", newWallet.PrivateKey)
			
				shouldStopThreads = true
				
			}
		}
		generatedCount++
		if generatedCount%1000000 == 0 {
			fmt.Printf("Status: %d wallets generated in %s\n", generatedCount, time.Since(startTime))
		}
	}
}

func main() {
	

		go generateWallet()
		time.Sleep(2*time.Second)
                os.Exit(0)

	fmt.Scanln()

}
