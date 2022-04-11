package social

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FacebookAcc struct{}

type FacebookInfoResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Picture  Data   `json:"picture"`
}

type Data struct {
	Data PictureData `json:"data"`
}

type PictureData struct {
	Url string `json:"url"`
}

func (fba FacebookAcc) GetInfo(accessToken string) (*SocialInfo, error) {
	// get info from facebook by access_token
	url := "https://graph.facebook.com/me/?fields=id,name,email,birthday,picture&access_token=" + accessToken

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error while getting info from facebook: %v", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Close = true
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("error while sending request to facebook to get user info: %v", err)
		return nil, err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal response
	var respObj *FacebookInfoResponse
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		return nil, err
	}

	finalResp := &SocialInfo{
		Id:       respObj.ID,
		Name:     respObj.Name,
		Birthday: respObj.Birthday,
		Email:    respObj.Email,
		Picture:  respObj.Picture.Data.Url,
	}

	return finalResp, nil
}
