import json
import sys
import uuid
import traceback
import boto3

def main(user_id: str, user_token: str):
    # dynamo_db = boto3.resource('dynamodb')
    dynamo_db = boto3.resource('dynamodb', endpoint_url='http://localhost:4566')
    table = dynamo_db.Table('user-table')
    response = table.get_item(Key={"userId": user_id})
    print(response)

    # Item が無く、新規作成の場合
    if 'Item' not in response:
        return {
            'statusCode': 200,
            'body': '{"message": "新規登録"}'
        }


    # すでにItemが作成されている場合
    item = response['Item']
    ios = item['ios']
    
    # アプリIDが登録されていない場合
    if app_id not in ios:
        ios[app_id] = {
            'sentEntryIds': []
        }
        item['ios'] = ios
        table.put_item(Item=item)
        return {
            'statusCode': 200,
            'body': '{"message": "追加完了"}'
        }


    # 何もしなくてokの場合
    return {
        'statusCode': 200,
        'body': '{"message": "何もしない"}'
    }




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
