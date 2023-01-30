/* Specify the name of the generated file's package. */ package copygen

import (
	"github.com/hyuti/pocketbase-clean-template/pkg/entity/model"
	"github.com/hyuti/pocketbase-clean-template/pkg/entity/model/usecase"
	useCaseModel "github.com/hyuti/pocketbase-clean-template/pkg/entity/model/usecase"
)

type Copygen interface {
	LoginInputToUserWhereInput(*usecase.LoginInput) *model.UserWhereInput
	PublicMeUseCaseUpdateInputToUserUpdateInput(
		*useCaseModel.PublicMeUseCaseUpdateInput,
	) *model.UserUpdateInput
}
