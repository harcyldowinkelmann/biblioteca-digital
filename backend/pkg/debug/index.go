package debug

import (
	"encoding/json"
	"fmt"
	"log"
)

func Json[T any](data T) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Erro ao converter para JSON: %v", err)
		return
	}

	fmt.Println(string(jsonData))
}
