package pkg

type Service struct{
	FetchNotices() (*[]presenter.Notice, error)
	InsertNotice(notice *entities.Notice) (*entities.Notice, error)
	DeleteNotice(ID string) error
	UpdateNotice(notice *entities.Notice) (*entities.Notice, error)
	FetchWeeklies() (*[]presenter.Weekly, error)
}


type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}