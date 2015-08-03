package system

import (
	"github.com/emicklei/go-restful"
	"github.com/garyburd/redigo/redis"
	"github.com/pelletier/go-toml"
	"gopkg.in/gorp.v1"
)

type Controller struct {
}

func (c *Controller) GetConfig(req *restful.Request) *toml.TomlTree {
	tmp := req.Attribute("app.config")
	if tmp != nil {
		val := tmp.(*toml.TomlTree)
		return val
	}
	return nil
}

func (c *Controller) GetSQL(req *restful.Request) *gorp.DbMap {
	tmp := req.Attribute("sql")
	if tmp != nil {
		val := tmp.(*gorp.DbMap)
		return val
	}
	return nil
}

func (c *Controller) GetRedis(req *restful.Request) *redis.Pool {
	tmp := req.Attribute("redis")
	if tmp != nil {
		val := tmp.(*redis.Pool)
		return val
	}
	return nil
}

func (c *Controller) GetQueries(req *restful.Request) *QueryManager {
	tmp := req.Attribute("queries")
	if tmp != nil {
		val := tmp.(*QueryManager)
		return val
	}
	return nil
}
