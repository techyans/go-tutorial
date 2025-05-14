package main

import "fmt"

func maps() {
	website := map[string]string{
		"google":   "www.google.com",
		"facebook": "www.facebook.com",
		"twitter":  "www.twitter.com",
	}

	fmt.Println(website)
	fmt.Println(website["google"])

	website["linkedin"] = "www.linkedin.com"
	fmt.Println(website)

	delete(website, "google")
	fmt.Println(website)
}
