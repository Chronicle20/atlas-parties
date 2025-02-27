package character

const (
	EnvEventMemberStatusTopic        = "EVENT_TOPIC_PARTY_MEMBER_STATUS"
	EventPartyMemberStatusTypeLogin  = "LOGIN"
	EventPartyMemberStatusTypeLogout = "LOGOUT"
)

type memberStatusEvent[E any] struct {
	WorldId     byte   `json:"worldId"`
	PartyId     uint32 `json:"partyId"`
	CharacterId uint32 `json:"characterId"`
	Type        string `json:"type"`
	Body        E      `json:"body"`
}

type memberLoginEventBody struct {
}

type memberLogoutEventBody struct {
}
