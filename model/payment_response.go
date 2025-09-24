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
	UserId            uuid.UUID `json:"user_id" gorm:"type:string"`
}

type AuthUserMembership struct {
    ID                   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    MembershipJoinDate   time.Time `gorm:"type:timestamptz;not null" json:"membership_join_date"`
    MembershipEndDate    time.Time `gorm:"type:timestamptz;not null" json:"membership_end_date"`
    UserId               uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    MembershipDurationId uuid.UUID `gorm:"type:uuid;not null" json:"membership_duration_id"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}

func (PaymentResponse) TableName() string {
    return "payment.payment_responses_rma"
}

type PaymentResult struct {
	Message string `json:"message"`
}

type GenericRmaResponse struct {
	Data  string `json:"data"`
	Error *RmaPaymentError
}
type Bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ArResponseData struct {
	TxnId  string `json:"txnId"`
	Status string `json:"status"`
	Banks  []Bank `json:"banks"`
}
