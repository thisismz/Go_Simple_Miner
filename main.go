package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type SimpleTransaction struct {
	ID     int
	from   string
	to     string
	amount int
}
type SimpleBlock struct {
	ID                int
	previousBlockHash string
	transactions      []SimpleTransaction
}

func blockHash(o interface{}, nonce string) []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v%s", o, nonce)))
	return h.Sum(nil)
}
func mine(o interface{}, difficulty int) {
	maxnonce := 100000000000
	perfixZeros := strings.Repeat("0", difficulty)
	var nonce int
	for nonce <= maxnonce {
		hashResult := fmt.Sprintf("%x", blockHash(o, strconv.Itoa(nonce)))
		nonce++
		if strings.HasPrefix(hashResult, perfixZeros) {
			println("Nouce is : ", nonce)
			println("The block Hash  is : ", hashResult)
			break
		}
	}

}
func main() {
	currentTime := time.Now()
	difficulty := 5 // increase number for make mining harder
	transone := SimpleTransaction{
		ID:     1,
		from:   "mahdi",
		to:     "mohammad",
		amount: 25,
	}
	tanstwo := SimpleTransaction{
		ID:     2,
		from:   "roya",
		to:     "negin",
		amount: 99,
	}
	block := SimpleBlock{
		ID:                1,
		previousBlockHash: "00000891861fd09d3cf161db54434c9e518cf80e94811e7762c3aee8a7af39af",
		transactions:      []SimpleTransaction{transone, tanstwo},
	}
	log.Printf("Start Mining")
	mine(block, difficulty)
	log.Printf("End of Mining")
	duration := time.Since(currentTime)
	log.Printf("Mining took %s", duration)

}
