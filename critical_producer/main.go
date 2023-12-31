package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	processor "github.com/zakhaev26/critical_producer/cpu"
	kafkaProducer "github.com/zakhaev26/critical_producer/kafka"
	mailer "github.com/zakhaev26/critical_producer/notifications"
	"github.com/zakhaev26/critical_producer/protobuf"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("OK")
	CriticalLog()
}

func CriticalLog() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env")
		return
	}
	CRITICAL_LOG_NODE_ID := os.Getenv("CRITICAL_LOG_NODE_ID")

	// thresholdMin := 5
	incidentCount := 0
	for {
		AvgCpuUsage, time_ := processor.FetchCpuUsage()
		if AvgCpuUsage < 80.0 {
			incidentCount++

			msg := &protobuf.KibanaMessage{
				CpuUsage:  AvgCpuUsage,
				Time:      time_,
				Timestamp: time.Now().Format(time.RFC3339),
			}

			data, err := proto.Marshal(msg)

			if err != nil {
				log.Fatal(err)
			}

			kafkaProducer.PushToKafka(CRITICAL_LOG_NODE_ID, string(data))

			content := `<!DOCTYPE html>
						<html lang="en">

						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>Watchdog - CPU Usage Alert</title>
							<style>
								body {
									font-family: 'Arial', sans-serif;
									background-color: #000;
									margin: 0;
									padding: 0;
									display: flex;
									justify-content: center;
									align-items: center;
									height: 100vh;
									color: #fff;
								}

								.container {
									background-color: #000;
									border-radius: 8px;
									box-shadow: 0 0 10px rgba(255, 255, 255, 0.1);
									padding: 20px;
									text-align: center;
									font-size: 20px;
									font-weight: 900;
								}

								h1 {
									color: #fff;
									font-weight: 1200;
									font-family: 'Courier New', Courier, monospace;
								}

								p {
									color: #ccc;
								}

								.cta-button {
									display: inline-block;
									background-color: #fff;
									color: #000;
									padding: 10px 20px;
									text-decoration: none;
									border-radius: 4px;
									transition: background-color 0.3s;

								}

								.cta-button:hover {
									background-color: #ddd;
								}
							</style>
						</head>

						<body>
							<div class="container">
								<h1>Watchdog Alert: High CPU Usage</h1>
								<p>Dear User,</p>
								<p>We have detected high CPU usage for ` + strconv.Itoa(incidentCount) + ` incidents on your system . This may indicate a potential issue that requires attention.</p>
								<p>Please take necessary actions to investigate and resolve the high CPU consumption.</p>
								<p>Thank you for using Watchdog to monitor your system.</p>
								<a href="#" class="cta-button">Login to Watchdog</a>
							</div>
						</body>

						</html>
			`

			_ = mailer.SendMail(string(content), "Cpu OverClocking - Watchdog Metrics", "b422056@iiit-bh.ac.in")
			time.Sleep(5 * time.Second)
		} else {
			continue
		}
	}
}
