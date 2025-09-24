package helpers

func GetResponseDescription(responseCode string) string {
    switch responseCode {
    case "00":
        return "Approved"
    case "03":
        return "Invalid Beneficiary"
    case "05":
        return "Beneficiary Account Closed"
    case "12":
        return "Invalid Transaction"
    case "13":
        return "Invalid Amount"
    case "14":
        return "Invalid Remitter Account"
    case "20":
        return "Invalid Response"
    case "30":
        return "Transaction Not Supported Or Format Error"
    case "45":
        return "Duplicate Beneficiary Order Number"
    case "47":
        return "Invalid Currency"
    case "48":
        return "Transaction Limit Exceeded"
    case "51":
        return "Insufficient Funds"
    case "53":
        return "No Savings Account"
    case "57":
        return "Transaction Not Permitted"
    case "61":
        return "Withdrawal Limit Exceeded"
    case "65":
        return "Withdrawal Frequency Exceeded"
    case "76":
        return "Transaction Not Found"
    case "78":
        return "Decryption Failed"
    case "80":
        return "Buyer Cancel Transaction"
    case "84":
        return "Invalid Transaction Type"
    case "85":
        return "Internal Error At Bank System"
    case "BC":
        return "Transaction Cancelled By Customer"
    case "FE":
        return "Internal Error"
    case "OA":
        return "Session Timeout at BFS Secure Entry Page"
    case "OE":
        return "Transaction Rejected As Not In Operating Hours"
    case "OF":
        return "Transaction Timeout"
    case "SB":
        return "Invalid Beneficiary Bank Code"
    case "XE":
        return "Invalid Message"
    case "XT":
        return "Invalid Transaction Type"
    default:
        return "Unknown Error"
    }
}