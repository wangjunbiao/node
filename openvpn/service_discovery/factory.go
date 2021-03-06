package service_discovery

import (
	"github.com/mysterium/node/datasize"
	"github.com/mysterium/node/identity"
	"github.com/mysterium/node/money"
	"github.com/mysterium/node/openvpn/service_discovery/dto"
	dto_discovery "github.com/mysterium/node/service_discovery/dto"
	"time"
)

var (
	locationUnknown = dto_discovery.Location{}
)

func NewServiceProposal(
	providerId identity.Identity,
	providerContact dto_discovery.Contact,
) dto_discovery.ServiceProposal {
	return NewServiceProposalWithLocation(providerId, providerContact, locationUnknown)
}

func NewServiceProposalWithLocation(
	identity identity.Identity,
	providerContact dto_discovery.Contact,
	nodeLocation dto_discovery.Location,
) dto_discovery.ServiceProposal {
	return dto_discovery.ServiceProposal{
		Id:          1,
		Format:      "service-proposal/v1",
		ServiceType: "openvpn",
		ServiceDefinition: dto.ServiceDefinition{
			Location:          nodeLocation,
			LocationOriginate: nodeLocation,
			SessionBandwidth:  dto.Bandwidth(10 * datasize.MB),
		},
		PaymentMethodType: dto.PAYMENT_METHOD_PER_TIME,
		PaymentMethod: dto.PaymentMethodPerTime{
			// 15 MYST/month = 0,5 MYST/day = 0,125 MYST/hour
			Price:    money.NewMoney(0.125, money.CURRENCY_MYST),
			Duration: 1 * time.Hour,
		},
		ProviderId:       identity.Address,
		ProviderContacts: []dto_discovery.Contact{providerContact},
	}
}
