package redirect

import (
	"fmt"
	"net/http"
)

func CheckRedirect(url string) (string,error){
	if url == "" {
		return "",nil
	}
	client := &http.Client{
		CheckRedirect: CustomCheckRedirect,
	}
	request, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(request)
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 301 || resp.StatusCode == 302 {
		respUrl,err := resp.Location()
		if err != nil {
			return "",err
		}
		url = respUrl.String()
	}
	return url,nil
}

func CustomCheckRedirect(req *http.Request, via []*http.Request)  error {
	//自用，将url根据需求进行组合
	return http.ErrUseLastResponse
}
