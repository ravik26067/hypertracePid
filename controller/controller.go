package controller

import (
	"bytes"
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
	//startServer()
	makeRequest()
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

	requestBody = `[
  {
    "variables": {},
    "query": "{\n  entities(\n    scope: \"API\"\n    limit: 100\n    between: {startTime: \"2021-10-07T10:01:47.055Z\", endTime: \"2021-10-07T10:16:47.055Z\"}\n    filterBy: [{operator: EQUALS, value: \"e2ef8e5d-1cdb-3ff5-ab3e-203279920399\", type: ID, idType: API}]\n  ) {\n    results {\n      id\n      name: attribute(key: \"name\")\n      duration: metric(key: \"duration\") {\n        p99: percentile(size: 99) {\n          value\n          __typename\n        }\n        p50: percentile(size: 50) {\n          value\n          __typename\n        }\n        __typename\n      }\n      errorCount: metric(key: \"errorCount\") {\n        avg {\n          value\n          __typename\n        }\n        sum {\n          value\n          __typename\n        }\n        avgrate_sec: avgrate(units: SECONDS, size: 1) {\n          value\n          __typename\n        }\n        __typename\n      }\n      numCalls: metric(key: \"numCalls\") {\n        avg {\n          value\n          __typename\n        }\n        sum {\n          value\n          __typename\n        }\n        avgrate_sec: avgrate(units: SECONDS, size: 1) {\n          value\n          __typename\n        }\n        __typename\n      }\n      outgoingEdges_API: outgoingEdges(neighborType: API) {\n        results {\n          duration: metric(key: \"duration\") {\n            p99: percentile(size: 99) {\n              value\n              __typename\n            }\n            p50: percentile(size: 50) {\n              value\n              __typename\n            }\n            __typename\n          }\n          errorCount: metric(key: \"errorCount\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          numCalls: metric(key: \"numCalls\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          neighbor {\n            id\n            name: attribute(key: \"name\")\n            duration: metric(key: \"duration\") {\n              p99: percentile(size: 99) {\n                value\n                __typename\n              }\n              p50: percentile(size: 50) {\n                value\n                __typename\n              }\n              __typename\n            }\n            errorCount: metric(key: \"errorCount\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            numCalls: metric(key: \"numCalls\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      outgoingEdges_BACKEND: outgoingEdges(neighborType: BACKEND) {\n        results {\n          duration: metric(key: \"duration\") {\n            p99: percentile(size: 99) {\n              value\n              __typename\n            }\n            p50: percentile(size: 50) {\n              value\n              __typename\n            }\n            __typename\n          }\n          errorCount: metric(key: \"errorCount\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          numCalls: metric(key: \"numCalls\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          neighbor {\n            id\n            name: attribute(key: \"name\")\n            type: attribute(key: \"type\")\n            duration: metric(key: \"duration\") {\n              p99: percentile(size: 99) {\n                value\n                __typename\n              }\n              p50: percentile(size: 50) {\n                value\n                __typename\n              }\n              __typename\n            }\n            errorCount: metric(key: \"errorCount\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            numCalls: metric(key: \"numCalls\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      incomingEdges_SERVICE: incomingEdges(neighborType: SERVICE) {\n        results {\n          duration: metric(key: \"duration\") {\n            p99: percentile(size: 99) {\n              value\n              __typename\n            }\n            p50: percentile(size: 50) {\n              value\n              __typename\n            }\n            __typename\n          }\n          errorCount: metric(key: \"errorCount\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          numCalls: metric(key: \"numCalls\") {\n            avg {\n              value\n              __typename\n            }\n            sum {\n              value\n              __typename\n            }\n            avgrate_sec: avgrate(units: SECONDS, size: 1) {\n              value\n              __typename\n            }\n            __typename\n          }\n          neighbor {\n            id\n            name: attribute(key: \"name\")\n            duration: metric(key: \"duration\") {\n              p99: percentile(size: 99) {\n                value\n                __typename\n              }\n              p50: percentile(size: 50) {\n                value\n                __typename\n              }\n              __typename\n            }\n            errorCount: metric(key: \"errorCount\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            numCalls: metric(key: \"numCalls\") {\n              avg {\n                value\n                __typename\n              }\n              sum {\n                value\n                __typename\n              }\n              avgrate_sec: avgrate(units: SECONDS, size: 1) {\n                value\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"
  },
  {
    "variables": {},
    "query": "{\n  entities(\n    scope: \"API\"\n    limit: 1\n    between: {startTime: \"2021-10-07T10:01:47.055Z\", endTime: \"2021-10-07T10:16:47.055Z\"}\n    filterBy: [{operator: EQUALS, value: \"e2ef8e5d-1cdb-3ff5-ab3e-203279920399\", type: ID, idType: API}]\n    includeInactive: true\n  ) {\n    results {\n      id\n      name: attribute(key: \"name\")\n      __typename\n    }\n    __typename\n  }\n}\n"
  }
]`
	if err!=nil{
		fmt.Println("Unable To Marshal")
		return
	}

	cookie := `_gaexp=GAX1.2.3mdAZUKqSLClB8mBuzttMQ.18983.1; _gcl_au=1.1.900131272.1633334581; firstAttribUtm={"utm_source":"","utm_campaign":"","utm_medium":"","utm_term":"","utm_content":"","utm_adgroup":"","timestamp":"2021-10-04T13:33:01%2B05:30"}; lastAttribUtm={"utm_source":"direct","utm_medium":"website","utm_campaign":"direct","utm_content":"not set","utm_keyword":"not set"}; rzp_utm={"attributions":[],"first_page":"admin-dashboard.razorpay.com/signin","new_user":true}; clientId=d08ce78b-a010-4b2c-a8d7-a7792ff2bb2d; _rdt_uuid=1633334581845.40ff8dc3-47d7-4b28-8270-3fc009c45048; _ga=GA1.2.809373854.1633334582; _gid=GA1.2.10977730.1633334582; ajs_anonymous_id=%228e699ec7-79e6-446a-804a-927ce8636802%22; _fbp=fb.1.1633334589056.1747874148; ajs_user_id=%22HsbBfFzeXZ1qyX%22; WZRK_G=253a7b4d7f4e4589a72244c82dea0c16; campaignStartTime=Mon, 04 Oct 2021 13:45:57 GMT; _uetsid=823724a024e911eca5f38fa1fc085271; _uetvid=8237ae1024e911ec9f74e92cb94abe38; _sso_proxy=aRXUlwLVBLf7SyKuf_oEWL341jCVNCqsNVRELo4C_q5tE9UivFqnu3E258xKCvUHLgJMPIAAmHWWRXCrJAUGZlz1Y5eJfdJrMxwqpZLPYvdgwZYwUKaF4jX9NyayC-xCLStINfSzC5sop7Qig1iQC49nJJLStdyDK3y8-6-JrqmdxCoTET0r8u6cjxzQLf_6FXypOJ8yzAbHwlf_nz72RK1CJg-g_aMx1cijhObR2pAt9M-6kidhNsrYtu3YB3KNvxnndNEft7g5_U6qkStSpLbH7mr5zXGuqJzqYAs5y4W9ancn_XewFkVrWrobruqQb31s7Zdw3KMDygxWg_zCahgXw14249u1MXMpf8OoNQt7dHvrnyScxxh-KAOMRCDe881EFw2tU4pW-z7-LZdHwHk__PbFTqyRFsaNG_dHu1FaxfnsAopQKjExVkWD6QGnEmmLsuWDkHx4qwxxjCXRB-fd9j5e5pjJa1sXD5gkC50-LfXA6xE1cLbLvHeGQ4M0w7w4hThdCcUfjDmedtv5WqJvsaFd46Il5Muk-_IfZC5YDY-4DWWe4clrBHDsQ8AhP5NitNVoNyfA0S1ec0l0szGZjlq60upSdRiYUTv1MT6xbLsIFNAyeGB46MmovGLD3cJEHDkvNgtrKGL_ZMTtHA5RcrkZcQgt4OuXtDFPmUgQwd_7SMjtPLYBBWKg-If0Mtd408LPa7itgKAvUHs9sASNnDanTS8jj08mUR66ezpcpxY4tYcVNWUwKM8jVg-sPhYxmHNRqQL7H_vhjficcGn8xMvi3ytD`
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