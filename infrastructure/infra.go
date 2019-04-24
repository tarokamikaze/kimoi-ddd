package infrastructure

import (
	"fmt"
	"github.com/tarokamikaze/kimoi-ddd/domain"
	"reflect"
)

type (
	ActiveRecordUser struct {
		ID       string
		MailAddr string
	}
	UserRepositoryImpl struct {
	}
)

func fromEntity(e domain.EntityUser) ActiveRecordUser {
	rt := reflect.ValueOf(e)
	return ActiveRecordUser{
		ID:       rt.FieldByName("id").String(),
		MailAddr: rt.FieldByName("mailAddr").String(),
	}
}
func (r *UserRepositoryImpl) Save(e domain.EntityUser) error {
	ar := fromEntity(e)
	// save
	fmt.Printf("ActiveRecord ID: %s MailAddr: %s\n", ar.ID, ar.MailAddr)
	return nil
}
