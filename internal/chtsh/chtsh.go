package chtsh

import (
	"fmt"
	"io"
	"net/http"
)

const ApiUrl string = "https://cht.sh"

func GetChtsh(tech string, args []string) string {
	formatedArg := ""
	if len(args) > 0 {
		for _, arg := range args {
			formatedArg += arg + "+"
		}
		formatedArg = "/" + formatedArg[:len(formatedArg)-1]
	}

	url := fmt.Sprintf("%s/%s%s", ApiUrl, tech, formatedArg)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "curl/8")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
