package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetAvatarDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	for _, a := range g.Player.DbAvatar.Avatar {
		avatarList := new(proto.Avatar)
		avatarList.FirstMetTimestamp = a.FirstMetTimestamp
		avatarList.BaseAvatarId = a.AvatarId
		avatarList.Promotion = a.Promotion
		avatarList.Rank = a.Rank
		avatarList.Level = a.Level
		avatarList.Exp = a.Exp
		if a.AvatarId/100 == 80 {

		} else {
			avatarList.SkilltreeList = GetKilltreeList(a.AvatarId, 1)
		}
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *Game) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)

	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Rank++
	g.SubtractMaterial(req.BaseAvatarId+10000, 1)
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	rsp := new(proto.GetChallengeScRsp)
	g.send(cmd.RankUpAvatarScRsp, rsp)
}

func (g *Game) DressAvatarCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.DressAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressAvatarCsReq)

	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].BaseAvatarId = req.BaseAvatarId
	g.Player.DbAvatar.Avatar[req.BaseAvatarId].EquipmentUniqueId = req.EquipmentUniqueId

	g.DressAvatarPlayerSyncScNotify(req.BaseAvatarId, req.EquipmentUniqueId)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.DressAvatarScRsp, rsp)
}

func (g *Game) DressAvatarPlayerSyncScNotify(avatarId, equipmentUniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync:    &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		EquipmentList: make([]*proto.Equipment, 0),
	}
	avatardb := g.Player.DbAvatar.Avatar[avatarId]
	avatar := &proto.Avatar{
		SkilltreeList:     GetKilltreeList(avatarId, 1),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      make([]uint32, 0),
		FirstMetTimestamp: avatardb.FirstMetTimestamp,
		Promotion:         avatardb.Promotion,
		Level:             avatardb.Level,
	}
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	equipment := g.Player.DbItem.EquipmentMap[equipmentUniqueId]

	equipmentList := &proto.Equipment{
		Exp:          equipment.Exp,
		Promotion:    equipment.Promotion,
		Level:        equipment.Level,
		BaseAvatarId: equipment.BaseAvatarId,
		IsProtected:  equipment.IsProtected,
		Rank:         equipment.Rank,
		UniqueId:     equipment.UniqueId,
		Tid:          equipment.Tid,
	}

	notify.EquipmentList = append(notify.EquipmentList, equipmentList)

	g.send(cmd.PlayerSyncScNotify, notify)

	g.UpDataPlayer()
}
