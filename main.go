package main

/**
This is a fork of the blockchain article found at https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc

I made some changes that I think make things easier to understand. The article was concise and (rightly) focused on the blockchain aspects. My changes are
mostly for usability, maintainability, and clearer delineation of responsibilities. For example, we made sure all returns are the same shape and that we don't
directly return arrays (OWASP JSON Array security risk).
*/
import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	go func() {
		t := time.Now()
		firstBlock := Block{
			Index:     0,
			Timestamp: t.String(),
			BPM:       0,
			PrevHash:  "",
			Hash:      "",
		}
		Blockchain = append(Blockchain, firstBlock)
	}()

	log.Fatal(run())
}
