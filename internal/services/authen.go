package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"note-keeper-backend/internal/models"
	"note-keeper-backend/internal/social"
	"note-keeper-backend/internal/utils"
	"strconv"
)

const FacebookLogin = "facebook"
const GoggleLogin = "goggle"
const LenToken = 26

func (s *Service) LoginUsingSocialAccessToken(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	var r models.LoginUsingSocialAccessTokenRequest

	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("error while decoding request: " + err.Error())
		return
	}
	fmt.Printf("request %v", r)
	// get social info
	var socialInstant = &social.Social{}

	// maybe login by google,twitter later
	switch r.Type {
	case FacebookLogin:
		socialInstant = social.NewSocialInstant(social.FacebookAcc{})
	}

	// get social info from api of Facebook(or Goggle, Twitter, ...)
	socialInfo, err := socialInstant.PerformGetInfo(r.SocialToken)
	if err != nil {
		fmt.Println("error while performing getInfo: " + err.Error())
		return
	}

	fmt.Printf("\nsocialInfo: %v\n", socialInfo)

	if socialInfo.Id == "" {
		fmt.Println("invalid social token")
		http.Error(w, http.StatusText(401), 401)
		return
	}

	// check if already have row in tb_login_social
	loginSocialModel, existed, err := s.loginSocialStore.GetBySocialId(socialInfo.Id)

	if err != nil {
		fmt.Println("error while performing GetBySocialId: " + err.Error())
		return
	}
	// not existed
	if !existed {
		// 1.1 generate token
		accessToken := utils.GenerateSecureToken(LenToken)

		userModel := &models.ModelUser{
			UserName:    socialInfo.Name,
			AccessToken: accessToken,
		}

		tx := s.mainStore.Begin()

		// 1.2 create user in tb_users
		if err = tx.Create(userModel).Error; err != nil {
			fmt.Println("error while performing getInfo: " + err.Error())
			tx.Rollback()
			return
		}

		// 1.3 create row in tb_login_social
		switch r.Type {
		case FacebookLogin:
			loginSocialModel.TypeProvider = FacebookLogin
		}
		loginSocialModel.SocialId = socialInfo.Id
		loginSocialModel.UserId = userModel.UserId

		if err = tx.Create(loginSocialModel).Error; err != nil {
			fmt.Println("error while performing getInfo: " + err.Error())
			tx.Rollback()
			return
		}

		if err = tx.Commit().Error; err != nil {
			fmt.Println("error while commiting transaction: " + err.Error())
			return
		}

		// 1.4 response user info
		w.Header().Set("Content-Type", "application/json")
		account := models.Account{
			UserName:    userModel.UserName,
			AccessToken: userModel.AccessToken,
		}
		res := models.LoginUsingSocialAccessTokenResponse{
			Account: account,
		}
		res.FlagInfo = models.FlagInfo{
			Flag:    143,
			Message: "OK",
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling response: " + err.Error())
			return
		}
		w.Write(resJson)
		return
	} else {
		// 2.1. update access_token in tb_users
		updateTokenObject := map[string]interface{}{
			"access_token": utils.GenerateSecureToken(LenToken),
		}

		if err = s.userStore.UpdateMapByUserID(loginSocialModel.UserId, updateTokenObject); err != nil {
			fmt.Println("error while UpdateMapByUserID: " + err.Error())
			return
		} else {
			// 2.2. get user info
			userModel, _, err := s.userStore.GetByUserId(loginSocialModel.UserId)
			if err != nil {
				fmt.Println("error while getting user info: " + err.Error())
				return
			}

			// 2.3. response user info
			w.Header().Set("Content-Type", "application/json")
			account := models.Account{
				UserName:    userModel.UserName,
				AccessToken: userModel.AccessToken,
			}
			res := models.LoginUsingSocialAccessTokenResponse{
				Account: account,
			}
			res.FlagInfo = models.FlagInfo{
				Flag:    143,
				Message: "OK",
			}
			fmt.Printf("@res :%v", res)

			resJson, err := json.Marshal(res)
			if err != nil {
				fmt.Println("error while Marshaling response: " + err.Error())
				return
			}

			w.Write(resJson)
			return
		}
	}
}

func (s *Service) Logout(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error

	var res models.LogoutResponse

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	// delete access token
	if err = s.userStore.ClearAccessToken(uint32(userIdUint)); err != nil {
		fmt.Println("error while Marshaling response: " + err.Error())
		return
	}

	res.FlagInfo = models.FlagInfo{
		Flag:    143,
		Message: "OK",
	}
}
