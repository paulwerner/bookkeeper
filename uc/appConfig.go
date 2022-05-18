package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (i interactor) AppConfig() (*d.AppConfig, error) {
	return d.Config(), nil
}
