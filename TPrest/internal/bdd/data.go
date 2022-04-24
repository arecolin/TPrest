package bdd

import (
	"encoding/json"
	"fmt"
	. "internal/entities"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func DbOpen(fileName string) {
	db, _ = bolt.Open(fileName, 0600, nil)

}

func DbPath() string {
	return db.Path()
}

func CreateBucket(bucketName string) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		fmt.Println("Bucket created:", bucketName)
		return nil
	})
}

func DbClose() {
	db.Close()
	fmt.Println("Database closed")
}

func SaveStudent(student Student) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		encoded, err := json.Marshal(student)
		err = b.Put([]byte(student.Id), encoded)
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}
		return nil
	})
	fmt.Println("Student saved:", student)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		v := b.Get([]byte(student.Id))
		fmt.Printf("The student is: %s\n", string(v))
		return nil
	})
}

func SaveLanguage(language Language) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		encoded, err := json.Marshal(language)
		err = b.Put([]byte(language.Code), encoded)
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}
		return nil
	})
	fmt.Println("Language saved:", language)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		v := b.Get([]byte(language.Code))
		fmt.Printf("The language is: %s\n", string(v))
		return nil
	})
}

func DbGetStudent(key string) string {
	var value string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		v := b.Get([]byte(key))
		value = string(v)
		return nil
	})
	return value
}

func DbGetLanguage(key string) string {
	var value string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		v := b.Get([]byte(key))
		value = string(v)
		return nil
	})
	return value
}

func DbDeleteStudent(key string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		err := b.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("delete key: %s", err)
		}
		return nil
	})
}

func DbDeleteLanguage(key string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		err := b.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("delete key: %s", err)
		}
		return nil
	})
}

func DbGetAll(bucketName string) map[string]string {
	var data = make(map[string]string)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			data[string(k)] = string(v)
		}
		return nil
	})
	return data
}

func DbUpdateStudent(student Student) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		encoded, err := json.Marshal(student)
		err = b.Put([]byte(student.Id), encoded)
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}
		return nil
	})
	fmt.Println("Student updated:", student)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("students"))
		v := b.Get([]byte(student.Id))
		fmt.Printf("The student is: %s\n", string(v))
		return nil
	})
}

func DbUpdateLanguage(language Language) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		encoded, err := json.Marshal(language)
		err = b.Put([]byte(language.Code), encoded)
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}
		return nil
	})
	fmt.Println("Language updated:", language)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("languages"))
		v := b.Get([]byte(language.Code))
		fmt.Printf("The language is: %s\n", string(v))
		return nil
	})
}
