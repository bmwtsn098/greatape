package core

import . "rail.town/infrastructure/components/contracts"

//region IDispatcher Implementation

func (dispatcher *dispatcher) NewEchoResult(document IDocument) IEchoResult {
	return NewEchoResult(document, nil)
}

func (dispatcher *dispatcher) NewSignupResult(token string, code string) ISignupResult {
	return NewSignupResult(token, code, nil)
}

func (dispatcher *dispatcher) NewVerifyResult(token string) IVerifyResult {
	return NewVerifyResult(token, nil)
}

func (dispatcher *dispatcher) NewLoginResult(username string, token string) ILoginResult {
	return NewLoginResult(username, token, nil)
}

func (dispatcher *dispatcher) NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string) IGetProfileByUserResult {
	return NewGetProfileByUserResult(username, displayName, avatar, banner, summary, github, nil)
}

func (dispatcher *dispatcher) NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string) IUpdateProfileByUserResult {
	return NewUpdateProfileByUserResult(displayName, avatar, banner, summary, github, nil)
}

func (dispatcher *dispatcher) NewLogoutResult() ILogoutResult {
	return NewLogoutResult(nil)
}

func (dispatcher *dispatcher) NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string) IWebfingerResult {
	return NewWebfingerResult(aliases, links, subject, nil)
}

//endregion
