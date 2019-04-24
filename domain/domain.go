package domain

type (
	UserID string
	MailAddr string
	EntityUser struct {
		id       UserID
		mailAddr MailAddr
	}
	UserRepository interface {
		Save(e EntityUser) error
	}
	DomainService struct {
		Repo UserRepository
	}
)

func (ds *DomainService) Kimoi(newAddr MailAddr) error {
	// ここはdomain service想定
	e := NewEntityUser()
	e.mailAddr = newAddr

	return ds.Repo.Save(e)
}
func NewEntityUser() EntityUser {
	return EntityUser{
		id:       "xx1",
		mailAddr: "foo@bar.com",
	}
}
