package services

import (
	"fmt"

	"github.com/taubyte/dreamland/core/registry"
	commonIface "github.com/taubyte/go-interfaces/common"
	peer "github.com/taubyte/go-interfaces/p2p/peer"
	billingIface "github.com/taubyte/go-interfaces/services/billing"
)

func (u *Universe) CreateBillingService(config *commonIface.ServiceConfig) (peer.Node, error) {
	var err error

	if registry.Registry.Billing.Service == nil {
		return nil, fmt.Errorf(`service is nil, have you imported _ "bitbucket.org/taubyte/billing/service"`)
	}

	billing, err := registry.Registry.Billing.Service(u.ctx, config)
	if err != nil {
		return nil, err
	}

	_billing, ok := billing.(billingIface.Service)
	if !ok {
		return nil, fmt.Errorf("failed type casting billing into a service")
	}

	u.registerService("billing", _billing)
	u.toClose(_billing)

	return _billing.Node(), nil
}

func (s *Simple) CreateBillingClient(config *commonIface.ClientConfig) error {
	if registry.Registry.Billing.Client == nil {
		return fmt.Errorf(`flient is nil, have you imported _ "bitbucket.org/taubyte/billing/api/p2p"`)
	}

	_billing, err := registry.Registry.Billing.Client(s.Node, config)
	if err != nil {
		return err
	}

	var ok bool
	s.Clients.billing, ok = _billing.(billingIface.Client)
	if !ok {
		return fmt.Errorf("setting billing client failed, not OK")
	}

	return nil

}
