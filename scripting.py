import requests
import os


APITOKEN = os.environ["APITOKEN"]

url = f"https://webtranslateit.com/api/organizations/{APITOKEN}/collaborations?email={email}"

payload = {}
headers = {}
emails = [""]  # emails here

for email in emails:
    response = requests.request("DELETE", url, headers=headers, data=payload)
    if response.status_code != 200:
        print(f'{email}: {response.status_code}')
