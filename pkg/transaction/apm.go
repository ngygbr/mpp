package transaction

import "mpp/pkg/model"

var ValidTypes = []string{
	"dragonpay",
	"alipay",
	"sepa",
}

func APMTransaction(transaction model.Transaction) (model.Transaction, error){
	//TODO

	return model.Transaction{}, nil
}
