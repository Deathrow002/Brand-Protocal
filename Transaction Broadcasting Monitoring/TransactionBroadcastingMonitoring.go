package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	broadcastURL = "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	checkURL     = "https://mock-node-wgqbnxruha-as.a.run.app/check/"
)

func TransactionBroadcastingMonitoring() {
	var symbol_input string
	var price_input int

	current_time_seconds := int(time.Now().Unix())

	fmt.Print("Input Symbol: ")
	fmt.Scanln(&symbol_input)

	fmt.Print("Input Price: ")
	fmt.Scanln(&price_input)

	Data := map[string]interface{}{
		"symbol":    symbol_input,
		"price":     price_input,
		"timestamp": current_time_seconds,
	}

	payload, err := json.Marshal(Data)

	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	resp, err := http.Post(broadcastURL, "application/json", bytes.NewReader(payload))

	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		fmt.Println("Error sending data:", string(body))
		return
	}

	fmt.Println("Data sent successfully!")

	var responseData map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	txHash, ok := responseData["tx_hash"].(string)
	if !ok {
		fmt.Println("Error: missing tx_hash in response")
		return
	}

	fmt.Println("Transaction Hash:", txHash)

	for {
		checkURL := checkURL + txHash

		resp, err := http.Get(checkURL)
		if err != nil {
			fmt.Println("Error fetching transaction status:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error fetching transaction status:", resp.Status)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var statusData map[string]interface{}
		err = json.Unmarshal(body, &statusData)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		txStatus, ok := statusData["tx_status"].(string)
		if !ok {
			fmt.Println("Error: missing tx_status in response")
			return
		}

		fmt.Println("Transaction Status:", txStatus)

		if txStatus == "CONFIRMED" || txStatus == "FAILED" || txStatus == "DNE" {
			break
		}

		time.Sleep(3 * time.Second)
	}

}

func main() {
	TransactionBroadcastingMonitoring()
}
