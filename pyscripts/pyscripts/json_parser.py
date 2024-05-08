import json

# Opening JSON file
with open('../datadogresults.json') as f:

    # returns JSON object as
    # a dictionary
    data = json.load(f)

    # Iterating through the json
    # list
    for i in data['data']['buckets']:
        print(i["by"]["@pandora.payment.reference"])
