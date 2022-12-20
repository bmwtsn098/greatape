package contracts

import . "github.com/xeronith/diamante/contracts/server"

// noinspection GoSnakeCaseUsage
const (
	//SystemCallOperation
	SYSTEM_CALL_REQUEST = 0x00001000
	SYSTEM_CALL_RESULT  = 0xF0001000

	//EchoOperation
	ECHO_REQUEST = 0x0541BD72
	ECHO_RESULT  = 0xAB2FF7D4

	//SignupOperation
	SIGNUP_REQUEST = 0x48DB23BF
	SIGNUP_RESULT  = 0x83D062B4

	//VerifyOperation
	VERIFY_REQUEST = 0x8B78F7F6
	VERIFY_RESULT  = 0x2C8A8A49

	//LoginOperation
	LOGIN_REQUEST = 0xF480F151
	LOGIN_RESULT  = 0xBE819605

	//GetProfileByUserOperation
	GET_PROFILE_BY_USER_REQUEST = 0xEAB16E71
	GET_PROFILE_BY_USER_RESULT  = 0x8EECDE97

	//UpdateProfileByUserOperation
	UPDATE_PROFILE_BY_USER_REQUEST = 0xC25AB0BA
	UPDATE_PROFILE_BY_USER_RESULT  = 0x678A8BAF

	//LogoutOperation
	LOGOUT_REQUEST = 0x447AFA34
	LOGOUT_RESULT  = 0x9412D17F

	//WebfingerOperation
	WEBFINGER_REQUEST = 0x01FD357C
	WEBFINGER_RESULT  = 0xCC81EC52

	//GetActorOperation
	GET_ACTOR_REQUEST = 0x5C4AC410
	GET_ACTOR_RESULT  = 0x136B82A8

	//FollowActorOperation
	FOLLOW_ACTOR_REQUEST = 0xD30C2420
	FOLLOW_ACTOR_RESULT  = 0x30154D74

	//AuthorizeInteractionOperation
	AUTHORIZE_INTERACTION_REQUEST = 0x59EA7612
	AUTHORIZE_INTERACTION_RESULT  = 0xB38E936F

	//GetFollowersOperation
	GET_FOLLOWERS_REQUEST = 0x3F20FD65
	GET_FOLLOWERS_RESULT  = 0x7F3E2EB5

	//GetFollowingOperation
	GET_FOLLOWING_REQUEST = 0xF9841DB9
	GET_FOLLOWING_RESULT  = 0xD707408F

	//PostToOutboxOperation
	POST_TO_OUTBOX_REQUEST = 0x9E489553
	POST_TO_OUTBOX_RESULT  = 0xC6C56614

	//GetOutboxOperation
	GET_OUTBOX_REQUEST = 0x527B6997
	GET_OUTBOX_RESULT  = 0xF94E37A0
)

var OPCODES = Opcodes{
	0x00000000: "N/A",
	0x0541BD72: "ECHO",
	0xAB2FF7D4: "Echo",
	0x48DB23BF: "SIGNUP",
	0x83D062B4: "Signup",
	0x8B78F7F6: "VERIFY",
	0x2C8A8A49: "Verify",
	0xF480F151: "LOGIN",
	0xBE819605: "Login",
	0xEAB16E71: "GET_PROFILE_BY_USER",
	0x8EECDE97: "GetProfileByUser",
	0xC25AB0BA: "UPDATE_PROFILE_BY_USER",
	0x678A8BAF: "UpdateProfileByUser",
	0x447AFA34: "LOGOUT",
	0x9412D17F: "Logout",
	0x01FD357C: "WEBFINGER",
	0xCC81EC52: "Webfinger",
	0x5C4AC410: "GET_ACTOR",
	0x136B82A8: "GetActor",
	0xD30C2420: "FOLLOW_ACTOR",
	0x30154D74: "FollowActor",
	0x59EA7612: "AUTHORIZE_INTERACTION",
	0xB38E936F: "AuthorizeInteraction",
	0x3F20FD65: "GET_FOLLOWERS",
	0x7F3E2EB5: "GetFollowers",
	0xF9841DB9: "GET_FOLLOWING",
	0xD707408F: "GetFollowing",
	0x9E489553: "POST_TO_OUTBOX",
	0xC6C56614: "PostToOutbox",
	0x527B6997: "GET_OUTBOX",
	0xF94E37A0: "GetOutbox",
}
