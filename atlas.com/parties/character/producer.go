package character

import (
	"github.com/Chronicle20/atlas-kafka/producer"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/segmentio/kafka-go"
)

func loginEventProvider(partyId uint32, worldId byte, characterId uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(partyId))
	value := &memberStatusEvent[memberLoginEventBody]{
		PartyId:     partyId,
		WorldId:     worldId,
		CharacterId: characterId,
		Type:        EventPartyMemberStatusTypeLogin,
		Body:        memberLoginEventBody{},
	}
	return producer.SingleMessageProvider(key, value)
}

func logoutEventProvider(partyId uint32, worldId byte, characterId uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(partyId))
	value := &memberStatusEvent[memberLogoutEventBody]{
		PartyId:     partyId,
		WorldId:     worldId,
		CharacterId: characterId,
		Type:        EventPartyMemberStatusTypeLogout,
		Body:        memberLogoutEventBody{},
	}
	return producer.SingleMessageProvider(key, value)
}
