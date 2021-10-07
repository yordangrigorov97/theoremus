# Kafka Consumer

## Usage notes

When starting docker-compose, it is normal to see a message or two from _Kafka Consumer_, saying that it cannot reach Kafka on the specified URI. The consumer will automatically retry the connection after 2 seconds, which is usually enough for Kafka to catch up and start allowing connections. The following messages from _Kafka Consumer_ should contain the messages that are being sent to MongoDB.

You can watch the messages appear in Mongo with Mongo Express on _localhost:27080_

## Overview

1. Fetches messages using the _kafka-go_ library
2. Adds keys for the day and hour contained in the message ("IDDay", "IDHour")
3. Sends the extended messsage to MongoDB
