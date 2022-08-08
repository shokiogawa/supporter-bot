package command

import (
	"encoding/json"
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/domain/repository_interface"
	"net/http"
	"os"
)

type SaveUserUseCase struct {
	userRepository repository_interface.UserRepository
}

func NewSaveUserUseCase(userRepository repository_interface.UserRepository) *SaveUserUseCase {
	usecase := new(SaveUserUseCase)
	usecase.userRepository = userRepository
	return usecase
}

func (usecase *SaveUserUseCase) Invoke(userId string) (err error) {
	//TODO: インターフェース化すべき
	profile, err := GetLineProfile(userId)
	user := entity.NewUser(profile.UserId, profile.PictureUrl, profile.DisplayName)
	//userを使用し永続化
	err = usecase.userRepository.Save(user)
	if err != nil {
		return
	}
	return
}

//移動予定
type LineProfile struct {
	UserId      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureUrl  string `json:"pictureUrl"`
}

func GetLineProfile(userId string) (lineProfile LineProfile, err error) {
	url := fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userId)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil {
		return
	}
	client := new(http.Client)
	response, err := client.Do(request)
	//遅延実行(returnの直前にBodyをcloseする。)
	defer func() {
		err = response.Body.Close()
		if err != nil {
			fmt.Println("line profile close failed")
		}
	}()
	err = json.NewDecoder(response.Body).Decode(&lineProfile)
	if err != nil {
		return
	}
	return
}
