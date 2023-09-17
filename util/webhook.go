package util

import (
	"bytes"
	"fmt"
	"net/http"
)

type Post struct {
	Content string `json:"content"`
}

func SendWebhook() (err error) {
	fmt.Println("Hello World")
	postURL := "https://discord.com/api/webhooks/1105303399324463164/51VehhrrvlGYLTBB_mozEx6_CRGzb4d-CBERYp-4jb2_eN2qRGNNzauRrZPvgo3OwNmg"

	body := []byte(`{
		"content": "",
		"username": "BIGWILLY",
		"avatar_url": "https://styles.redditmedia.com/t5_2sqho/styles/communityIcon_kdkwphdj62eb1.png?width=256&s=ca237ed1392dcddd57ec42412f307d69c78b3b42",
		"embeds": [{
		  "title": "This is a Test Embed",
		  "description": "This is a random description test test test test test",
		  "url": "https://discord.com/developers/docs/resources/channel#embed-object-embed-structure",
		  "color": 49151,
		  "thumbnail": {
			"url": "https://styles.redditmedia.com/t5_2sqho/styles/communityIcon_kdkwphdj62eb1.png?width=256&s=ca237ed1392dcddd57ec42412f307d69c78b3b42"
		  },
		  "timestamp": "2015-12-31T12:00:00.000Z",
		  "author": {
			"name": "ANDREWTESTAUTHOR"
		  },
		  "footer": {
			"text": "this is the footer"
		  },
		  "fields": [
			{
			"name": "This is a test field",
			"value": "testFieldValue",
			"inline": true
			}, 
			{
			"name": "anotehr test field",
			"value": "some bs here",
			"inline": true
			}
		  ]
		}]
	  }`)

	r, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	return nil
}
