package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gopkg.in/telegram-bot-api.v4"
)

var unsplashID string
var telegramID string

func parseConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	//err := viper.ReadInConfig()
	//if err != nil {
	//	logger.Printf("Unable to open config file: %s, using env vars", err.Error())
	//}
	viper.SetEnvPrefix("bot")
	viper.BindEnv("unsplashID")
	viper.BindEnv("telegramID")

	unsplashID = viper.GetString("unsplashID")
	telegramID = viper.GetString("telegramID")

	if (unsplashID == "") || (telegramID == "") {
		return errors.New("You should use config file or environment variables")
	}

	return nil

}

func echoTelegram(event *json.RawMessage, context *sparta.LambdaContext, w http.ResponseWriter, logger *logrus.Logger) {

	var message tgbotapi.Message
	value := gjson.Get(string(*event), "body.message")
	err := json.Unmarshal([]byte(value.Raw), &message)
	if err != nil {
		logger.Error("Failed to unmarshal event data: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	logger.Info("Usuario :", message.Chat.UserName)

	switch message.Command() {
	case "start":
		logger.Info("start")
		txt := "Wellcome to unsplash bot, you can use /random and /search commands at the moment."
		sendHTML(&message, logger, txt)
		txt = "All photos are downloaded from <a href=\"https://unsplash.com\">unsplash.com</a>"
		sendHTML(&message, logger, txt)
	case "random":
		logger.Info("random: ", message.CommandArguments())
		photo, err := getRandomImage(logger)
		if err != nil {
			logger.Panic("ERROR: ", err.Error())
		}
		sendPhoto(&message, logger, photo)
	case "search":
		logger.Info("Search: ", message.CommandArguments())
		photos, err := searchImages(message.CommandArguments(), logger)
		if err != nil {
			logger.Panic("ERROR: ", err.Error())
		}
		if len(photos) == 0 {
			logger.Info("No hay fotos")
		} else {
			//TODO: Only get 3 images
			for v := 0; v < len(photos); v++ {
				sendPhoto(&message, logger, photos[v])
				if v == 2 {
					break
				}
			}
		}
	default:
		logger.Info("Sin comando - saliendo")
	}
}

func appendTelegramLambda(api *sparta.API, lambdaFunctions []*sparta.LambdaAWSInfo) []*sparta.LambdaAWSInfo {
	loptions := &sparta.LambdaFunctionOptions{
		Timeout: 10,
	}
	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{}, echoTelegram, loptions)
	apiGatewayResource, _ := api.NewResource("/v1/bot", lambdaFn)
	apiGatewayResource.NewMethod(http.MethodPost, http.StatusCreated)

	return append(lambdaFunctions, lambdaFn)
}

func spartaLambdaData(api *sparta.API) []*sparta.LambdaAWSInfo {
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = appendTelegramLambda(api, lambdaFunctions)
	return lambdaFunctions
}

func main() {

	parseConfig()

	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway("MySpartaApi", stage)

	stackName := "SpartaApplication"
	sparta.Main(stackName,
		"Simple Sparta Application",
		spartaLambdaData(apiGateway),
		apiGateway,
		nil)
}
