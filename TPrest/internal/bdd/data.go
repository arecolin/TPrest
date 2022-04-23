package bdd

import (
	"fmt"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func DbOpen(fileName string){
	db, _ = bolt.Open(fileName, 0600, nil)

}

func DbPath() string {
	return db.Path()	
}

func DbCheckIfOpen() bool {
	return db.IsOpen()
}

func DbClose() {
	db.Close()
}

func DbPut(key string, value string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mydata"))
		err := b.Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}
		return nil
	})
}

func DbGet(key int) JSON {
	var value JSON
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mydata"))
		v := b.Get([]byte(key))
		value = JSON(v)
		return nil
	})
	return value
}

func DbDelete(key int) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mydata"))
		err := b.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("delete key: %s", err)
		}
		return nil
	})
}


func DbGetAll() map[int]JSON {
	var data = make(map[int]JSON)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mydata"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			data[int(k)] = JSON(v)
		}
		return nil
	})
	return data
}

