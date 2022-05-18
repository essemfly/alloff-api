package omnious

import (
	"encoding/json"
	"fmt"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"github.com/lessbutter/alloff-api/internal/utils"
	"io/ioutil"
)

const (
	PostTaggerURL = "https://api.omnious.com/tagger/v2.12/tags"
)

func GetOmniousData(imgUrl string) (*dto.OmniousResult, error) {
	omniousKey := config.OmniousKey
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

	resp, err := utils.MakeRequest(PostTaggerURL, method, header, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result dto.PostResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	omniousResult, err := MapPostResponseToResult(&result)
	if err != nil {
		return nil, err
	}

	return omniousResult, nil
}
