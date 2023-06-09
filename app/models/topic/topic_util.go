// Package topic 工具
package topic

import (
	"gohub-lesson/pkg/app"
	"gohub-lesson/pkg/database"
	"gohub-lesson/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idStr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id", idStr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", clause.Column{Name: field}, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var m Topic

	database.DB.Select("id").
		Where("? = ?", clause.Column{Name: field}, value).
		Take(&m)

	return m.ID > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)

	return
}
