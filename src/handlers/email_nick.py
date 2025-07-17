import json
import os
import boto3

def emailHandler(event, context):
    print("Received event:", json.dumps(event, indent=2))

    # Send message to SNS
    sns = boto3.client("sns")
    topic_arn = os.environ.get("SNS_TOPIC_ARN")
    if not topic_arn:
        raise ValueError("SNS_TOPIC_ARN environment variable is not set")
    
    if "body" not in event:
        raise ValueError("Event body is missing")
    

    body = json.loads(event["body"])
    topicMessage = f"Name: {body["name"]} Email: {body["email"]} Message: {body["message"]}"

    sns.publish(
        TopicArn=topic_arn,
        Message=json.dumps(topicMessage),
        Subject=f"New Email for Nick the Cloud Guy from {body["name"]}"
    )

    return {
        "statusCode": 200,
        "headers": {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "https://nickthecloudguy.com"
        },
        "body": json.dumps({
            "message": f"Message sent to {body["email"]} successfully!",
        })
    }
