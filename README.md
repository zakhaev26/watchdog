First Run a Zookeeper Server to manage zkNodes and Kafka Brokers.

```bash
# Run Using Docker (preferrably dont use detach mode to keep on checking whether the Zookeeper instance is alive or dead: 
 docker run -p 2181:2181 zookeeper

```

Then run a Kafka Server on 9092 Port with zookeeper PORT(default : 2181),Kafka's own Server Port,Replication Factor=1
 (We don't really care for fault tolerance at the moment) 

```bash
# Using confluentinc/cp-kafka image 
docker run -p 9092:9092 \
-e KAFKA_ZOOKEEPER_CONNECT=<YOUR_PRIVATE_IP>:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://<YOUR_PRIVATE_IP>:9092 \
-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
confluentinc/cp-kafka
```
Kafka Server Configuration Done.

```bash  
  cd server 
  # I used yarn for package management,install it using:
  npm i -g yarn
  yarn install   
```

Node server setup done.

Create a topic in Kafka server using: 
```bash
  node admin.js 
  # Creates a topic named football-score-updates
```

Topic Creation done.

Start the node server : 
```bash
npm start 
```

Copy the path of subscriber.html and publisher.html , paste in browser and test it out
