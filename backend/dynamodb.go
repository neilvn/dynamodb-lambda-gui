package backend

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamodb2 "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (a *App) ListTables() ([]string, error) {
	tables := []string{}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	input := &dynamodb.ListTablesInput{}

	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
			return []string{}, err
		}

		// Append table names to our slice
		for _, n := range result.TableNames {
			tables = append(tables, *n)
		}

		// If LastEvaluatedTableName is nil, we've retrieved all tables
		if result.LastEvaluatedTableName == nil {
			break
		}

		// Set the ExclusiveStartTableName for the next iteration
		input.ExclusiveStartTableName = result.LastEvaluatedTableName
	}

	return tables, nil
}

func (a *App) ScanTable(tableName string) ([]map[string]interface{}, error) {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}

	// Create a DynamoDB client
	client := dynamodb2.NewFromConfig(cfg)

	// Prepare the scan input
	input := &dynamodb2.ScanInput{
		TableName: aws.String(tableName),
	}

	// Perform the scan operation
	result, err := client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %v", err)
	}

	// Convert the result to []map[string]interface{}
	var items []map[string]interface{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal items: %v", err)
	}

	return items, nil
}
