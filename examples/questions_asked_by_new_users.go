package main

import "fmt"
import "time"
import "github.com/laktek/Stack-on-Go/stackongo"

func main() {

	fmt.Printf("Questions asked by users registered in last 24 hours:\n")

	from_date := time.Now().Unix() - (60 * 60 * 24)

	session := stackongo.NewSession("stackoverflow")

	//set the params
	user_params := make(stackongo.Params)
	user_params.Add("fromdate", from_date)
	user_params.Add("sort", "creation")

	users, err := session.AllUsers(user_params)

	if err != nil {
		fmt.Printf(err.Error())
	}

	var user_ids []int
	for _, user := range users.Items {
		user_ids = append(user_ids, user.User_id)
	}

	question_params := make(stackongo.Params)
	question_params.Add("sort", "creation")

	questions, err2 := session.QuestionsFromUsers(user_ids, question_params)

	if err2 != nil {
		fmt.Printf(err2.Error())
	}

	for _, question := range questions.Items {
		fmt.Printf("%v\n", question.Title)
		fmt.Printf("Asked By: %v on %v\n", question.Owner.Display_name, time.Unix(question.Creation_date, 0).UTC())
		fmt.Printf("Link: %v\n\n", question.Link)

	}
}
