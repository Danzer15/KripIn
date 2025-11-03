package logic

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func Logic() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input text : ")
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error :", err)
	}

	plaintext = strings.TrimSpace(plaintext)

	hash := sha256.Sum256([]byte(plaintext))
	hashString := hex.EncodeToString(hash[:])

	fmt.Println("Plain Text : ", plaintext)
	fmt.Println("hash : ", hashString)

}
