package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/justinas/alice"
	"golang.org/x/crypto/bcrypt"

	"tessa/blockchain"
	"tessa/config"
	"tessa/database"
	"tessa/utils"
)

func apiHandlerAuth(middlewares alice.Chain, router *Router) {
	router.Get("/api/mnemonic", middlewares.ThenFunc(apiAuthMnemonic))
	router.Post("/api/signup", middlewares.ThenFunc(apiAuthSignup))
	router.Post("/api/login", middlewares.ThenFunc(apiAuthLogin))
	router.Get("/api/logout", middlewares.ThenFunc(apiAuthLogout))

	router.Post("/api/otpsend", middlewares.ThenFunc(apiAuthOTPSend))
	router.Post("/api/pinreset", middlewares.ThenFunc(apiAuthPinReset))
	router.Post("/api/otpverify", middlewares.ThenFunc(apiAuthOTPVerify))

}

func apiAuthMnemonic(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")
	json.NewEncoder(httpRes).Encode(Message{
		Code: http.StatusOK,
		Body: map[string]interface{}{
			"mnemonic": strings.Split(blockchain.ETHNewMnemonic(), " "),
		},
	})
}

func apiAuthSignup(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	message := &Message{Code: http.StatusInternalServerError}

	var formStruct struct {
		Username, Password,
		Confirm, Mnemonic,
		Email, Mobile,

		Title, Firstname, Lastname,
		Othername, Fullname, Gender,

		Street, City, State, Country,
		Occupation, Referrer,

		BVNNo, IDType, IDNumber, Image string
	}

	if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
		message.Message = "Error Decoding Form Values " + err.Error()
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	if formStruct.Mnemonic == "" {
		message.Message += "Mnemonic " + IsRequired
	} else {
		if !blockchain.ETHIsMnemonicValid(formStruct.Mnemonic) {
			message.Message += "Mnemonic Checksum failed. \n"
		}
	}

	if formStruct.Username == "" {
		message.Message += "Username " + IsRequired
	}

	if formStruct.Password == "" {
		message.Message += "Password " + IsRequired
	}

	if formStruct.Confirm == "" {
		message.Message += "Confirm Password " + IsRequired
	}

	if message.Message == "" {
		if formStruct.Password != formStruct.Confirm {
			message.Message += "Passwords do not match "
		}
	}

	if strings.HasSuffix(message.Message, "\n") {
		message.Message = message.Message[:len(message.Message)-2]
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	// var ownerID, campaignID uint64
	// if claims := utils.VerifyJWT(httpRes, httpReq); claims != nil {
	// 	ownerID, campaignID = apiClaims(claims)
	// }

	usersList := []database.Users{}
	sqlQuery := "select id from users where username = $1"
	config.Get().Postgres.Select(&usersList, sqlQuery, formStruct.Username)

	if len(usersList) > 0 {
		message.Message = fmt.Sprintf("Sorry this Username [%s] is taken", formStruct.Username)
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	mnemonic := formStruct.Mnemonic
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(formStruct.Password), bcrypt.DefaultCost)
	formStruct.Password = base64.StdEncoding.EncodeToString(config.Encrypt([]byte(passwordHash)))
	formStruct.Mnemonic = base64.StdEncoding.EncodeToString(config.Encrypt([]byte(formStruct.Mnemonic)))

	//Check if Mnemonic is duplicated, if so restore and replace
	var userID uint64
	user := database.Users{}
	config.Get().Postgres.Get(&userID, "select userid from wallets where mnemonic = $1", formStruct.Mnemonic)
	if userID > 0 {
		user.Update(
			map[string]interface{}{
				"ID":        userID,
				"FailedMax": 4,
				"Workflow":  "enabled",
				"Username":  formStruct.Username,
				"Password":  formStruct.Password,
			})
		message.Code = http.StatusOK
		message.Message = "Wallet Restored!! - Please Log In"
		json.NewEncoder(httpRes).Encode(message)
		return
	}
	//Check if Mnemonic is duplicated, if so restore and replace

	//All Seems Clear, Create New User Now Now
	profile := database.Profiles{}
	if formStruct.Firstname != "" || formStruct.Lastname != "" || formStruct.Email != "" ||
		formStruct.Mobile != "" || formStruct.Referrer != "" {

		fullname := ""
		if formStruct.Title != "" {
			fullname = formStruct.Title
		}

		if formStruct.Firstname != "" {
			fullname = fmt.Sprintf("%s %s", fullname, formStruct.Firstname)
		}

		if formStruct.Lastname != "" {
			fullname = fmt.Sprintf("%s %s", fullname, formStruct.Lastname)
		}
		formStruct.Fullname = fullname

		profile.Create(map[string]interface{}{
			"Workflow":   "enabled",
			"Title":      formStruct.Title,
			"Fullname":   formStruct.Fullname,
			"Firstname":  formStruct.Firstname,
			"Lastname":   formStruct.Lastname,
			"Othername":  formStruct.Othername,
			"Email":      formStruct.Email,
			"Mobile":     formStruct.Mobile,
			"City":       formStruct.City,
			"State":      formStruct.State,
			"Street":     formStruct.Street,
			"Country":    formStruct.Country,
			"Referrer":   formStruct.Referrer,
			"Occupation": formStruct.Occupation,
			"BVNNo":      formStruct.BVNNo,
			"IDType":     formStruct.IDType,
			"IDNumber":   formStruct.IDNumber,
			"Image":      formStruct.Image,
		})
	}

	user.Create(map[string]interface{}{
		"ProfileID": profile.ID,
		"FailedMax": 4,
		"Workflow":  "enabled",
		"Username":  formStruct.Username,
		"Password":  formStruct.Password,
	})

	wallet := database.Wallets{}
	wallet.Create(map[string]interface{}{
		"UserID":    user.ID,
		"ProfileID": profile.ID,
		"Mnemonic":  formStruct.Mnemonic,
		"Network":   "rinkeby",
		"Title":     "Tessa - " + utils.RandomString(3),
	})

	//generate account from wallet mnemonic
	account := database.Accounts{}
	_, fromAddress := blockchain.EthGenerateKey(mnemonic, 1)
	account.Create(map[string]interface{}{
		"Level":    1,
		"Priority": 1,
		"Title":    "Account 1",
		"Workflow": "enabled",
		"Address":  fromAddress.Hex(),

		"UserID":    user.ID,
		"WalletID":  wallet.ID,
		"ProfileID": profile.ID,
	})
	//generate account from wallet mnemonic

	message.Code = http.StatusOK
	message.Body = map[string]string{"Redirect": "/signin"}
	message.Message = "Wallet Created!! - Please Log In"

	//All Seems Clear, Create New User Now Now

	tableHits := database.Hits{}
	tableHits.Code = formStruct.Username
	tableHits.Title = fmt.Sprintf("New Client Signup: [%v] - %s", formStruct.Username, message.Message)

	tableHits.UserAgent = httpReq.UserAgent()
	tableHits.IPAddress = httpReq.RemoteAddr
	tableHits.Workflow = "enabled"
	// tableHits.Description = fmt.Sprintf("Fields: %v", formStruct)
	tableHits.Create(tableHits.ToMap())

	json.NewEncoder(httpRes).Encode(message)
}

func apiAuthLogin(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")
	message := &Message{Code: http.StatusInternalServerError}

	var formStruct struct {
		Username, Password, Path string
	}

	if err := json.NewDecoder(httpReq.Body).Decode(&formStruct); err != nil {
		message.Message = "Error Decoding Form Values " + err.Error()
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	// var ownerID, campaignID uint64
	// if claims := utils.VerifyJWT(httpRes, httpReq); claims != nil {
	// 	ownerID, campaignID = apiClaims(claims)
	// }

	user := database.Users{}
	sqlQuery := "select * from users where username = $1 limit 1"
	config.Get().Postgres.Get(&user, sqlQuery, formStruct.Username)

	if user.ID == 0 {
		message.Message = "Invalid Login!"
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	if user.Workflow != "enabled" {
		message.Message = "Username is " + user.Workflow
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	userMap := map[string]interface{}{"ID": user.ID}
	passwordHash, _ := base64.StdEncoding.DecodeString(user.Password)

	// if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(formStruct.Password)); err != nil {
	if err := bcrypt.CompareHashAndPassword(config.Decrypt(passwordHash), []byte(formStruct.Password)); err != nil {
		log.Println(err.Error())
		user.Failed++
		if user.FailedMax <= user.Failed {
			userMap["Workflow"] = "blocked"
			userMap["Failed"] = user.FailedMax
			message.Message = fmt.Sprintf("Account blocked - too many failed logins")
		} else {
			log.Println(user.Failed)
			userMap["Failed"] = user.Failed
			message.Message = fmt.Sprintf("%v attempts left", user.FailedMax-user.Failed)
		}
		user.Update(userMap)
		json.NewEncoder(httpRes).Encode(message)
		return
	}

	// All Seems Clear, Validate User Password
	userMap["Failed"] = uint64(0)
	user.Update(userMap)

	//Get the Profile Details
	profile := database.Profiles{}
	config.Get().Postgres.Get(&profile, "select * from profiles where userid = $1 limit 1", user.ID)
	//Get the Profile Details

	//Get Wallet List
	walletID := uint64(0)
	var myWallets []database.Wallets
	config.Get().Postgres.Select(&myWallets, "select id, title, network from wallets where userid = $1", user.ID)
	if len(myWallets) == 1 {
		walletID = myWallets[0].ID
		message.Body = map[string]string{"Redirect": "/dashboard"}
	} else {
		message.Body = map[string]interface{}{"Wallets": myWallets}
	}
	message.Message = "User Verified"
	message.Code = http.StatusOK
	//Get Wallet List

	// set our claims
	cookieExpires := time.Now().Add(time.Minute * 60)
	jwtClaims := jwt.MapClaims{
		"ID":        user.ID,
		"Username":  user.Username,
		"Email":     profile.Email,
		"Mobile":    profile.Mobile,
		"ProfileID": user.ProfileID,
		"WalletID":  walletID,
		"exp":       cookieExpires.Unix(),
	}

	if jwtToken, err := utils.GenerateJWT(jwtClaims); err == nil {
		cookieMonster := &http.Cookie{
			Name: config.Get().COOKIE, Value: jwtToken, Expires: cookieExpires, Path: "/",
		}
		http.SetCookie(httpRes, cookieMonster)
		httpReq.AddCookie(cookieMonster)
	}
	//All Seems Clear, Validate User Password and Generate Token

	tableHits := database.Hits{}
	tableHits.Code = formStruct.Username
	tableHits.Title = fmt.Sprintf("User Login: [%v] - %s", formStruct.Username, message.Message)
	tableHits.UserAgent = httpReq.UserAgent()
	tableHits.IPAddress = httpReq.RemoteAddr
	tableHits.Workflow = "enabled"
	// tableHits.Description = fmt.Sprintf("Fields: %v", formStruct)
	tableHits.Create(tableHits.ToMap())

	json.NewEncoder(httpRes).Encode(message)
}

func apiAuthLogout(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	message := new(Message)
	message.Code = http.StatusInternalServerError

	claims := utils.VerifyJWT(httpRes, httpReq)
	if claims == nil || claims["Username"] == nil {
		message.Code = http.StatusUnauthorized
		message.Error = "Invalid Authorization"
		message.Message = "Authorization Required"
		message.Body = map[string]string{"Redirect": "/"}
		return
	}

	username := "--unknown--"
	if claims["Username"] != nil {
		username = claims["Username"].(string)
	}

	tableHits := database.Hits{}
	tableHits.Code = username
	tableHits.Title = fmt.Sprintf("User Logout: [%v]", username)
	tableHits.UserAgent = httpReq.UserAgent()
	tableHits.IPAddress = httpReq.RemoteAddr
	tableHits.Workflow = "enabled"
	// tableHits.Description = fmt.Sprintf("Claims: %v", claims)
	tableHits.Create(tableHits.ToMap())

	message.Code = http.StatusOK
	message.Message = "Logout Successful"
	message.Body = map[string]string{"Redirect": "/"}

	cookieMonster := &http.Cookie{
		Name: config.Get().COOKIE, Value: "deleted", Path: "/",
		Expires: time.Now().Add(-(time.Hour * 24 * 30 * 12)), // set the expire time
	}
	http.SetCookie(httpRes, cookieMonster)
	httpReq.AddCookie(cookieMonster)

	json.NewEncoder(httpRes).Encode(message)
}

func apiAuthOTPSend(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusInternalServerError
	statusMessage := "Mobile number is needed to generate a One Time Pin"

	user := database.Users{}
	err := json.NewDecoder(httpReq.Body).Decode(&user)
	if err == nil {

		if user.Username != "" {

			userExists := database.Users{}
			sqlQuery := "select id, username from users where username = $1 limit 1"
			config.Get().Postgres.Get(&userExists, sqlQuery, user.Username)

			lSendOTP := true

			switch user.Code {
			case "reset":
				if userExists.ID == 0 {
					lSendOTP = false
					statusMessage = "Mobile number is not registered"
				}
			default:
				if userExists.ID != 0 {
					lSendOTP = false
					statusMessage = "Mobile number is registered"
				}
			}

			if lSendOTP {
				expirationTime := time.Duration(5)
				statusCode = http.StatusOK
				statusMessage = fmt.Sprintf("A One Time Pin has been sent to you - it expires in %d minutes", expirationTime)
				//set other otp linked to mobile to expire
				// config.Get().Postgres.Exec("update activations set workflow = 'expired' where userid = $1", user.ID)
				config.Get().Postgres.Exec("update activations set workflow = 'expired' where workflow='active' and mobile = $1", user.Username)

				//All Seems Clear, Generate One Time Pin and SMS User
				otpCODE := utils.RandomString(4)
				smsMessage := fmt.Sprintf("Your One Time Pin is: %s it expires in %d minutes", otpCODE, expirationTime)

				activation := database.Activations{}
				activation.Create(
					map[string]interface{}{
						"Code":        otpCODE,
						"Title":       "One Time Pin (OTP)",
						"Mobile":      user.Username,
						"UserID":      userExists.ID,
						"Workflow":    "active",
						"Description": smsMessage + " [sent to] " + user.Username,
						"Expirydate":  time.Now().Add(+(time.Minute * expirationTime)).Format(utils.TimeFormat),
					})

				if len(user.Username) > 10 {
					user.Username = user.Username[len(user.Username)-10:]
				}
				// user.Username = "+234" + user.Username
				// go utils.SendSMSTwilio("", user.Username, smsMessage)

				statusMessage = fmt.Sprintf("Your One Time Pin is [%v] ", otpCODE)
			}
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Message: statusMessage,
	})
}

func apiAuthOTPVerify(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusInternalServerError
	statusMessage := "Incorrect OTP"

	userOTP := database.Users{}
	err := json.NewDecoder(httpReq.Body).Decode(&userOTP)
	if err == nil {

		// user := database.Users{}
		// sqlQuery := "select username, id from users where username = $1 limit 1"
		// config.Get().Postgres.Get(&user, sqlQuery, userOTP.Username)

		if userOTP.Username != "" {
			activation := database.Activations{}
			// sqlQueryOTP := "select * from activations where workflow = 'active' and userid = $1 and code = $2 limit 1"
			sqlQueryOTP := "select * from activations where workflow = 'active' and mobile = $1 and code = $2 limit 1"
			config.Get().Postgres.Get(&activation, sqlQueryOTP, userOTP.Username, userOTP.Code)

			if activation.Code != "" {
				secondsExpired := utils.GetDifferenceInSeconds("", activation.Expirydate)
				if secondsExpired > 0 {
					//set otp to expired
					statusMessage = "Expired OTP"
					activation.Workflow = "expired"
				} else {
					//set otp to verified
					statusCode = http.StatusOK
					statusMessage = "OTP Verified"
					activation.Workflow = "verified"
				}
				activation.Update(activation.ToMap())
			}
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Message: statusMessage,
	})
}

func apiAuthPinReset(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	statusCode := http.StatusInternalServerError
	statusMessage := "PIN Reset Error"

	userOTP := database.Users{}
	err := json.NewDecoder(httpReq.Body).Decode(&userOTP)
	if err == nil {

		user := database.Users{}
		sqlQuery := "select username, id from users where username = $1 limit 1"
		config.Get().Postgres.Get(&user, sqlQuery, userOTP.Username)

		if user.Username != "" {
			activation := database.Activations{}
			sqlQueryOTP := "select * from activations where workflow = 'verified' and mobile = $1 and code = $2 limit 1"
			config.Get().Postgres.Get(&activation, sqlQueryOTP, user.Username, userOTP.Code)

			if activation.Code != "" {
				secondsExpired := utils.GetDifferenceInSeconds("", activation.Expirydate)
				if secondsExpired > 0 {
					//set otp to expired
					statusMessage += " - Expired OTP"
				} else {
					//set otp to verified
					statusMessage = "PIN Reset Successful"

					passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userOTP.Password), bcrypt.DefaultCost)
					user.Password = base64.StdEncoding.EncodeToString(config.Encrypt([]byte(passwordHash)))

					user.Failed = 0
					user.FailedMax = 5
					user.Workflow = "enabled"

					user.Update(user.ToMap())

					statusCode = http.StatusOK
				}
				activation.Workflow = "expired"
				activation.Update(activation.ToMap())
			}
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Message: statusMessage,
	})
}
