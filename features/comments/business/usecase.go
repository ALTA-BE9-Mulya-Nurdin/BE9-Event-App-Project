package business

import (
	"be9/event/features/comments"
)

type commentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(cmnt comments.Data) comments.Business {
	return &commentUseCase{
		commentData: cmnt,
	}
}

func (uc *commentUseCase) GetDataAll() (data []comments.Core, err error) {
	data, err = uc.commentData.GetDataAll()
	return data, err
}
