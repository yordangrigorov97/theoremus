def get_database():
    from pymongo import MongoClient

    CONNECTION_STRING = "mongodb://root:root@mongo:27017"

    client = MongoClient(CONNECTION_STRING)

    return client['theoremus']


def aggregate(fromTS, toTS, how):
    import arrow
    fromDT = arrow.get(fromTS).datetime
    toDT = arrow.get(toTS).datetime
    db = get_database()
    collection = db["vehicles"]

    # example 2021-09-24T01:40:02Z
    # time_filter = collection.find(
    #       {"data.date-time.system":{"$gte": fromDT, "$lte": toDT}})
    result = collection.aggregate([
        {"$match": {"data.date-time.system": {"$gte": fromDT, "$lte": toDT}}},
        {"$group": {
            "_id": {
                "vehicle-id": "$a",
                "IDDay": "$b"
                },
            "count": {"$sum": 1}
            }
        }])

    return str(result)
    # for item in item_details:
    #     # This does not give a very readable output
    #     print(item)
