package utils

import (
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/kristabdi/bnmo-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func (Db *DbInstance) InitSeeding() error {
	var err error
	faker.SetGenerateUniqueValues(true)

	if !Db.Migrator().HasTable(&models.User{}) {
		err = Db.Migrator().CreateTable(&models.User{})
		if err != nil {
			return err
		}

		for i := 0; i < 50; i++ {
			newUser := models.User{}
			err = faker.FakeData(&newUser)
			var hashed []byte
			hashed, err = bcrypt.GenerateFromPassword([]byte(faker.Sentence()), 14)
			if err != nil {
				return err
			}

			newUser.Password = string(hashed)
			newUser.Balance = uint64((rand.Intn(100000-1000) + 1000) * 1000)
			result := Db.Create(&newUser)
			if result.Error != nil {
				return err
			}

		}
	}

	log.Println("User table done!")

	if !Db.Migrator().HasTable(&models.Request{}) {
		err := Db.Migrator().CreateTable(&models.Request{})
		if err != nil {
			return err
		}
		for i := 1; i <= 50; i++ {
			amount := (rand.Intn(500-100) + 100) * 1000
			rand.Seed(time.Now().UnixNano())
			newRequest := models.Request{
				UserID:     int64(i),
				Amount:     uint64(amount),
				IsAdd:      rand.Intn(2) == 1,
				IsApproved: true,
			}
			result := Db.Create(&newRequest)
			if result.Error != nil {
				return err
			}
		}
	}

	log.Println("Request table done!")

	if !Db.Migrator().HasTable(&models.Transaction{}) {
		err := Db.Migrator().CreateTable(&models.Transaction{})
		if err != nil {
			return err
		}

		var idAvailable []uint
		for i := 1; i <= 50; i++ {
			idAvailable = append(idAvailable, uint(i))
		}
		rand.Shuffle(len(idAvailable), func(i, j int) { idAvailable[i], idAvailable[j] = idAvailable[j], idAvailable[i] })

		for i := 0; i < 20; i++ {
			amount := (rand.Intn(100-20) + 20) * 1000
			newTransaction := models.Transaction{
				IdFrom:   int64(idAvailable[0]),
				IdTo:     int64(idAvailable[1]),
				Amount:   uint64(amount),
				UserFrom: models.User{},
				UserTo:   models.User{},
			}

			result := Db.Create(&newTransaction)
			if result.Error != nil {
				return err
			}

			idAvailable = idAvailable[2:]
		}
	}

	log.Println("Transaction table done!")

	faker.SetGenerateUniqueValues(false)
	return nil
}
