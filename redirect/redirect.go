package redirect

import (
	"errors"
	"net/http"
)

func CheckRedirect(url string) (string,error){
	if url == "" {
		return "",nil
	}
	client := &http.Client{
		CheckRedirect: CustomCheckRedirect,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return url,err
	}
	resp, err := client.Do(request)
	if err != nil {
		return url,err
	}
	if resp.StatusCode == 301 || resp.StatusCode == 302 {
		respUrl,err := resp.Location()
		if err != nil {
			return url,err
		}
		url = respUrl.String()
	}
	return url,nil
}

func CustomCheckRedirect(req *http.Request, via []*http.Request)  error {
	//自用，将url根据需求进行组合
	if len(via) >= 1 {
		return errors.New("stopped after 1 redirects")
	}
	return nil
}
