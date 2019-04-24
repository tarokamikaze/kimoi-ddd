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
		// golangのメジャーORMライブラリは、public propertyしかマッピングしてくれない。
		// 一方、データ整合性を担保するためにはDDDのEntityのproperty を一部 package private にする必要がある
		// つまりActive Record(DTO?) 的な Struct (Class) はインフラ層で定義して、public propertyな状態にする必要がある

		// 一方でDDD的には、ドメイン層からインフラ層への依存はご法度。
		// つまり、インフラ層でEntityのprivate propertyから無理やり値を取得してActive Recordに反映するしかない。
		// つらい
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
