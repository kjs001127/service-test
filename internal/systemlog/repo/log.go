package repo

import (
	"context"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/channel-io/ch-app-store/internal/systemlog/model"
	"github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

const ddbTableName = "app_system_log"

type SystemLogRepository struct {
	ddbCli *dynamodb.Client
}

func NewSystemLogRepository(ddbCli *dynamodb.Client) *SystemLogRepository {
	return &SystemLogRepository{ddbCli: ddbCli}
}

func (s *SystemLogRepository) Save(ctx context.Context, input *model.SystemLog) error {
	ddbInput := &dynamodb.PutItemInput{
		Item:      marshalToDDBItem(input),
		TableName: aws.String(ddbTableName),
	}

	if _, err := s.ddbCli.PutItem(ctx, ddbInput); err != nil {
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
		Limit:                     &req.Limit,
	}

	output, err := s.ddbCli.Query(ctx, ddbInput)
	if err != nil {
		return nil, err
	}

	return unmarshalToModels(output.Items)
}

func allAttributes() []string {
	return []string{
		"id",
		"chatKey",
		"channelId",
		"message",
		"appId",
		"createdAt",
		"expiresAt",
	}
}

func keyValueExpression(req *svc.QueryRequest) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		":pk": &types.AttributeValueMemberS{Value: toChatKey(req.ChatType, req.ChatId)},
		":sk": &types.AttributeValueMemberS{Value: req.CursorID},
	}
}

func keyNameExpression() map[string]string {
	return map[string]string{
		"#pk": "chatKey",
		"#sk": "id",
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

func unmarshalToModels(ddbModels []map[string]types.AttributeValue) ([]*model.SystemLog, error) {
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

func unmarshalToModel(ddbModel map[string]types.AttributeValue) (*model.SystemLog, error) {
	var ret model.SystemLog
	ret.AppID = ddbModel["id"].(*types.AttributeValueMemberS).Value
	ret.ChannelID = ddbModel["channelId"].(*types.AttributeValueMemberS).Value
	ret.Message = ddbModel["message"].(*types.AttributeValueMemberS).Value

	chatType, chatId := fromChatKey(ddbModel["chatKey"].(*types.AttributeValueMemberS).Value)
	ret.ChatType = chatType
	ret.ChatId = chatId

	createdAt, err := strconv.ParseInt(ddbModel["createdAt"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return nil, err
	}
	ret.CreatedAt = createdAt

	expiresAt, err := strconv.ParseInt(ddbModel["expiresAt"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return nil, err
	}
	ret.ExpiresAt = expiresAt

	return &ret, nil
}

func marshalToDDBItem(input *model.SystemLog) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id":        &types.AttributeValueMemberS{Value: input.Id},
		"chatKey":   &types.AttributeValueMemberS{Value: toChatKey(input.ChatId, input.ChatId)},
		"channelId": &types.AttributeValueMemberS{Value: input.ChannelID},
		"message":   &types.AttributeValueMemberS{Value: input.Message},
		"appId":     &types.AttributeValueMemberS{Value: input.AppID},
		"createdAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(input.CreatedAt, 10)},
		"expiresAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(input.ExpiresAt, 10)},
	}
}

func toChatKey(chatType string, chatId string) string {
	return chatType + "-" + chatId
}

func fromChatKey(chatKey string) (string, string) {
	chatType, chatId, _ := strings.Cut(chatKey, "-")
	return chatType, chatId
}
