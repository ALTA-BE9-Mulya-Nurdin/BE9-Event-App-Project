package business

import "be9/event/features/categorys"

type categoryUseCase struct {
	categoryData categorys.Data
}

func NewCategoryBusiness(ctgory categorys.Data) categorys.Business {
	return &categoryUseCase{
		categoryData: ctgory,
	}
}

func (uc *categoryUseCase) GetDataAll() (data []categorys.Core, err error) {
	data, err = uc.categoryData.GetDataAll()
	return data, err
}
