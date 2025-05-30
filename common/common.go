package common

import (
	"github.com/sunthewhat/secure-docs-api/type/shared"
	"github.com/sunthewhat/secure-docs-api/type/shared/query"
	"go.mongodb.org/mongo-driver/mongo"
)

var Config *shared.Config
var Gorm *query.Query
var Mongo *mongo.Database
