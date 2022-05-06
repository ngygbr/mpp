package db

import (
	"encoding/json"

	"github.com/ngygbr/mpp/pkg/model"

	"github.com/dgraph-io/badger/v3"
)

func Create(transaction *model.Transaction) error {
	transactionInByte, _ := json.Marshal(transaction)

	if err := Database.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(transaction.ID), transactionInByte)
		return err
	}); err != nil {
		return err
	}

	return nil
}

func Update(transaction *model.Transaction) error {
	err := Delete(transaction.ID)
	if err != nil {
		return err
	}

	err = Create(transaction)
	if err != nil {
		return err
	}

	return nil
}

func GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction

	if err := Database.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			transaction, err := GetByID(string(k))
			if err != nil {
				return err
			}
			transactions = append(transactions, transaction)
		}
		return nil
	}); err != nil {
		return []model.Transaction(nil), err
	}

	return transactions, nil
}

func GetByID(id string) (model.Transaction, error) {
	var transaction model.Transaction

	if err := Database.View(func(txn *badger.Txn) error {
		i, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}
		return i.Value(func(val []byte) error {
			return json.Unmarshal(val, &transaction)
		})

	}); err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func Delete(id string) error{
	if err := Database.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func DeleteAll() error {
	if err := Database.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := Delete(string(k))
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
