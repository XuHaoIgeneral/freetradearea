package service

import (
	"apiserver/model"
	"sync"
	"apiserver/util"
	"fmt"
)

func ListUsers(username string,offset,limit int) ([]*model.UserInfo,uint64,error){
	infos:=make([]*model.UserInfo,0)
	users,count,err:=model.ListUser(username,offset,limit)
	if err!=nil {
		return nil,count,err
	}

	ids:=[]uint64{}
	for _,user:=range users{
		ids=append(ids,user.Id)
	}

	wg:=sync.WaitGroup{}

	userList:=model.UserList{
		Lock:new(sync.Mutex),
		IdMap:make(map[uint64]*model.UserInfo,len(users)),
	}

	errChan:=make(chan error,1)
	finished:=make(chan bool,1)

	//使用go程写入  花里胡哨的，还不是要加锁，还去开了新的go程，炫技？
	for _,u:=range users{
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			shortId,err:=util.GenShortId()
			if err!=nil {
				errChan<-err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id]=&model.UserInfo{
				Id:u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	//GO程监听事态
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err:=<-errChan:
		return nil, count, err
	}
	for _,id:=range ids{
		infos=append(infos,userList.IdMap[id])
	}
	return infos, count, nil
}