def get_database():
    from pymongo import MongoClient
    import pymongo

    CONNECTION_STRING = "mongodb://root:root@mongo:27017"

    client = MongoClient(CONNECTION_STRING)

    return client['theoremus']
    
# This is added so that many files can reuse the function get_database()
def aggregate(fromTS, toTS, how):
    import arrow
    fromDT = arrow.get(fromTS).datetime
    toDT = arrow.get(toTS).datetime
    db = get_database()
    collection = db["vehicles"]

    # example 2021-09-24T01:40:02Z
    result = collection.find({"data.datetime.system":{"$gte": fromDT, "$lte": toDT}}).count()
    return str(result)
    # for item in item_details:
    #     # This does not give a very readable output
    #     print(item)
