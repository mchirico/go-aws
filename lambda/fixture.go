package lambda

func KinesisExample() string {
	return `{
		"Records": [
			{
				"kinesis": {
					"kinesisSchemaVersion": "1.0",
					"partitionKey": "1",
					"sequenceNumber": "49590338271490256608559692538361571095921575989136588898",
					"data": "SGVsbG8sIHRoaXMgaXMgYSB0ZXN0Lg==",
					"approximateArrivalTimestamp": 1545084650.987
				},
				"eventSource": "aws:kinesis",
				"eventVersion": "1.0",
				"eventID": "shardId-000000000006:49590338271490256608559692538361571095921575989136588898",
				"eventName": "aws:kinesis:record",
				"invokeIdentityArn": "arn:aws:iam::123456789012:role/lambda-role",
				"awsRegion": "us-east-2",
				"eventSourceARN": "arn:aws:kinesis:us-east-2:123456789012:stream/lambda-stream"
			},
			{
				"kinesis": {
					"kinesisSchemaVersion": "1.0",
					"partitionKey": "1",
					"sequenceNumber": "49590338271490256608559692540925702759324208523137515618",
					"data": "VGhpcyBpcyBvbmx5IGEgdGVzdC4=",
					"approximateArrivalTimestamp": 1545084711.166
				},
				"eventSource": "aws:kinesis",
				"eventVersion": "1.0",
				"eventID": "shardId-000000000006:49590338271490256608559692540925702759324208523137515618",
				"eventName": "aws:kinesis:record",
				"invokeIdentityArn": "arn:aws:iam::123456789012:role/lambda-role",
				"awsRegion": "us-east-2",
				"eventSourceARN": "arn:aws:kinesis:us-east-2:123456789012:stream/lambda-stream"
			}
		]
	}`
}

func SNSExample() string {
	return `{
		"Records": [
		  {
			"EventVersion": "1.0",
			"EventSubscriptionArn": "arn:aws:sns:us-east-2:123456789012:sns-lambda:21be56ed-a058-49f5-8c98-aedd2564c486",
			"EventSource": "aws:sns",
			"Sns": {
			  "SignatureVersion": "1",
			  "Timestamp": "2019-01-02T12:45:07.000Z",
			  "Signature": "tcc6faL2yUC6dgZdmrwh1Y4cGa/ebXEkAi6RibDsvpi+tE/1+82j...65r==",
			  "SigningCertUrl": "https://sns.us-east-2.amazonaws.com/SimpleNotificationService-ac565b8b1a6c5d002d285f9598aa1d9b.pem",
			  "MessageId": "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
			  "Message": "Hello from SNS!",
			  "MessageAttributes": {
				"Test": {
				  "Type": "String",
				  "Value": "TestString"
				},
				"TestBinary": {
				  "Type": "Binary",
				  "Value": "TestBinary"
				}
			  },
			  "Type": "Notification",
			  "UnsubscribeUrl": "https://sns.us-east-2.amazonaws.com/?Action=Unsubscribe&amp;SubscriptionArn=arn:aws:sns:us-east-2:123456789012:test-lambda:21be56ed-a058-49f5-8c98-aedd2564c486",
			  "TopicArn":"arn:aws:sns:us-east-2:123456789012:sns-lambda",
			  "Subject": "TestInvoke"
			}
		  }
		]
	  }`
}
