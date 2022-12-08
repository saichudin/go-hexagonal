package merchant

import (
	"e-menu-tentakel/core/model"
)

type WeborderAdapter interface {
	GetDetailWeblink(weblinkUrl string) (*model.WebLinkUri, error)
}
