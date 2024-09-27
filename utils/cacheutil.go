package utils

import (
	"pigeon/entity"

	"github.com/robfig/cron"
)

var jokesCache []entity.Joke // 笑话缓存

func init() {
	c := cron.New()
	err := c.AddFunc("@every 30m", cacheJokesFromDB)
	if err != nil {
		panic(err)
	}
	c.Start()
}

// 查询数据库并缓存数据到内存的函数
func cacheJokesFromDB() {
	// 这里编写查询数据库获取笑话并将结果存储到 jokes 切片的代码
	// 示例：假设你有一个函数 getJokesFromDB 可以从数据库获取笑话
	// jokes = getJokesFromDB()
}
