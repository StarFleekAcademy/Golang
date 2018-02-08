package main

import "fmt"

func main() {
	var userName string = "cthulhuplus"
	var endPoint string = fmt.Sprintf("%s%s%s", "users/", userName, "/repos")
	var sourceUrl string = "https://api.github.com/"
	var GITHUB_API_OATH_KEY = "?access_token=YOUSHALLPASS"

	url := fmt.Sprintf("%s%s%s", sourceUrl, endPoint, GITHUB_API_OATH_KEY)

	fmt.Println(url)
}
