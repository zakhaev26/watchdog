# WatchDog - Server Monitoring Tool

![Alt text](image.png)
## Overview

WatchDog is a fault-tolerant server monitoring tool designed to empower server administrators with real-time insights into server health and performance. Its core functionality includes continuous CPU health log collection at one-second intervals. These logs are transferred to Elastic Cloud, where administrators can leverage Kibana for in-depth analytics and visualization of server loads. WatchDog also monitor's cpu health and usage and has a triggering mechanism that promptly notifies administrators via mails when server loads exceed predefined threshold values, ensuring proactive management of system performance.

## Features

1. Continuous Monitoring :WatchDog excels in providing continuous monitoring of CPU health, offering administrators a granular view of server performance with one-second interval logs.

2. Elastic Cloud Integration :The integration with Elastic Cloud and Kibana enables administrators to perform comprehensive analytics and visualize server loads, facilitating informed decision-making.

3. Alerting Mechanism :WatchDog features a triggering mechanism that sends timely notifications to administrators when server loads surpass predefined thresholds. This approach allows for prompt intervention to maintain optimal system performance.

## To build and run WatchDog from source, follow these steps:

> Note: MUST HAVE INSTALLED Go Compiler and Protoc Compiler.

Ensure Apache Kafka and Zookeeper are running either natively or using Docker. 

- Run Zookeeper:

``` bash

docker run -p 2181:2181 zookeeper
```

- Run Kafka Server:

```bash
docker run -p 9092:9092 \
-e KAFKA_ZOOKEEPER_CONNECT=<PRIVATE_IP>:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://<PRIVATE_IP>:9092 \
-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
confluentinc/cp-kafka
```
## Run the Mailing Microservice:

1. Build the Docker image:

```bash
docker build --build-arg EMAIL_SENDER_NAME="WatchDog" \
             --build-arg EMAIL_SENDER_ADDRESS=<mailingAddress@example.com> \
             --build-arg EMAIL_SENDER_PASSWORD=<YourEmailPassword> \
             -t mailing-api-image .
```

2. Run the Docker container:

```bash
    docker run -p 6969:6969 \
               -e EMAIL_SENDER_NAME="WatchDog" \
               -e EMAIL_SENDER_ADDRESS=<mailingAddress@example.com> \
               -e EMAIL_SENDER_PASSWORD=<YourEmailPassword> \
               mailing-api-image
```


## Simplifying the Process with Bash Script

- For a more straightforward process, use the provided bash script:

```bash

#!/bin/bash

# Set your environment variables
export EMAIL_SENDER_ADDRESS="your_email@example.com"
export EMAIL_SENDER_PASSWORD="YourEmailPassword"

# Build the Docker image
docker build --build-arg EMAIL_SENDER_NAME="WatchDog" \
             --build-arg EMAIL_SENDER_ADDRESS=<mailingAddress@example.com> \
             --build-arg EMAIL_SENDER_PASSWORD=<YourEmailPassword> \
             -t mailing-api-image .

# Run the Docker container
docker run -p 6969:6969 \
           -e EMAIL_SENDER_NAME="WatchDog" \
           -e EMAIL_SENDER_ADDRESS=<mailingAddress@example.com> \
           -e EMAIL_SENDER_PASSWORD=<YourEmailPassword> \
           mailing-api-image
```

- Make the script executable:

```bash
chmod +x run-mailing-api.sh
```

- Run the script:

```bash
./run-mailing-api.sh
```

This script encapsulates the build and run steps, making it easy to execute the entire process with a single command.

## Setting up Environment Variables
WatchDog provides a CLI tool that extracts crucial information such as Index Names of Log Nodes, API Keys, and SMTP Mail service configurations.

- Run the CLI Interface using :

```bash
go run cli/main.go
```

## Running WatchDog Microservices

For starting the execution of WatchDog microservices, a convenient bash script (watchdog-console-runner.sh) has been provided. This script automates the startup process for the monitoring components. Follow the steps below to run the microservices:

1.  Make the Bash Script Executable

```bash
chmod +x watchdog-console-runner.sh
```

2.  Run the Bash Script

```bash
./watchdog-console-runner.sh
```

The watchdog-console-runner.sh script orchestrates the execution of WatchDog microservices in the background.  These programs represent different components of WatchDog, such as frequent and critical consumers and producers.

By running the watchdog-console-runner.sh script, you initiate the WatchDog monitoring components, providing a way to set up the system for continuous monitoring.

<!-- 
## Comparison with Other Services

WatchDog stands out with its one-second interval log collection and seamless integration with Elastic Cloud. This allows for real-time analytics and visualization, providing a competitive edge over services like Prometheus. Additionally, WatchDog's alerting mechanism ensures proactive system management, distinguishing it as a comprehensive solution for server monitoring.
Apache Kafka and Elastic Search

WatchDog leverages Apache Kafka for efficient log transmission and Elastic Search for advanced analytics. Kafka facilitates scalable and fault-tolerant log streaming, while Elastic Search powers in-depth analysis and visualization through Kibana, making WatchDog a powerful and versatile server monitoring tool.
Conclusion

WatchDog is an tool for server administrators, offering a    approach to server monitoring with continuous log collection, seamless integration with Elastic Cloud, and a robust alerting mechanism. Its architecture, utilizing Apache Kafka and Elastic Search, ensures scalability and efficiency in monitoring server health and performance. -->