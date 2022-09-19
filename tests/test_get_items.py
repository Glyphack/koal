import json

import requests

HOST = "http://localhost:8090/v1/"
REGISTER = "auth/register"
LOGIN = "auth/login"

payload = '{\n  "email": "s@ss.com",\n  "password": "yahoo@yahoo"\n}'
headers = {"Content-Type": "application/json"}
# response = requests.request("POST", url, data=payload, headers=headers)
# print(response.text)

# token = json.loads(response.text)["token"]

token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE2NjM1MjIyNTMsInN1YiI6InNAc3MuY29tIn0.M6wvoQcdWz6RyDmwrML6ru60mjBKfQkeYNOcOQ_x2WM"
headers = {
    "Content-Type": "application/json",
    "Authorization": f"bearer {token}",
}
payload = '{\n  "name": "pj2"\n}'
url = "http://localhost:8090/v1/todo/projects"
response = requests.request("POST", url, data=payload, headers=headers)
print(response.text)
project_id = json.loads(response.text)["createdProject"]["id"]

url = "http://localhost:8090/v1/todo/items"
payload = {
    "projectId": project_id,
    "title": "done-pj2",
    "description": "blahblah",
}
response = requests.request("POST", url, json=payload, headers=headers)
print(response.text)

querystring = {"projectIds": f"{project_id}"}
payload = ""
response = requests.request(
    "GET", url, data=payload, headers=headers, params=querystring
)

print(response.text)
