package line

//
//import (
//	"encoding/json"
//	//"encoding/json"
//	//"fmt"
//	//"net/http"
//	//"os"
//	"fmt"
//	"net/http"
//	"os"
//)
//
//type LineProfile struct {
//	UserId        string `json:"userId"`
//	DisplayName   string `json:"displayName"`
//	PictureUrl    string `json:"pictureUrl"`
//}
//
//func GetLineProfile(userId string)(lineProfile LineProfile,err error){
//	url := fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userId)
//	request, err := http.NewRequest("GET", url, nil)
//	request.Header.Set("Authorization", "Bearer "+ os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
//	if err != nil{
//		return
//	}
//	client := new(http.Client)
//	response, err := client.Do(request)
//	//遅延実行(returnの直前にBodyをcloseする。)
//	defer func() {
//		err = response.Body.Close()
//		if err != nil{
//			fmt.Println("line profile close failed")
//		}
//	}()
//	err = json.NewDecoder(response.Body).Decode(&lineProfile)
//	if err != nil{
//		return
//	}
//	return
//}
