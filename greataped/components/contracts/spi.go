package contracts

import . "github.com/xeronith/diamante/contracts/security"

var SpiPassThroughFilter = func(ISpi) bool { return true }

type (
	Spis               []ISpi
	SpiIterator        func(ISpi)
	SpiCondition       func(ISpi) bool
	SpiFilterPredicate func(ISpi) bool
	SpiMapPredicate    func(ISpi) ISpi
	SpiCacheCallback   func()

	ISpi interface {
	}

	ISpiCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ISpi
		Append(spi ISpi)
		ForEach(SpiIterator)
		Array() Spis
	}

	ISpiManager interface {
		ISystemComponent
		OnCacheChanged(SpiCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition SpiCondition) bool
		ListSpis(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISpiCollection
		GetSpi(id int64, editor Identity) (ISpi, error)
		AddSpi(editor Identity) (ISpi, error)
		AddSpiWithCustomId(id int64, editor Identity) (ISpi, error)
		AddSpiObject(spi ISpi, editor Identity) (ISpi, error)
		AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error)
		AddSpiWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ISpi, error)
		AddSpiObjectAtomic(transaction ITransaction, spi ISpi, editor Identity) (ISpi, error)
		Log(source string, editor Identity, payload string)
		UpdateSpi(id int64, editor Identity) (ISpi, error)
		UpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error)
		UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		UpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error)
		AddOrUpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error)
		AddOrUpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error)
		RemoveSpi(id int64, editor Identity) (ISpi, error)
		RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		Find(id int64) ISpi
		ForEach(iterator SpiIterator)
		Filter(predicate SpiFilterPredicate) ISpiCollection
		Map(predicate SpiMapPredicate) ISpiCollection
		Echo(document IDocument, editor Identity) (IEchoResult, error)
		Signup(username string, email string, password string, editor Identity) (ISignupResult, error)
		Verify(email string, token string, code string, editor Identity) (IVerifyResult, error)
		Login(email string, password string, editor Identity) (ILoginResult, error)
		GetProfileByUser(editor Identity) (IGetProfileByUserResult, error)
		UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string, editor Identity) (IUpdateProfileByUserResult, error)
		Logout(editor Identity) (ILogoutResult, error)
		Webfinger(resource string, editor Identity) (IWebfingerResult, error)
	}

	IEchoResult interface {
		Document() IDocument
	}

	ISignupResult interface {
		Token() string
		Code() string
	}

	IVerifyResult interface {
		Token() string
	}

	ILoginResult interface {
		Username() string
		Token() string
	}

	IGetProfileByUserResult interface {
		Username() string
		DisplayName() string
		Avatar() string
		Banner() string
		Summary() string
		Github() string
	}

	IUpdateProfileByUserResult interface {
		DisplayName() string
		Avatar() string
		Banner() string
		Summary() string
		Github() string
	}

	ILogoutResult interface {
	}

	IWebfingerResult interface {
		Aliases() []string
		Links() []IActivityPubLink
		Subject() string
	}
)
