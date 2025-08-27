package converter

import (
	"github.com/tsfans/go/server/model/dto"
	"gorm.io/gorm"
)

func CommonToDto(common *gorm.Model) *dto.Common {
	return &dto.Common{
		ID:        common.ID,
		CreatedAt: common.CreatedAt,
		UpdatedAt: common.UpdatedAt,
	}
}
