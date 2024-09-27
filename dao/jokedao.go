package dao

import (
	"log"
	"pigeon/entity"
	"pigeon/utils"
)

func GetJokeByPage(deviceId string, page int, pageSize int) ([]entity.Joke, error) {
	// query := `select t.joke_id, t.content from(
	// select joke_id, content, row_number() over(order by joke_id) as rn from dim_joke_di) as t
	// where t.rn between ? and ?`
	query := `select t.joke_id, t.content from(
	select t1.joke_id, t1.content, row_number() over(order by t2.like_rate desc) as rn
	from(select joke_id, content from dim_joke_di) as t1
	left join(
		select joke_id, if(sum(case when act_type='show' then 1 else 0 end)=0, 0
		,sum(case when act_type='like' then 1 else 0 end)/sum(case when act_type='show' then 1 else 0 end)) as like_rate
		from dwd_joke_act_rt group by joke_id
		) as t2
	on t1.joke_id = t2.joke_id
	left join(select joke_id from dwd_joke_act_rt where device_id = ? group by joke_id) as t3
	on t1.joke_id = t3.joke_id where t3.joke_id is null) as t
	where t.rn between ? and ?`
	rows, err := utils.Select(query, deviceId, (page-1)*pageSize+1, page*pageSize)
	if err != nil {
		log.Fatalf("Select error: %v", err)
	}
	defer rows.Close()

	jokes := make([]entity.Joke, 0)
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

func InsertUserAct(deviceId string, jokeId string, actType string) error {
	query := "insert into dwd_joke_act_rt values(?, ?, ?, ?);"
	actTime, _ := utils.StringToTime(utils.GetNowTime())
	_, err := utils.Insert(query, deviceId, jokeId, actType, actTime)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}

	return err
}

func GetUserShowHistory(deviceId string) ([]string, error) {
	query := "select joke_id from dwd_joke_act_rt where device_id = ? group by joke_id"
	rows, err := utils.Select(query, deviceId)
	if err != nil {
		log.Fatalf("Select error: %v", err)
	}
	defer rows.Close()

	jokes := make([]string, 0)
	for rows.Next() {
		var jokeId string
		err := rows.Scan(&jokeId)
		if err != nil {
			return nil, err
		}
		jokes = append(jokes, jokeId)
	}

	return jokes, nil
}
