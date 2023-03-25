package main

import (
	"chlp/kafka-example-food-delivery/cmd"
	"chlp/kafka-example-food-delivery/internal/helper"
	"chlp/kafka-example-food-delivery/internal/kafka"
	"chlp/kafka-example-food-delivery/internal/refrigerator"
	"context"
	"errors"
	"log"
)

func main() {
	mainCtx, mainCtxCancel := context.WithCancel(context.Background())

	r := kafka.NewReader(
		mainCtx,
		[]string{cmd.KafkaAddr},
		cmd.KafkaLogin,
		cmd.KafkaPassword,
		cmd.KafkaGroup,
		cmd.KafkaTopic,
	)
	if r == nil {
		log.Fatalln("kafka reader is nil", errors.New("nil reader"))
	}
	kafka.SubscribeWithHandler(mainCtx, r, refrigerator.GotNewFoodHandler)

	helper.ListenSigInt()
	helper.CloseApp(mainCtxCancel)
}
