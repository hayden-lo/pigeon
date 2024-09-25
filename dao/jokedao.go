package dao

import (
	"log"
	"pigeon/entity"
	"pigeon/utils"
	"time"
)

func GetJokeByPage(page int, pageSize int) ([]entity.Joke, error) {
	query := `select t.joke_id, t.content from(
	select joke_id, content, row_number() over(order by joke_id) as rn from dim_joke_di) as t
	where t.rn between ? and ?`
	rows, err := utils.Select(query, page*pageSize, (page+1)*pageSize)
	if err != nil {
		log.Fatalf("Select error: %v", err)
	}
	defer rows.Close()

	var jokes []entity.Joke
	for rows.Next() {
		var joke entity.Joke
		err := rows.Scan(&joke.ID, &joke.Content)
		if err != nil {
			return nil, err
		}
		jokes = append(jokes, joke)
	}

	return jokes, nil
}

func InsertUserAct(jokeId string, actType string) error {
	query := "insert into dwd_joke_act_rt values('?','?', ?);"
	actTime := time.Now().Unix()
	_, err := utils.Insert(query, jokeId, actType, actTime)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}

	return err
}
