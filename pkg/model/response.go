package model

type ActionResponse struct {
	Status       string         `json:"status"`
	StatusCode   uint64         `json:"status_code"`
	Message      string         `json:"message"`
	Transaction  *Transaction   `json:"transaction,omitempty"`
	Transactions *[]Transaction `json:"transactions,omitempty"`
}

type TransactionResponse struct {
	Status     string       `json:"status"`
	StatusCode uint64       `json:"status_code"`
	Message    string       `json:"message"`
	Data       *Transaction `json:"data,omitempty"`
}

func (resp *TransactionResponse) SetTransactionResponse(transaction *Transaction, code uint64, err error) {
	resp.StatusCode = code

	if err != nil {
		resp.Status = "failed"
		resp.Message = err.Error()
	} else {
		resp.Status = "success"
		resp.Message = "success"
	}

	resp.Data = transaction
}

func (resp *ActionResponse) SetActionResponse(code uint64, err error, transaction *Transaction, transactions *[]Transaction) {
	resp.StatusCode = code

	if err != nil {
		resp.Status = "failed"
		resp.Message = err.Error()
	} else {
		resp.Status = "success"
		resp.Message = "success"
	}

	resp.Transaction = transaction
	resp.Transactions = transactions
}
