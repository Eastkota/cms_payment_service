package model

import (
	"time"
	
	"github.com/google/uuid"
)

type PaymentResponse struct {
	BfsBfsTxnId       string `json:"bfs_bfsTxnId" gorm:"type:string;"`
	BfsDebitAuthNo    string `json:"bfs_debitAuthNo" gorm:"type:varchar(15)"`
	BfsRemitterName   string `json:"bfs_remitterName" gorm:"type:varchar(50)"`
	BfsTxnCurrency    string `json:"bfs_txnCurrency" gorm:"type:varchar(50)"`
	BfsBfsTxnTime     string `json:"bfs_bfsTxnTime" gorm:"type:varchar(50)"`
	BfsBenfId         string `json:"bfs_benfId" gorm:"type:string"`
	BfsRemitterBankId string `json:"bfs_remitterBankId" gorm:"type:string"`
	BfsOrderNo        string `json:"bfs_orderNo" gorm:"type:string"`
	BfsDebitAuthCode  string `json:"bfs_debitAuthCode" gorm:"type:string"`
	BfsTxnAmount      string `json:"bfs_txnAmount" gorm:"type:string"`
	BfsBenfTxnTime    string `json:"bfs_benfTxnTime" gorm:"type:varchar(50)"`
	BfsMsgType        string `json:"bfs_msgType" gorm:"type:varchar(50)"`
	MembershipDurationId uuid.UUID    `json:"membership_duration_id" gorm:"type:uuid"`
	UserId            uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

func (PaymentResponse) TableName() string {
    return "payment.payment_responses_rma"
}

type StripeResponse struct {
	Amount                string    `gorm:"type:varchar" json:"amount"`
	TransactionId         string    `gorm:"type:varchar" json:"transaction_id"`
	UserId                uuid.UUID `gorm:"type:uuid" json:"user_id"`
	MembershipDurationId  uuid.UUID `gorm:"type:uuid" json:"membership_duration_id"`
	Status                string    `gorm:"type:varchar" json:"status"`
	Remarks               string    `gorm:"type:varchar" json:"remarks"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Product               string    `gorm:"type:varchar" json:"product"`
}

func (StripeResponse) TableName() string {
    return "payment.payment_responses_stripe"
}


type GenericCmsPaymentResponse struct {
	Data  string `json:"data"`
	Error *CmsPaymentError
}
