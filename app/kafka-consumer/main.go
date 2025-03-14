package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	var offset int64 = -1
	var limit = -1 // -1 means all
	var debug = false
	var action, version, addresses, topic, username, password string
	flag.Int64Var(&offset, "offset", -1, "offset")
	flag.IntVar(&limit, "limit", -1, "limit")
	flag.StringVar(&action, "action", "consume", "action: consume(default), show-partition")
	flag.StringVar(&version, "version", "0.10.2.1", "version")
	flag.StringVar(&addresses, "bootstrap-server", os.Getenv("KAFKA_BOOTSTRAP_SERVER"), "bootstrap-server")
	flag.StringVar(&topic, "topic", os.Getenv("KAFKA_TOPIC"), "topic")
	flag.StringVar(&username, "username", os.Getenv("KAFKA_USERNAME"), "username")
	flag.StringVar(&password, "password", os.Getenv("KAFKA_PASSWORD"), "password")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	if debug {
		sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
	}
	if offset > 0 {
		offset = -offset // 避免越界
	}
	var err error

	addrs := strings.Split(addresses, ",")
	cfg := sarama.NewConfig()
	cfg.Version, err = sarama.ParseKafkaVersion(version)
	if err != nil {
		panic(err)
	}
	cfg.ClientID = "test"
	cfg.Metadata.Full = false
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.User = username
	cfg.Net.SASL.Password = password
	cfg.Net.SASL.Handshake = true
	cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	cfg.Net.SASL.Version = sarama.SASLHandshakeV0 // 这个必须正确
	client, err := sarama.NewClient(addrs, cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Close()
	}()
	ps, err := client.Partitions(topic)
	if err != nil {
		panic(err)
	}
	offsets := map[int32]int64{}
	for _, p := range ps {
		offsetNow, err := client.GetOffset(topic, p, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		offsets[p] = offsetNow + offset
		offsetOldest, err := client.GetOffset(topic, p, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		if offsets[p] < offsetOldest {
			offsets[p] = offsetOldest
		}
		if action == "show-partition" {
			fmt.Printf("patition %d: %d, %d\n", p, offsetNow, offsets[p])
		}
	}
	if action == "show-partition" {
		return
	}
	if limit == 0 {
		return
	}

	consumer, err := sarama.NewConsumer(addrs, cfg)
	if err != nil {
		panic(err)
	}

	var lock sync.Mutex
	syncRun := func(f func()) {
		lock.Lock()
		f()
		lock.Unlock()
	}

	count := 0
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	for _, partition := range ps {
		wg.Add(1)
		go func(partition int32) {
			defer wg.Done()
			// 设置从某个时间点开始消费
			//fmt.Printf("%d, %d\n", partition, offsets[partition])
			consumer, err := consumer.ConsumePartition(topic, partition, offsets[partition])
			if err != nil {
				fmt.Printf("partition %d: %d %s\n", partition, offsets[partition], err.Error())
				return
			}
			defer func() {
				_ = consumer.Close()
			}()
			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-consumer.Messages():
					syncRun(func() {
						if limit > 0 && count >= limit {
							cancel()
							return
						}
						count++
						fmt.Printf("%s\n", string(msg.Value))
						//fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
					})
				}
			}
		}(partition)
	}
	wg.Wait()
}
