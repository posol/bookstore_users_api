package services

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemService struct {
}

func (i *itemService) GetItem() {

}

func (i itemService) SaveItem() {

}
