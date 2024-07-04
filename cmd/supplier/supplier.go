package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var sqsClient *sqs.Client

type Message struct {
	URI  string `json:"uri"`
	Body string `json:"body"`
}

func main() {
	sqsClient = initSQSClient()

	r := gin.Default()
	r.Use(ErrorMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/lambda/:input", messageHandler())

	r.Run(":8080")
}

func initSQSClient() *sqs.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return sqs.NewFromConfig(cfg)
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.Header("Content-Type", "application/json")

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": c.Errors.Last().Error(),
			})

			c.Abort()
			return
		}
	}
}

func sendMessageToSQS(client *sqs.Client, queueUrl string, msg Message) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(string(bytes)),
	}

	_, err = client.SendMessage(context.Background(), input)
	return err
}

func messageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		type InputURI struct {
			UriInput string `uri:"input" binding:"required"`
		}
		type InputJson struct {
			Payload string `json:"payload" binding:"required"`
		}

		var inputURI InputURI
		var inputJson InputJson

		if err := c.ShouldBindUri(&inputURI); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := c.ShouldBindJSON(&inputJson); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		sqsPrefix := "sqs"
		region := "ap-southeast-2"
		queueName := "HelperSQS"
		queueNum := "064657251892"

		// https://sqs.ap-southeast-2.amazonaws.com/064657251892/HelperSQS
		domain := fmt.Sprintf("https://%s.%s.amazonaws.com", sqsPrefix, region) //"ap-southeast-2.amazonaws.com/064657251892/"
		queueURL := fmt.Sprintf("%s/%s/%s", domain, queueNum, queueName)

		//var msg Message
		//
		//err := json.NewDecoder(c.Request.Body).Decode(&msg)
		//if err != nil {
		//	c.AbortWithError(http.StatusInternalServerError, err)
		//	return
		//}

		msg := Message{
			URI:  inputURI.UriInput,
			Body: inputJson.Payload,
		}

		if err := sendMessageToSQS(sqsClient, queueURL, msg); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(200, gin.H{"body": msg.Body})
	}
}
