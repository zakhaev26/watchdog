package main

import (
	"fmt"
	"strconv"
	"time"

	processor "github.com/zakhaev26/critical_producer/cpu"
	kafkaProducer "github.com/zakhaev26/critical_producer/kafka"
	mailer "github.com/zakhaev26/critical_producer/notifications"
)

func main() {
	fmt.Println("OK")
	CriticalLog()
}

func CriticalLog() {
	thresholdMin := 5
	incidentCount := 0
	for {
		AvgCpuUsage := processor.FetchCpuUsage()

		if AvgCpuUsage < 80.0 {
			incidentCount++

			logMessage := "Critical event detected - CPU Usage: " + strconv.FormatFloat(AvgCpuUsage, 'f', 2, 64)
			kafkaProducer.PushToKafka("watchdog-critical-logs", logMessage)

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
			time.Sleep(time.Minute * time.Duration(thresholdMin))
		} else {
			continue
		}
	}
}
