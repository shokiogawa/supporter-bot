package line_controller

import (
	"household.api/src/usecase/command"
)

type UserController struct {
	saveUseCase command.SaveUserUseCase
}

func NewUserController(saveUseCase command.SaveUserUseCase) *UserController {
	con := new(UserController)
	con.saveUseCase = saveUseCase
	return con
}

func (con *UserController) SaveUser(userId string) (replyMessage string, err error) {
	err = con.saveUseCase.Invoke(userId)
	if err != nil {
		return
	}
	replyMessage = "ユーザー登録に成功しました！"
	return
}
