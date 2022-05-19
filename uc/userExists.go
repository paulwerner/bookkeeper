package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) UserExists(id d.UserID) (ok bool) {
	ok, err := i.userRW.Exists(id)
	if err != nil {
		return false
	}
	return
}
