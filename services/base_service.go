package services

type BaseDomainService interface {
	Get(foo any)
	Update()
	Create()
}
