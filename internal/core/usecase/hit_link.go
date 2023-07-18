package usecase

import (
	"context"
	"time"

	"gotiny/internal/core/model"
	"gotiny/internal/core/port"
)

type HitLink struct {
	repository        port.LinksRepository
	ipRepository      port.IpRepository
	ipCacheRepository port.IpCacheRepository
}

func NewHitLink(
	repository port.LinksRepository,
	ipRepository port.IpRepository,
	ipCacheRepository port.IpCacheRepository,
) *HitLink {
	return &HitLink{
		repository:        repository,
		ipRepository:      ipRepository,
		ipCacheRepository: ipCacheRepository,
	}
}

func (u *HitLink) Call(
	ctx context.Context,
	id string,
	requestData model.RedirecsRequestData,
) (*string, error) {
	link, err := u.repository.GetLinkById(ctx, id)
	if err != nil {
		return nil, &model.AppError{
			Type:    string(model.UnknownError),
			Context: err,
		}
	}

	if link == nil {
		return nil, nil
	}

	go u.saveAnalitics(link, requestData)

	if !link.Valid() {
		go u.repository.DeleteLinkById(context.Background(), id)
		return nil, nil
	}

	go u.repository.IncrementHitCount(context.Background(), id)

	return &link.OriginalLink, nil
}

func (u *HitLink) saveAnalitics(link *model.Link, requestData model.RedirecsRequestData) {
	if link.TrackUntil == nil {
		return
	}

	if link.TrackUntil.Before(time.Now()) {
		return
	}
	ipDetails := u.getIpDetails(requestData.Ip)

	u.repository.SaveHitAnalitics(context.Background(), link.Id, model.LinkHitAnalitics{
		IpDetails:   ipDetails,
		RequestData: requestData,
	})
}

func (u *HitLink) getIpDetails(ip string) *model.IpDetails {
	ipDetails, err := u.ipCacheRepository.GetIpDetails(context.Background(), ip)
	if err != nil || ipDetails.Ip == "" {
		return nil
	}

	if ipDetails != nil {
		return ipDetails
	}

	newDetails, err := u.ipRepository.GetIpDetails(context.Background(), ip)
	if err != nil {
		return nil
	}

	go u.ipCacheRepository.SaveIpDetails(context.Background(), newDetails)

	return &newDetails
}
