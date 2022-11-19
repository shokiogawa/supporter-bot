package query_service_interface

type CommonQueryService interface {
	GetPublicUserId(lineUserId string) (publicUserId string, err error)
}
