package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"io/ioutil"
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

func makeRequest(){
	url := "https://hypertrace-ui.razorpay.com/graphql"
	fmt.Println("URL:>", url)

	requestBody,err := json.Marshal(map[string]interface{}{
		"query": "{\\n  spans(\\n    limit: 1000\\n    between: {startTime: \\\"2021-10-05T13:18:03.590Z\\\", endTime: \\\"2021-10-05T14:18:03.590Z\\\"}\\n    filterBy: [{operator: EQUALS, value: \\\"8ab7f71bf02d0eff\\\", type: ID, idType: API_TRACE }]\\n    orderBy: [{ key: \\\"duration\\\"}]\\n  ) {\\n    results {\\n      id\\n      logEvents {\\n        results {\\n          attributes: attribute(key: \\\"attributes\\\")\\n          timestamp: attribute(key: \\\"timestamp\\\")\\n          summary: attribute(key: \\\"summary\\\")\\n          __typename\\n        }\\n        __typename\\n      }\\n      displayEntityName: attribute(key: \\\"displayEntityName\\\")\\n      displaySpanName: attribute(key: \\\"displaySpanName\\\")\\n      duration: attribute(key: \\\"duration\\\")\\n      endTime: attribute(key: \\\"endTime\\\")\\n      parentSpanId: attribute(key: \\\"parentSpanId\\\")\\n      protocolName: attribute(key: \\\"protocolName\\\")\\n      spanTags: attribute(key: \\\"spanTags\\\")\\n      startTime: attribute(key: \\\"startTime\\\")\\n      type: attribute(key: \\\"type\\\")\\n      errorCount: attribute(key: \\\"errorCount\\\")\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n",
		})
	if err!=nil{
		fmt.Println("Unable To Marshal")
		return
	}

	cookie := `_gaexp=GAX1.2.3mdAZUKqSLClB8mBuzttMQ.18983.1; _gcl_au=1.1.900131272.1633334581; firstAttribUtm={"utm_source":"","utm_campaign":"","utm_medium":"","utm_term":"","utm_content":"","utm_adgroup":"","timestamp":"2021-10-04T13:33:01%2B05:30"}; lastAttribUtm={"utm_source":"direct","utm_medium":"website","utm_campaign":"direct","utm_content":"not set","utm_keyword":"not set"}; rzp_utm={"attributions":[],"first_page":"admin-dashboard.razorpay.com/signin","new_user":true}; clientId=d08ce78b-a010-4b2c-a8d7-a7792ff2bb2d; _rdt_uuid=1633334581845.40ff8dc3-47d7-4b28-8270-3fc009c45048; _ga=GA1.2.809373854.1633334582; _gid=GA1.2.10977730.1633334582; ajs_anonymous_id=%228e699ec7-79e6-446a-804a-927ce8636802%22; _fbp=fb.1.1633334589056.1747874148; ajs_user_id=%22HsbBfFzeXZ1qyX%22; WZRK_G=253a7b4d7f4e4589a72244c82dea0c16; campaignStartTime=Mon, 04 Oct 2021 13:45:57 GMT; _uetsid=823724a024e911eca5f38fa1fc085271; _uetvid=8237ae1024e911ec9f74e92cb94abe38; _sso_proxy=aRXUlwLVBLf7SyKuf_oEWL341jCVNCqsNVRELo4C_q5tE9UivFqnu3E258xKCvUHLgJMPIAAmHWWRXCrJAUGZlz1Y5eJfdJrMxwqpZLPYvdgwZYwUKaF4jX9NyayC-xCLStINfSzC5sop7Qig1iQC49nJJLStdyDK3y8-6-JrqmdxCoTET0r8u6cjxzQLf_6FXypOJ8yzAbHwlf_nz72RK1CJg-g_aMx1cijhObR2pAt9M-6kidhNsrYtu3YB3KNvxnndNEft7g5_U6qkStSpLbH7mr5zXGuqJzqYAs5y4W9ancn_XewFkVrWrobruqQb31s7Zdw3KMDygxWg_zCahgXw14249u1MXMpf8OoNQt7dHvrnyScxxh-KAOMRCDe881EFw2tU4pW-z7-LZdHwHk__PbFTqyRFsaNG_dHu1FaxfnsAopQKjExVkWD6QGnEmmLsuWDkHx4qwxxjCXRB-fd9j5e5pjJa1sXD5gkC50-LfXA6xE1cLbLvHeGQ4M0w7w4hThdCcUfjDmedtv5WqJvsaFd46Il5Muk-_IfZC5YDY-4DWWe4clrBHDsQ8AhP5NitNVoNyfA0S1ec0l0szGZjlq60upSdRiYUTv1MT6xbLsIFNAyeGB46MmovGLD3cJEHDkvNgtrKGL_ZMTtHA5RcrkZcQgt4OuXtDFPmUgQwd_7SMjtPLYBBWKg-If0Mtd408LPa7itgKAvUHs9sASNnDanTS8jj08mUR66ezpcpxY4tYcVNWUwKM8jVg-sPhYxmHNRqQL7H_vhjficcGn8xMvi3ytD; _sso_proxy=wdVpPXnkMynNxBmlAvhj4xQhE7EjXjWkQMH-sBuejZfyS_hn8hAmue4hj34tfbfiOsAOt6PcLki5ZUWf7HTl2KbnF4vg58i5fKwcwZQv3xoMDltmLOrcxtEsla2ArF84VVe9CZ9Pa9MLZBqMyzFmS-SaxjFVhxj6w4HUOqt53DH-2qVGpPGzRlKbxFrUavQznXjVWQX-Zx589I0J5r_G20rdnttHaSelPzklLKDclBgS3I02ceptsNB1xV9QXsx3Qui6WK9IOs8YUutBijWh-iJ2AimFYuEeu-7nNRbAQXJuNduidXS9G3XUYwOaTJ0YJY2j1wNqlID_dhfc6YvSVw2BOVFR0evRXXMaeVGtju2LX3ZFfu_ONx4Oc9c5rmRP5rouF73xh3gE_Mc9F2HKzBcyyeppOGQhu_WtoNMZs6OLO8OfM1TOq5C9AuIHi1LLb7y9koAf-nlTqelYo1gPml6lXsVVcZCyK9NfUtiBRoKwDZuVzp323VveeOU-UtWggvCchDGvN5kYLOJB3kThKj18PpRFL6oqJ8SMQ713LlelqUJwAqSiy0d05u0_YSZZ0Ue2O_46TRCnGbKiotbkXCTIVhk6JduEWZLnC3Yakpdcdg3cbKNk5eGK7MNGGsxK3YV2QKwVmzBsK06oU1xlBJv3X_oDtye7pV_BPv0Yx54a5PV1kkys24NQ6oREPpp8n2q6uXOa4PvV9Ra9q1Ejr8V2TCJ7J57_37xs8TU7Zl5ZF56AFD4qtksUwjbWTcZEs3SVF2cOQrxhM4_60dsCgKKM5nNzYZ5-GnM`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err!=nil{
		fmt.Println("Unable to Create a new http request")
		return
	}else{
		req.Header.Set("Cookie", cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
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