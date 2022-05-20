package uc

import d "github.com/paulwerner/bookkeeper/pkg/domain"

func (i interactor) UserExists(id d.UserID) (ok bool) {
	ok, err := i.userStore.Exists(id)
	if err != nil {
		return false
	}
	return
}
