package config

import (
	"os"
)

func PostgresUser() string     { return os.Getenv("POSTGRES_USER") }
func PostgresPassword() string { return os.Getenv("POSTGRES_PASSWORD") }
func PostgresHost() string     { return os.Getenv("POSTGRES_HOST") }
func PostgresPort() string     { return os.Getenv("POSTGRES_PORT") }
func PostgresDB() string       { return os.Getenv("POSTGRES_DB") }

const (
    Database                  = "educareRmaPayment"
	PaymentResponseCollection = "payment_responses"
	BENF_ID                   = "BE10000169"
	BENF_BANK_CODE            = "01"
	CURRENCY                  = "BTN"
	PAYMENT_PRIVATE_KEY       = "keys/rma-payment.key"
	PAYMENT_PUBLIC_KEY        = "keys/public.key"
	RMA_PAYMENT_API           = "https://bfssecure.rma.org.bt/BFSSecure/nvpapi"
)


