package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// const protocol = "tcp"
const protocol = "mqtt"

// const broker = "jba41a51.ala.cn-hangzhou.emqxsl.cn"
const broker = "broker.emqx.io"

// const broker = "jba41a51.ala.cn-hangzhou.emqxsl.cn"
const port = 1883

// const port = 8883
const topic = "topic/volvo"
const username = "volvo"
const password = "volvo"

func main() {
	client := createMqttClient()
	go subscribe(client)        // 在主函数里, 我们用另起一个 go 协程来订阅消息
	time.Sleep(time.Second * 1) // 暂停一秒等待 subscribe 完成
	publish(client)
}

func createMqttClient() mqtt.Client {
	connectAddress := fmt.Sprintf("%s://%s:%d", protocol, broker, port)
	fmt.Println("Connecting to MQTT Broker " + connectAddress)
	rand.Seed(time.Now().UnixNano())
	clientID := fmt.Sprintf("go-client-%d", rand.Int())

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	opts.SetKeepAlive(time.Second * 60)

	// Optional: 设置CA证书
	// opts.SetTLSConfig(loadTLSConfig("caFilePath"))

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

func publish(client mqtt.Client) {
	qos := 0
	msgCount := 0
	for {
		if msgCount > 10 {
			break
		}
		payload := fmt.Sprintf("message: %d!", msgCount)
		if token := client.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
			fmt.Printf("publish failed, topic: %s, payload: %s\n", topic, payload)
		} else {
			fmt.Printf("publish success, topic: %s, payload: %s\n", topic, payload)
		}
		msgCount++
		time.Sleep(time.Second * 1)
	}
}

func subscribe(client mqtt.Client) {
	qos := 0
	client.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received `%s` from `%s` topic\n", msg.Payload(), msg.Topic())
	})
}

func loadTLSConfig(caFile string) *tls.Config {
	// load tls config
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = false
	if caFile != "" {
		certpool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(caFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		certpool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = certpool
	}
	return &tlsConfig
}
