package main

import (
	"a4m.dev/godev/dsa"
)

func main() {
	list := dsa.New[int64](true)
	list.Append(450)
	list.Append(340)
	list.Append(500)
	list.Append(-120)
	list.Append(90)
	list.Append(42)
	list.Print()
	// res, err := github.SearchIssues([]string{"C++", "python", "golang"})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, item := range res.Items {
	// 	fmt.Printf("#%-5d %s (%s - %s) - %s\n",
	// 		item.Number, item.Title, item.User.Login, item.User.HtmlUrl, item.HtmlUrl)
	// }
}
