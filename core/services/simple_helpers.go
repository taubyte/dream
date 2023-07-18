package services

import "github.com/taubyte/dreamland/core/common"

func (s *Simple) getAll() map[string]common.ClientCreationMethod {
	return map[string]common.ClientCreationMethod{
		"auth":    s.CreateAuthClient,
		"hoarder": s.CreateHoarderClient,
		"billing": s.CreateBillingClient,
		"monkey":  s.CreateMonkeyClient,
		"patrick": s.CreatePatrickClient,
		"seer":    s.CreateSeerClient,
		"tns":     s.CreateTNSClient,
		"console": s.CreateConsoleClient,
	}
}
