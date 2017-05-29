package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	gocf "github.com/crewjam/go-cloudformation"
	sparta "github.com/mweagle/Sparta"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gopkg.in/telegram-bot-api.v4"
)

const appName = "UnsplashTelegramBot"

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
	case "featured":
		logger.Info("Featured click")
		photos, err := getFeaturedImages(logger)
		if err != nil {
			logger.Panic("ERROR: ", err.Error())
		}

		if len(photos) == 0 {
			logger.Info("There is no photos")
		}
		for _, photo := range photos {
			sendPhoto(&message, logger, photo)
		}
	case "latest":
		logger.Info("Latest click")
		photos, err := getLatestImages(logger)
		if err != nil {
			logger.Panic("ERROR: ", err.Error())
		}

		for _, photo := range photos {
			sendPhoto(&message, logger, photo)
		}
	case "search":
		logger.Info("Search: ", message.CommandArguments())
		photos, err := searchImages(message.CommandArguments(), logger)
		if err != nil {
			logger.Panic("ERROR: ", err.Error())
		}
		if len(photos) == 0 {
			logger.Info("There is no photos")
		}
		for _, photo := range photos {
			sendPhoto(&message, logger, photo)
		}
	default:
		logger.Info("Without command - exiting")
	}
}

func appendTelegramLambda(api *sparta.API, lambdaFunctions []*sparta.LambdaAWSInfo) []*sparta.LambdaAWSInfo {
	envs := make(map[string]*gocf.StringExpr)
	envs["BOT_TELEGRAMID"] = &gocf.StringExpr{Literal: telegramID}
	envs["BOT_UNSPLASHID"] = &gocf.StringExpr{Literal: unsplashID}

	loptions := &sparta.LambdaFunctionOptions{
		Timeout:     10,
		Environment: envs,
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

	err := parseConfig()
	if err != nil {
		log.Panic("Error parsing config ", err.Error())
	}

	// Get --tags options
	parseErrors := sparta.ParseOptions(nil)
	if parseErrors != nil {
		os.Exit(2)
	}

	var StackName string

	// If --tags is production name -Production else use -Devel
	// for separate environments
	if sparta.OptionsGlobal.BuildTags == "production" {
		StackName = appName + "-Production"
	} else {
		StackName = appName + "-Devel"
	}

	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway(StackName, stage)

	sparta.Main(StackName,
		"Unsplash Telegram Bot",
		spartaLambdaData(apiGateway),
		apiGateway,
		nil)
}
