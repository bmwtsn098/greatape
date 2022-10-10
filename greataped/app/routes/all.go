package routes

import "contracts"

var All = []contracts.IRoute{
	Root,
	Health,
	Error,
	Upload,
	Profile,
	Signup,
	Verify,
	Login,
	GetProfile,
	UpdateProfile,
	WebFinger,
	User,
	Message,
	InboxPost,
	InboxGet,
	OutboxPost,
	OutboxGet,
	Followers,
	Following,
	Follow,
	Feed,
	FeedTypes,
	Authorize,
}
