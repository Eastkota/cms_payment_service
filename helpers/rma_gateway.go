package helpers

import (
	"rma_payment_service/config"
	"rma_payment_service/model"
	
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func GenerateOrderNo(length int) (string, error) {
	pool := "0123456789"
	poolLength := int64(len(pool))
	uniqueID := ""

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(poolLength))
		if err != nil {
			return "", err
		}
		uniqueID += string(pool[n.Int64()])
	}

	return uniqueID, nil
}

func FormatArResponse(urlEncodedString string) ([]model.Bank, string, string, error) {
	values, err := url.ParseQuery(urlEncodedString)
	if err != nil {
		return nil, "Invalid Response", "NA", fmt.Errorf("error parsing URL-encoded string: %v", err)
	}

	// Extract Data
	txnId := values.Get("bfs_bfsTxnId")
	status := values.Get("bfs_responseDesc")
	bankList := values.Get("bfs_bankList")
	if bankList == "" || status != "Success" {
		return nil, "Invalid Response", "NA", fmt.Errorf("BankList not found")
	}

	// Decode the bankList
	bankList = strings.ReplaceAll(bankList, "%7E", "#")
	bankEntries := strings.Split(bankList, "#")

	// Create a map to hold bank codes and names
	var banks []model.Bank
	// Loop through the entries and split code and name
	for _, entry := range bankEntries {
		parts := strings.Split(entry, "~") // Correct split by "~"
		if len(parts) >= 2 {               // Ensure there are at least two parts (code and name)
			code := parts[0]
			name := strings.ReplaceAll(parts[1], "+", " ") // Replace "+" with space
			banks = append(banks, model.Bank{Code: code, Name: name})
		}
	}
	return banks, status, txnId, nil
}

func GenerateCheckSum(checksumDataString string) (string, error) {
	privateKeyData, err := os.ReadFile(config.PAYMENT_PRIVATE_KEY)
	if err != nil {
		return "", fmt.Errorf("failed to read private key: %v", err)
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return "", fmt.Errorf("invalid private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Hash the checksum data string
	hashed := sha1.Sum([]byte(checksumDataString))

	// Sign the checksum data string
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	// Convert signature to hexadecimal string
	hexSignature := hex.EncodeToString(signature)

	// Load public key from file
	publicKeyData, err := os.ReadFile(config.PAYMENT_PUBLIC_KEY)
	if err != nil {
		return "", fmt.Errorf("failed to read public key: %v", err)
	}

	publicBlock, _ := pem.Decode(publicKeyData)
	if publicBlock == nil {
		return "", fmt.Errorf("invalid public key")
	}
	cert, err := x509.ParseCertificate(publicBlock.Bytes)
	if err != nil {
		return "", fmt.Errorf("error parsing certificate: %v", err)
	}

	// Extract the public key (usually *rsa.PublicKey)
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error asserting public key type: expected RSA")
	}

	pbuSignature, err := hex.DecodeString(hexSignature)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %v", err)
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hashed[:], pbuSignature)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %v", err)
	}
	return strings.ToUpper(hexSignature), nil
}

func RmaPostRequest(payloadData map[string]string) ([]byte, error) {
	queryString := url.Values{}
	for key, value := range payloadData {
		queryString.Set(key, value)
	}

	// Construct the full URL
	// fullURL := config.RMA_PAYMENT_API
	fullURL := config.RMA_PAYMENT_API + "?" + queryString.Encode()
	// fmt.Println("Full Request URL:", fullURL) // Debugging output

	// Create a client with increased timeout
	client := &http.Client{
		Timeout: 30 * time.Second, // Adjust timeout as necessary
	}

	// Create the POST request
	req, err := http.NewRequest("POST", fullURL, nil) // Consider using no body if params are in the URL
	// req, err := http.NewRequest("POST", fullURL, bytes.NewBufferString(queryString.Encode())) // Consider using no body if params are in the URL
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Printf("Request timed out: %v\n", err)
		} else {
			fmt.Printf("Request failed: %v\n", err)
		}
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return bodyBytes, nil
}

// func RmaPostRequest(payloadData map[string]string) ([]byte, error) {
// 	queryString := url.Values{}
// 	for key, value := range payloadData {
// 		queryString.Set(key, value)
// 	}

// 	fullURL := config.RMA_PAYMENT_API + "?" + queryString.Encode()
// 	// Send POST request
// 	resp, err := http.Post(fullURL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(queryString.Encode())))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to make request: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse response: %v", err)
// 	}

// 	return bodyBytes, nil
// }
