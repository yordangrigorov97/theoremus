import os
AGGREGATION_TYPES = {"hour": "IDHour", "day": "IDDay"}


def get_database():
    from pymongo import MongoClient
    MONGO_URI = os.getenv('MONGO_URI')
    if MONGO_URI is None:
        print("please set MONGO_URI environment variable")
        exit(1)

    client = MongoClient(MONGO_URI)

    return client['theoremus']


def aggregate(fromTS, toTS, how):
    import arrow
    agg_field = AGGREGATION_TYPES[how]
    fromDT = arrow.get(fromTS).datetime
    toDT = arrow.get(toTS).datetime
    db = get_database()
    collection = db["vehicles"]

    # example 2021-09-24T01:40:02Z
    result = collection.aggregate(
        [
           {
              "$match": {
                 "data.date-time.system": {
                    "$gte": fromDT,
                    "$lte": toDT
                 }
              }
           },
           {
              "$group": {
                 "_id": {
                    "vehicle-id": "$vehicle-id",
                    f"{agg_field}": f"${agg_field}"
                 },
                 "count": {
                    "$sum": 1
                 }
              }
           }
        ]
                                )

    # print(f"{result=}")
    # print(f"{str(result)=}")
    # import ipdb; ipdb.set_trace()

    return str(list(result))
    # for item in item_details:
    #     # This does not give a very readable output
    #     print(item)
