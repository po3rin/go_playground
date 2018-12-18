package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Groups []Group

type Group struct {
	Names []string `json:"name"`
}

func (g *Group) MarshalJSON() ([]byte, error) {
	var alias Group
	fmt.Println(g.Names)
	if g.Names == nil {
		alias.Names = []string{}
	} else {
		alias.Names = g.Names
	}
	return json.Marshal(alias)
}

func Hello(c *gin.Context) {
	group := getGroup()
	c.JSON(200, group)
}

func getGroup() Groups {
	return []Group{
		{
			Names: []string{"po3rin", "gopher"},
		},
		{
			Names: nil,
		},
		{
			Names: []string{},
		},
	}
}
