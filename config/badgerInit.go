package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger"
)

const (
	DOMAINSCOLLECTION = "Domains"
)

func BadgerInit(config *Conf, dbPath string) error {
	config.BadgerDB = make(map[string]*badger.DB)
	// config.BleveIndexes = make(map[string]bleve.Index)

	badgerCollections := []string{
		DOMAINSCOLLECTION,
	}

	for _, v := range badgerCollections {
		err := config.CreateBadgerCollection(dbPath, v)
		if err != nil {
			log.Println(err)
		}
	}

	return nil

}

func (config *Conf) SetBadgerKV(collection string, key []byte, v []byte, userMeta byte) error {
	kv := config.GetBadgerCollection(collection)

	return kv.Update(func(txn *badger.Txn) error {
		err := txn.SetWithMeta(key, v, userMeta)
		return err
	})

}

func (config *Conf) DeleteBadgerKV(collection string, key []byte) error {
	kv := config.GetBadgerCollection(collection)

	return kv.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})

}

func (config *Conf) GetBadgerKV(collection string, key []byte) ([]byte, error) {
	kv := config.GetBadgerCollection(collection)

	var item *badger.Item
	var val []byte
	var err error
	err = kv.View(func(txn *badger.Txn) error {
		log.Println(string(key))
		if item, err = txn.Get(key); err != nil {
			fmt.Printf("Error while getting key: %q : %#v\n", key, err)
			return err
		}
		// item.Key()
		log.Printf("%#v", item.Key())

		val, err = item.Value()
		if err != nil {
			log.Println("errororor: ", err)
		}
		return err
	})

	return val, err
}

func (config *Conf) CreateBadgerCollection(dbPath, collection string) error {
	path := filepath.Join(dbPath, collection)
	os.MkdirAll(path, 0755)
	opt := badger.DefaultOptions
	opt.Dir = path
	opt.ValueDir = path
	var err error

	config.BadgerDB[collection], err = badger.Open(opt)
	return err
}

func (config *Conf) GetBadgerCollection(collection string) *badger.DB {
	return config.BadgerDB[collection]
}

// DisconnectDB Disconnect from the DB
func (config *Conf) DisconnectDB() {
	for _, v := range config.BadgerDB {
		if err := v.Close(); err != nil {
			log.Println("close badger.Db error: ", err)
		}
	}
}
