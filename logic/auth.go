package main

import (
    "goim/libs/define"
    "encoding/json"
)

type TmgToken struct {
    Uid int64 `json:"Uid"`
    Rid int32 `json:"Rid"`
}

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
    Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
    return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
    var tmgtkn TmgToken
    err := json.Unmarshal([]byte(token), &tmgtkn)
    if err != nil {
        userId = 0
        roomId = define.NoRoom
    } else {
        userId = tmgtkn.Uid
        roomId = tmgtkn.Rid
    }
    return
}
