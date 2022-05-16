package uc

import d "github.com/paulwerner/bookkeeper/domain"

func (self interactor) AppConfig() (*d.AppConfig, error) {
	return d.Config(), nil
}
