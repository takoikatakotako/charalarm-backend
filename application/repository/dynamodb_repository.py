import json
import uuid
import boto3

# def func():
#     print('-- sub_mod1.func1 is called')

class DynamoDBRepository:
    def insert_user(self):
        dynamo_db = boto3.resource('dynamodb', endpoint_url='http://localhost:4566')
        table = dynamo_db.Table('user-table')
        item = {}
        item['userId'] = 'userId'
        item['userToken'] = 'userToken'
        table.put_item(Item=item)
        return {
            'statusCode': 200,
            'body': '{"message": "追加完了"}'
        }
