package service

import (
	"encoding/json"
	"fmt"
	"github.com/xiaoweihong/wolfweb/model/response"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

func GetRepos(url string) (result response.FseListResponseResult, err error) {
	resp, err := client.Get(fmt.Sprintf("http://%s/x-api/v1/repositories", url))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(all, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteRepoById(url, id string) (result response.FseResponseResult, err error) {
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("http://%s/x-api/v1/repositories/%s", url, id), nil)
	resp, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(all, &result)
	if err != nil {
		return result, err
	}

	return result, nil

}
