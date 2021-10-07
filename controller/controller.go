package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"net/http"
	"os"
)

func getInput(){

}

func RunServer() {
	startServer()
}

func startServer(){
	r := gin.Default()
	userRoute := r.Group("/razorpay")
	userRoute.GET("/", postMessage)

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func configs(){

	err := godotenv.Load("environment.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/receive", slashCommandHandler)

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)
}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("SLACK_VERIFICATION_TOKEN")) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/hyperint":
		params := &slack.Msg{Text: s.Text}
		fmt.Printf("Response Text: %v",params.Text)
		response := fmt.Sprintf("Response %v", params.Text)
		w.Write([]byte(response))

	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func postMessage(c * gin.Context){
	api:= slack.New("xoxb-2533926746038-2570903409222-D3xdEvN9CBt2e62tPifofuFg")
	channelId,timestamp,err := api.PostMessage("C02GZQ3U3S7",
		slack.MsgOptionText("Hello World",false))

	if err!=nil{
		fmt.Println("%s\n" ,err)
	}

	fmt.Printf("Message sent successfully to channel %s at %s" , channelId, timestamp)

}
