package merchant

import (
	"go-hexagonal/core/model"
)

type WeborderAdapter interface {
	GetDetailWeblink(weblinkUrl string) (*model.WebLinkUri, error)
}
