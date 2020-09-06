package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"strings"
	"time"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/facts cats"),
		tgbotapi.NewKeyboardButton("/facts numbers"),
		tgbotapi.NewKeyboardButton("/facts random"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/quote star wars"),
		tgbotapi.NewKeyboardButton("/quote anime"),
		tgbotapi.NewKeyboardButton("/quotes programming"),
	),
)

func CatFacts(catfact *Cat, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetCatFact(catfact)
	newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, catfact.Fact)
	_, _ = bot.Send(newmsg)
}

func ShowKeyboard(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose one the the following")
	msg.ReplyMarkup = numericKeyboard
	_, _ = bot.Send(msg)
}

func ProgramFacts(programFacts *Program, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetProgrammingFact(programFacts)
	newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, programFacts.En+"\n-"+programFacts.Author)
	_, _ = bot.Send(newmsg)
}

func NumberFacts(numbersFact *Fact, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetNumberFact(numbersFact)
	newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, numbersFact.Text)
	_, _ = bot.Send(newmsg)
}

func RandomFacts(randomFacts *Fact, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetRandomFact(randomFacts)
	newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, randomFacts.Text)
	_, _ = bot.Send(newmsg)
}

func StarwarsQuote(starwarsQuote *StarWars, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	GetStarWarsQuote(starwarsQuote)
	newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, starwarsQuote.Quote)
	_, _ = bot.Send(newmsg)
}

func cases(update tgbotapi.Update, catfact Cat, bot *tgbotapi.BotAPI, numbersFact Fact, programFacts Program, randomFact Fact, starwarsQuote StarWars) {
	message := strings.ToLower(update.Message.Text)
	if strings.Contains(message, "/facts cats") {
		CatFacts(&catfact, update, bot)
	} else if strings.Contains(message, "/facts numbers") {
		NumberFacts(&numbersFact, update, bot)
	} else if strings.Contains(message, "/quotes programming") {
		ProgramFacts(&programFacts, update, bot)
	} else if strings.Contains(message, "/facts random") {
		RandomFacts(&randomFact, update, bot)
	} else if strings.Contains(message, "/quote star wars") {
		StarwarsQuote(&starwarsQuote, update, bot)
	} else if strings.Contains(message, "/quote anime") {
		data := GetAnimeQuotes()
		resultingMap := AnimeStatus{}
		if err := json.Unmarshal(data, &resultingMap); err != nil {
			panic(err)
		}
		newmsg := tgbotapi.NewMessage(update.Message.Chat.ID, resultingMap.Data[0].Quote+"\n- "+resultingMap.Data[0].Character+" from "+resultingMap.Data[0].Anime)
		_, _ = bot.Send(newmsg)
	} else if strings.Contains(message, "/stop") {
		bot.StopReceivingUpdates()
	} else {
		ShowKeyboard(update, bot)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("Bot_Key")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	//random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//get cat facts
	var catfact Cat

	//var for programming facts
	var programFacts Program

	//var for numbers facts
	var numbersFact Fact

	//var for random facts
	var randomFact Fact

	//var  for starwarsQuote
	var starwarsQuote StarWars

	//gets the latest updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//makes a channel to reply
	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	//go through updates and rely based on message
	for update := range updates {
		//check for the null case
		if update.Message == nil {
			continue
		}

		//handle for private message
		if update.Message.Chat.IsPrivate() {
			cases(update, catfact, bot, numbersFact, programFacts, randomFact, starwarsQuote)
			// check if message in group was to bot
		} else if bot.IsMessageToMe(*(update.Message)) {
			cases(update, catfact, bot, numbersFact, programFacts, randomFact, starwarsQuote)
			//check if the reply message sent is from bot and not nil
		} else if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.From.UserName == "RandomFact1961Bot" {
			cases(update, catfact, bot, numbersFact, programFacts, randomFact, starwarsQuote)
		}
	}
}
