package services

import (
	"rma_payment_service/config"
	"rma_payment_service/helpers"
	"rma_payment_service/model"
	"rma_payment_service/repositories"

	"fmt"
	"net/url"
	"time"
	"github.com/google/uuid"
)

type RmaPaymentService struct {
	Repository repositories.Repository // Inject Repository
}

func NewRmaPaymentService(repository repositories.Repository) *RmaPaymentService {
	return &RmaPaymentService{Repository: repository}
}

func (ms *RmaPaymentService) MakeArRequest(amount float64, product, remitterEmail string) (*model.ArResponseData, error) {
	err := helpers.ValidateArRequest(amount, product, remitterEmail)
	if err != nil {
		return nil, err
	}
	orderNo, err := helpers.GenerateOrderNo(14)
	if err != nil {
		return nil, err
	}
	loc, _ := time.LoadLocation("Asia/Thimphu")
	date := time.Now().In(loc)
	txnTime := date.Format("20060102150405")

	payloadData := map[string]string{
		"bfs_msgType":       "AR",
		"bfs_benfTxnTime":   txnTime,
		"bfs_orderNo":       orderNo,
		"bfs_benfId":        config.BENF_ID,        // Replace with actual config value
		"bfs_benfBankCode":  config.BENF_BANK_CODE, // Replace with actual config value
		"bfs_txnCurrency":   config.CURRENCY,       // Replace with actual config value
		"bfs_txnAmount":     fmt.Sprintf("%.2f", amount),
		"bfs_remitterEmail": remitterEmail,
		"bfs_paymentDesc":   "Educareskill-" + product,
		"bfs_version":       "1.0",
		"bfs_checkSum":      "",
	}

	checksumDataString := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		payloadData["bfs_benfBankCode"],
		payloadData["bfs_benfId"],
		payloadData["bfs_benfTxnTime"],
		payloadData["bfs_msgType"],
		payloadData["bfs_orderNo"],
		payloadData["bfs_paymentDesc"],
		payloadData["bfs_remitterEmail"],
		payloadData["bfs_txnAmount"],
		payloadData["bfs_txnCurrency"],
		payloadData["bfs_version"],
	)

	// Load private key from file
	checksum, err := helpers.GenerateCheckSum(checksumDataString)
	if err != nil {
		return nil, err
	}
	payloadData["bfs_checkSum"] = checksum
	// Build query string from payloadData
	bodyBytes, err := helpers.RmaPostRequest(payloadData)
	if err != nil {
		return nil, err
	}
	bankList, status, txnId, err := helpers.FormatArResponse(string(bodyBytes))

	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	arResponse := model.ArResponseData{
		TxnId:  txnId,
		Status: status,
		Banks:  bankList,
	}

	return &arResponse, nil
}
func (ms *RmaPaymentService) MakeAeRequest(remitterAccNo, remitterBankId, txnId string) (string, error) {
	payloadData := map[string]string{
		"bfs_msgType":        "AE",
		"bfs_bfsTxnId":       txnId,
		"bfs_benfId":         config.BENF_ID,
		"bfs_remitterBankId": remitterBankId, // Replace with actual config value
		"bfs_remitterAccNo":  remitterAccNo,  // Replace with actual config value
		"bfs_checkSum":       "",
	}

	checksumDataString := fmt.Sprintf("%s|%s|%s|%s|%s",
		payloadData["bfs_benfId"],
		payloadData["bfs_bfsTxnId"],
		payloadData["bfs_msgType"],
		payloadData["bfs_remitterAccNo"],
		payloadData["bfs_remitterBankId"],
	)

	// Load private key from file
	checksum, err := helpers.GenerateCheckSum(checksumDataString)
	if err != nil {
		return "", err
	}
	payloadData["bfs_checkSum"] = checksum
	bodyBytes, err := helpers.RmaPostRequest(payloadData)
	if err != nil {
		return "", err
	}

	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("error parsing URL-encoded string: %v", err)
	}
	bfsRespondCode := values.Get("bfs_responseCode")
	bfsResponseDesc := values.Get("bfs_responseDesc")

	if bfsRespondCode != "00" || bfsResponseDesc != "Success" {
		return "", fmt.Errorf("error occured while for the payment")
	}

	return txnId, nil
}

func (ms *RmaPaymentService) MakeDrRequest(txnId, remitterOtp string, userId, membershipDurationId uuid.UUID) (string, error) {
	err := helpers.ValidateDrRequest(txnId, remitterOtp, userId, membershipDurationId)
	if err != nil {
		return "", err
	}

	membership, err := helpers.GetMembership(userId)
    if err != nil {
        return "", fmt.Errorf("failed to fetch user membership: %v", err)
    }

	payloadData := map[string]string{
		"bfs_msgType":     "DR",
		"bfs_bfsTxnId":    txnId,
		"bfs_benfId":      config.BENF_ID,
		"bfs_remitterOtp": remitterOtp, // Replace with actual config value
		"bfs_checkSum":    "",
	}

	checksumDataString := fmt.Sprintf("%s|%s|%s|%s",
		payloadData["bfs_benfId"],
		payloadData["bfs_bfsTxnId"],
		payloadData["bfs_msgType"],
		payloadData["bfs_remitterOtp"],
	)
	// Load private key from file
	checksum, err := helpers.GenerateCheckSum(checksumDataString)
	if err != nil {
		return "", err
	}
	payloadData["bfs_checkSum"] = checksum
	// Build query string from payloadData
	bodyBytes, err := helpers.RmaPostRequest(payloadData)
	if err != nil {
		return "", err
	}
	fmt.Println("Response from RMA:", string(bodyBytes))
	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("error parsing URL-encoded string: %v", err)
	}
	remarks := helpers.GetResponseDescription(values.Get("bfs_debitAuthCode"))

	paymentData := model.RmaPaymentResponseInput{
		BfsBfsTxnId:       values.Get("bfs_bfsTxnId"),
		BfsDebitAuthNo:    values.Get("bfs_debitAuthNo"),
		BfsRemitterName:   values.Get("bfs_remitterName"),
		BfsTxnCurrency:    values.Get("bfs_txnCurrency"),
		BfsBfsTxnTime:     values.Get("bfs_bfsTxnTime"),
		BfsBenfId:         values.Get("bfs_benfId"),
		BfsRemitterBankId: values.Get("bfs_remitterBankId"),
		BfsOrderNo:        values.Get("bfs_orderNo"),
		BfsDebitAuthCode:  values.Get("bfs_debitAuthCode"),
		BfsTxnAmount:      values.Get("bfs_txnAmount"),
		BfsBenfTxnTime:    values.Get("bfs_benfTxnTime"),
		BfsMsgType:        values.Get("bfs_msgType"),
		UserId:            userId,
		MembershipDurationId: membership.MembershipDurationId,
		Remarks:           remarks,
	}

	err = ms.StorePaymentResponse(paymentData)
	if err != nil {
		return remarks, fmt.Errorf("failed to find the user: %v", err)
	}
	return remarks, nil
}

func (ms *RmaPaymentService) StorePaymentResponse(paymentInput model.RmaPaymentResponseInput) error {
	return ms.Repository.StorePaymentResponse(paymentInput)
}
