import json
import sys
import uuid
import traceback
import boto3
from repository.dynamodb_repository import DynamoDBRepository


def main(user_id: str, user_token: str):
    print('hello')
    xxx = DynamoDBRepository()
    xxx.insert_user()




def lambda_handler(event, context):
    try:
        body = json.loads(event['body'])
        user_id = body['userId']
        user_token = body['userToken']
        main(user_id, user_token)
    except Exception as e:
        print(e)
        return {
            'statusCode': 500,
            'body': '{"message": "すでに作成されている場合"}'
        }


# python auth-anonymous-signup.py user_id user_token
if __name__ == '__main__':
    args = sys.argv
    user_id = args[1]
    user_token = args[2]
    body = { 
        "userId": user_id, 
        "userToken": user_token,
    }
    json_body = json.dumps(body)
    event = {}
    event['body'] = json_body
    lambda_handler(event, {})
