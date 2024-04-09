package repo

import (
	"context"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/channel-io/ch-app-store/internal/systemlog/model"
	"github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

const ddbTableName = "app_system_log"

type SystemLogRepository struct {
	ddbCli *dynamodb.DynamoDB
}

func NewSystemLogRepository(ddbCli *dynamodb.DynamoDB) *SystemLogRepository {
	return &SystemLogRepository{ddbCli: ddbCli}
}

func (s *SystemLogRepository) Save(ctx context.Context, input *model.SystemLog) error {
	ddbInput := &dynamodb.PutItemInput{
		Item:      marshalToDDBItem(input),
		TableName: aws.String(ddbTableName),
	}

	if _, err := s.ddbCli.PutItemWithContext(ctx, ddbInput); err != nil {
		return err
	}

	return nil
}

func (s *SystemLogRepository) Query(ctx context.Context, req *svc.QueryRequest) ([]*model.SystemLog, error) {
	ddbInput := &dynamodb.QueryInput{
		TableName:                 aws.String(ddbTableName),
		AttributesToGet:           allAttributes(),
		ExpressionAttributeNames:  keyNameExpression(),
		ExpressionAttributeValues: keyValueExpression(req),
		KeyConditionExpression:    rangeQueryExpression(req),
		Limit:                     aws.Int64(int64(req.Limit)),
	}

	output, err := s.ddbCli.QueryWithContext(ctx, ddbInput)
	if err != nil {
		return nil, err
	}

	return unmarshalToModels(output.Items)
}

func allAttributes() []*string {
	return []*string{
		aws.String("id"),
		aws.String("chatKey"),
		aws.String("channelId"),
		aws.String("message"),
		aws.String("appId"),
		aws.String("createdAt"),
		aws.String("expiresAt"),
	}
}

func keyValueExpression(req *svc.QueryRequest) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		":pk": {S: aws.String(toChatKey(req.ChatType, req.ChatId))},
		":sk": {S: aws.String(req.CursorID)},
	}
}

func keyNameExpression() map[string]*string {
	return map[string]*string{
		"#pk": aws.String("chatKey"),
		"#sk": aws.String("id"),
	}
}

func rangeQueryExpression(req *svc.QueryRequest) *string {
	switch req.Order {
	case svc.OrderAsc:
		return aws.String("#pk = :pk AND #sk > :sk")
	case svc.OrderDesc:
		return aws.String("#pk = :pk AND #sk < :sk")
	default:
		return aws.String("#pk = :pk AND #sk > :sk")
	}
}

func unmarshalToModels(ddbModels []map[string]*dynamodb.AttributeValue) ([]*model.SystemLog, error) {
	ret := make([]*model.SystemLog, 0, len(ddbModels))
	for _, ddbModel := range ddbModels {
		unmarshalled, err := unmarshalToModel(ddbModel)
		if err != nil {
			return nil, err
		}
		ret = append(ret, unmarshalled)
	}
	return ret, nil
}

func unmarshalToModel(ddbModel map[string]*dynamodb.AttributeValue) (*model.SystemLog, error) {
	var ret model.SystemLog
	ret.AppID = *ddbModel["id"].S
	ret.ChannelID = *ddbModel["channelId"].S
	ret.Message = *ddbModel["message"].S

	chatType, chatId := fromChatKey(*ddbModel["chatKey"].S)
	ret.ChatType = chatType
	ret.ChatId = chatId

	createdAt, err := strconv.ParseInt(*ddbModel["createdAt"].N, 10, 64)
	if err != nil {
		return nil, err
	}
	ret.CreatedAt = createdAt

	expiresAt, err := strconv.ParseInt(*ddbModel["expiresAt"].N, 10, 64)
	if err != nil {
		return nil, err
	}
	ret.ExpiresAt = expiresAt

	return &ret, nil
}

func marshalToDDBItem(input *model.SystemLog) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"id":        {S: aws.String(input.Id)},
		"chatKey":   {S: aws.String(toChatKey(input.ChatId, input.ChatId))},
		"channelId": {S: aws.String(input.ChannelID)},
		"message":   {S: aws.String(input.Message)},
		"appId":     {S: aws.String(input.AppID)},
		"createdAt": {N: aws.String(strconv.FormatInt(input.CreatedAt, 10))},
		"expiresAt": {N: aws.String(strconv.FormatInt(input.ExpiresAt, 10))},
	}
}

func toChatKey(chatType string, chatId string) string {
	return chatType + "-" + chatId
}

func fromChatKey(chatKey string) (string, string) {
	chatType, chatId, _ := strings.Cut(chatKey, "-")
	return chatType, chatId
}
