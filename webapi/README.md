# Web API

## Usage Notes

Example queries:

http://localhost:8080/vehicles/2020-09-24T01:40:02Z/2022-09-24T01:40:02Z/day

http://localhost:8080/vehicles/2020-09-24T01:40:02Z/2022-09-24T01:40:02Z/hour

http://localhost:8080/vehicles/2022-09-24T01:40:02Z/2023-09-24T01:40:02Z/day

http://localhost:8080/vehicles/2022-09-24T01:40:02Z/2023-09-24T01:40:02Z/hour


## Overview

A Django app without a GUI, returns a JSON message in the format:

```json

[{'_id': {'vehicle-id': str, 'IDDay': datetime.datime(..)}, 'count': int}, ..]

```

or

```json

[{'_id': {'vehicle-id': str, 'IDHour': datetime.datime(..)}, 'count': int}, ..]

```

Depending on the aggregation type
