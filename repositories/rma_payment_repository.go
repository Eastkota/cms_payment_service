package repositories

import (
	"fmt"
	"rma_payment_service/model"
	"gorm.io/gorm"
)

type RmaPaymentRepository struct{
	DB *gorm.DB
}

func NewRmaPaymentRepository(db *gorm.DB) *RmaPaymentRepository {
	return &RmaPaymentRepository{DB: db}
}

func (repo *RmaPaymentRepository) StorePaymentResponse(paymentInput model.RmaPaymentResponseInput) error {

    result := repo.DB.Create(&model.PaymentResponse{
		BfsBfsTxnId:       paymentInput.BfsBfsTxnId,
		BfsDebitAuthNo:    paymentInput.BfsDebitAuthNo,
		BfsRemitterName:   paymentInput.BfsRemitterName,
		BfsTxnCurrency:    paymentInput.BfsTxnCurrency,
		BfsBfsTxnTime:     paymentInput.BfsBfsTxnTime,
		BfsBenfId:         paymentInput.BfsBenfId,
		BfsRemitterBankId: paymentInput.BfsRemitterBankId,
		BfsOrderNo:        paymentInput.BfsOrderNo,
		BfsDebitAuthCode:  paymentInput.BfsDebitAuthCode,
		BfsTxnAmount:      paymentInput.BfsTxnAmount,
		BfsBenfTxnTime:    paymentInput.BfsBenfTxnTime,
		BfsMsgType:        paymentInput.BfsMsgType,
		UserId:            paymentInput.UserId,
	})

    if result.Error != nil {
        return fmt.Errorf("failed to save the payment responses: %v", result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("no rows were inserted")
    }
    
    return nil
}
