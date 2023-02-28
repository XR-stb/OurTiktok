GATEWAY_SRC_PATH=apps/gateway
USER_SRC_PATH=apps/user
PUBLISH_SRC_PATH=apps/publish
FAVORITE_SRC_PATH=apps/favorite
FEED_SRC_PATH=apps/feed
COMMENT_SRC_PATH=apps/comment
RELATION_SRC_PATH=apps/relation
MESSAGE_SRC_PATH=apps/message

all: buildUser buildPublish buildFavorite buildFeed buildComment buildRelation buildMessage buildGateway

buildUser:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(USER_SRC_PATH)/main $(USER_SRC_PATH)/*.go

buildPublish:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(PUBLISH_SRC_PATH)/main $(PUBLISH_SRC_PATH)/*.go

buildFavorite:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(FAVORITE_SRC_PATH)/main $(FAVORITE_SRC_PATH)/*.go

buildFeed:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(FEED_SRC_PATH)/main $(FEED_SRC_PATH)/*.go

buildComment:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(COMMENT_SRC_PATH)/main $(COMMENT_SRC_PATH)/*.go

buildRelation:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(RELATION_SRC_PATH)/main $(RELATION_SRC_PATH)/*.go

buildMessage:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(MESSAGE_SRC_PATH)/main $(MESSAGE_SRC_PATH)/*.go

buildGateway:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(GATEWAY_SRC_PATH)/main $(GATEWAY_SRC_PATH)/*.go