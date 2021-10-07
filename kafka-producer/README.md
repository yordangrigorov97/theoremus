# Kafka Producer

## Quickstart

1. Provide some data inside _data/raw_gps_data.csv_
2. Make sure the settings in _conf.json_ are correct

## Usage Notes

You can watch the messages appear in Kafka on Kafka UI: _localhost:9080_

## Overview

The producer reads messages from the configured data file (default=_data/raw_gps_data.csv_) and sends them to Kafka (default=_kafka:9092_) if they contain Longitude and Latitude data.

## Improvements

Currently, the producer reads the file using the Pandas library. We should ideally write our own csv parser or use a more lightweight csv library

