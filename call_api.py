import requests
import json

api_base_url = "https://44104577-93de-4813-9d9b-619db4249087.mock.pstmn.io/"

def get_user_role():

    api_url = '{0}get_user_role'.format(api_base_url)

    response = requests.get(api_url)

    if response.status_code == 200:
        return json.loads(response.content.decode('utf-8'))
    else:
        return None

get_user = get_user_role()

if get_user is not None:
    print("Here's your info: ")
    for k, v in get_user.items():
        print('{0}:{1}'.format(k, v))

else:
    print('[!] Request Failed')

