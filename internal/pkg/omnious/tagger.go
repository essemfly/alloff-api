package omnious

import (
	"encoding/json"
	"fmt"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/utils"
	"io/ioutil"
)

func GetOmniousData(imgUrl string) {
	omniousKey := config.OmniousKey
	url := "https://api.omnious.com/tagger/v2.12/tags"
	method := utils.REQUEST_POST
	header := utils.GetOmniousHeader(omniousKey)
	body := fmt.Sprintf(`
		{
		  "image": {
			"type": "url",
			"content": "%s"
		  }
		}
	`, imgUrl)

	resp, err := utils.MakeRequest(url, method, header, body)
	if err != nil {
		fmt.Println("TODO: 에러처리1 : ", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("TODO 에러처리2 : ", err)
	}

	var result PostResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("TODO 에러처리3 : ", err)
	}

	fmt.Println(result)
}
