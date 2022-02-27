package main

import (
	"consumer/domains"
	"fmt"

	"github.com/Shopify/sarama"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	connection_string := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		"root",
		"example",
		"mysql",
		"3306",
		"micro",
	)

	dial := mysql.Open(connection_string)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", db.Config.Logger)

	//testDB(db)
	testKafka()
}

func testDB(db *gorm.DB) {
	userRepo := domains.NewUserRepository(db)
	users, _ := userRepo.FindAll()
	for _, v := range users {
		fmt.Printf("%#v,%#v\n", v.FirstName, v.LastName)
	}
}

func testKafka() {
	servers := []string{"kafka:9092"}
	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		fmt.Println("NewConsumer Error")
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("sample_topic", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("ConsumePartition Error")
		panic(err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg:%+v {%s}\n", msg, msg.Value)
		}
	}
}
