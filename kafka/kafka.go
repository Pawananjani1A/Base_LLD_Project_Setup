package kafka

import (
	"base_app/config"
	"base_app/constants"
	"base_app/logger"
	"context"
	"fmt"
	"sync"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

// TODO: Setup proper producers and consumers to replicate in local
var (
	producer     *kafkago.Writer
	producerOnce sync.Once
	readerClient *kafkago.Reader
)

const (
	Value10  = 10
	Value5   = 5
	MaxBytes = 1000000
)

func CheckKafkaConnectivityAndTopic(brokers []string, topicName string) error {
	if config.GetServerEnvironment() != constants.Production {
		logger.Log.Debugf("Skipping kafka connectivity check for topic %s in %s environment", topicName, config.GetServerEnvironment())
		return nil
	}

	if len(brokers) == 0 {
		return fmt.Errorf("kafka brokers is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := kafkago.DialContext(ctx, "tcp", brokers[0]) // Check first broker
	if err != nil {
		return fmt.Errorf("failed to connect to broker %s: %w", brokers[0], err)
	}
	defer conn.Close()

	// Check if topic exists
	partitions, err := conn.ReadPartitions(topicName)
	if err != nil {
		return fmt.Errorf("failed to read partitions for topic %s: %w", topicName, err)
	}
	if len(partitions) == 0 {
		return fmt.Errorf("topic %s exists but has no partitions (or metadata issue)", topicName)
	}

	logger.Log.Infof("Successfully connected to broker and verified topic %s exists with %d partitions.", topicName, len(partitions))
	return nil
}

func SetupReaderClient() error {
	kafkaConfig := config.KafkaConf()

	dialerElementViewed := &kafkago.Dialer{
		Timeout:   Value10 * time.Second,
		DualStack: true,
		KeepAlive: Value5 * time.Second,
	}

	brokers := config.GetKafkaBrokers()
	topicName := kafkaConfig.TopicsMap.KafkaTopic

	err := CheckKafkaConnectivityAndTopic(brokers, topicName)
	if err != nil {
		logger.Log.Errorf("Error connecting to kafka for topic %s: %+v", topicName, err)
		panic(err)
	}

	readerConfig := kafkago.ReaderConfig{
		Brokers:               brokers,
		Dialer:                dialerElementViewed,
		StartOffset:           kafkago.FirstOffset,
		MaxWait:               Value5 * time.Second,
		CommitInterval:        Value5 * time.Second,
		MaxBytes:              MaxBytes,
		WatchPartitionChanges: true,
		Logger:                kafkago.LoggerFunc(logger.Log.Infof),
		ErrorLogger:           kafkago.LoggerFunc(logger.Log.Errorf),
	}

	logger.Log.Info(fmt.Sprintf("GroupId: %s and Topic: %s", kafkaConfig.GroupId, kafkaConfig.TopicsMap.KafkaTopic))

	readerConfig.GroupID = kafkaConfig.GroupId
	readerConfig.Topic = topicName
	readerClient = kafkago.NewReader(readerConfig)
	return nil
}

func SetupWriterClient(topicName string) *kafkago.Writer {
	producerOnce.Do(func() {
		var requiredAcks kafkago.RequiredAcks
		kafkaConfig := config.KafkaConf()

		if kafkaConfig.KafkaProducerRequiredAcks == 0 {
			requiredAcks = kafkago.RequireNone
		} else if kafkaConfig.KafkaProducerRequiredAcks == 1 {
			requiredAcks = kafkago.RequireOne
		} else {
			requiredAcks = kafkago.RequireAll
		}

		brokers := kafkaConfig.Brokers
		err := CheckKafkaConnectivityAndTopic(brokers, topicName)
		if err != nil {
			logger.Log.Errorf("Error connecting to kafka for writer topic %s: %+v", topicName, err)
			panic(err)
		}
		producer = &kafkago.Writer{
			Addr:         kafkago.TCP(brokers...),
			Topic:        topicName,
			Balancer:     &kafkago.Hash{},
			RequiredAcks: requiredAcks,
			BatchTimeout: Value10 * time.Millisecond,
		}
	})

	return producer
}
