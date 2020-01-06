package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func sendEmail(name string, email string, message string) error {
	session, err := session.NewSession()
	//(&aws.Config{
	//	Region: aws.String("us-east-1"),
	//})
	if err != nil {
		fmt.Println("NewSession error:", err)
		return err
	}

	topic := os.Getenv("SNS_TOPIC_ARN")
	topicMessage := fmt.Sprintf("Name: %v Email: %v Message: %v", name, email, message)
	fmt.Printf("Sending message %v to topic %v", topicMessage, topic)

	client := sns.New(session)
	input := &sns.PublishInput{
		Message:  aws.String(topicMessage),
		TopicArn: aws.String(topic),
	}

	result, err := client.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return err
	}

	fmt.Println(result)

	return nil
}
