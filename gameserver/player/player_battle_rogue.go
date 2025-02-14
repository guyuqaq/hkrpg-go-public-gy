package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

// 区域通知
func (g *GamePlayer) SyncRogueMapRoomScNotify() {
	rogue := g.GetDbRogue()

	notify := &proto.SyncRogueMapRoomScNotify{
		CurRoom: &proto.RogueRoom{
			CurStatus: proto.RogueRoomStatus(rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomStatus),
			SiteId:    rogue.CurRogue.CurSiteId,
			RoomId:    rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId,
		},
		MapId: rogue.CurRogue.RogueMapID,
	}
	g.Send(cmd.SyncRogueMapRoomScNotify, notify)
}

func (g *GamePlayer) SyncRogueVirtualItemInfoScNotify() {
	notify := &proto.SyncRogueVirtualItemInfoScNotify{
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			// RogueCoin: g.GetDbRogue().CurRogue.CosmicFragment,
			// X:     8,
		},
	}

	g.Send(cmd.SyncRogueVirtualItemInfoScNotify, notify)
}

func (g *GamePlayer) SyncRogueCommonPendingActionScNotify(buffIdList []uint32) {
	dbRogue := g.GetDbRogue()
	rogue := g.GetRogue()
	rogue.BuffNum += uint32(len(buffIdList))
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		Action: &proto.RogueCommonPendingAction{
			QueuePosition: rogue.BuffNum,
			RogueAction:   &proto.RogueAction{
				/*
					BuffSelectInfo: &proto.RogueCommonBuffSelectInfo{
						HandbookUnlockBuffIdList: make([]uint32, 0),
						CanRoll:                  true,
						MazeBuffList:             make([]*proto.RogueCommonBuff, 0),
						SelectBuffSourceHint:     1,
						SourceCurCount:           dbRogue.CurRogue.CurSiteId,
						SourceTotalCount:         1,
					},
				*/
			},
		},
		RogueVersionId: dbRogue.CurRogue.RogueMapID,
	}

	/*
		for _, buffId := range buffIdList {
			rogueCommonBuff := &proto.RogueCommonBuff{
				BuffLevel:  1,
				BuffId: buffId,
			}
			notify.RogueCommonPendingAction.RogueAction.BuffSelectInfo.MazeBuffList = append(notify.RogueCommonPendingAction.RogueAction.BuffSelectInfo.MazeBuffList, rogueCommonBuff)
		}
	*/

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

func (g *GamePlayer) SyncEntityBuffChangeListScNotify(buffIdList []uint32) {
	// rogue := g.GetRogue()
	notify := &proto.SyncEntityBuffChangeListScNotify{
		EntityBuffChangeList: make([]*proto.EntityBuffChange, 0),
	}
	// if len(buffIdList) == 0 {
	// 	for id, buff := range rogue.BuffList {
	// 		rntityBuffChangeInfo := &proto.EntityBuffChange{
	// 			AddBuffInfo: &proto.BuffInfo{
	// 				BuffId:    id,
	// 				Level:     buff.Level,
	// 				AddTimeMs: buff.AddTimeMs,
	// 				LifeTime:  -1,
	// 				Count:     4294967295,
	// 			},
	// 			RemoveBuffId: 0,
	// 			EntityId:     6291457,
	// 		}
	// 		notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, rntityBuffChangeInfo)
	// 	}
	// } else {
	// 	for _, id := range buffIdList {
	// 		if rogue.BuffList[id] == nil {
	// 			continue
	// 		}
	// 		rntityBuffChangeInfo := &proto.EntityBuffChange{
	// 			AddBuffInfo: &proto.BuffInfo{
	// 				BuffId:    id,
	// 				Level:     rogue.BuffList[id].Level,
	// 				AddTimeMs: rogue.BuffList[id].AddTimeMs,
	// 				LifeTime:  -1,
	// 				Count:     4294967295,
	// 			},
	// 			RemoveBuffId: 0,
	// 			EntityId:     6291457,
	// 		}
	// 		notify.EntityBuffChangeList = append(notify.EntityBuffChangeList, rntityBuffChangeInfo)
	// 	}
	// }

	g.Send(cmd.SyncEntityBuffChangeListScNotify, notify)
}

// func (g *GamePlayer) CommonRogueUpdateScNotify() {
// 	rogue := g.GetDbRogue()
// 	notify := &proto.CommonRogueUpdateScNotify{
// 		RogueUpdate: &proto.RogueUpdate{
// 			MapId:  rogue.CurRogue.RogueMapID,
// 			AreaId: rogue.CurRogue.CurAreaId,
// 		},
// 	}
//
// 	g.Send(cmd.CommonRogueUpdateScNotify, notify)
// }

func (g *GamePlayer) SyncRogueCommonActionResultScNotify(buffId uint32) {
	rogue := g.GetRogue()
	if rogue.BuffList[buffId] == nil {
		return
	}
	notify := &proto.SyncRogueCommonActionResultScNotify{
		/*
			Action: &proto.RogueActionResult{
				ActionData: &proto.RogueActionResultData{
					AddBuffList: &proto.RogueBuffData{
						Level:  rogue.BuffList[buffId].Level,
						BuffId: buffId,
					},
				},
				Source: proto.RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_SELECT,
			},
		*/
		RogueVersionId: g.GetDbRogue().CurRogue.RogueMapID,
	}

	g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}

func (g *GamePlayer) GetRogueScoreRewardInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.Info = &proto.RogueScoreRewardInfo{
		// TODO 注意时间
		BeginTime:            1706472000,
		EndTime:              4294967295,
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               22,
		ExploreScore:         5000, // 本期分数
	}

	g.Send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueTalentInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueTalentInfoScRsp{
		RogueTalentInfo: &proto.RogueTalentInfo{
			RogueTalentList: make([]*proto.RogueTalent, 0),
		},
	}

	for _, talent := range gdconf.GetTalentIDList() {
		rogueTalent := &proto.RogueTalent{
			Status:   proto.RogueTalentStatus_ROGUE_TALENT_STATUS_ENABLE,
			TalentId: talent,
		}
		rsp.RogueTalentInfo.RogueTalentList = append(rsp.RogueTalentInfo.RogueTalentList, rogueTalent)
	}

	g.Send(cmd.GetRogueTalentInfoScRsp, rsp)
}

func (g *GamePlayer) GetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetRogueHandbookDataScRsp{
		HandbookInfo: &proto.RogueHandbook{
			// MiracleList: make([]*proto.RogueHandbookMiracle, 0),
			// RogueEvent:  make([]*proto.RogueHandbookEvent, 0),
			// BuffList:    make([]*proto.RogueHandbookBuff, 0),
		},
	}

	// 解锁全部祝福
	// allBuffList := gdconf.GetAllBuff()
	// for _, id := range allBuffList {
	// 	buff := &proto.RogueHandbookBuff{
	// 		BuffId: id,
	// 	}
	// 	rsp.HandbookInfo.BuffList = append(rsp.HandbookInfo.BuffList, buff)
	// }

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *GamePlayer) GetRogueInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueInfoScRsp)
	rsp.RogueInfo = g.GetRogueInfo()

	g.Send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *GamePlayer) StartRogueCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.StartRogueCsReq, payloadMsg)
	req := msg.(*proto.StartRogueCsReq)

	if req.AreaId == 0 {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	g.NewRogue(req.BaseAvatarIdList, req.AreaId)
	g.NewRogueState()

	//
	g.SyncRogueVirtualItemInfoScNotify()
	syncRogueStatusScNotify := &proto.SyncRogueStatusScNotify{
		Status: proto.RogueStatus_ROGUE_STATUS_DOING,
	}
	g.Send(cmd.SyncRogueStatusScNotify, syncRogueStatusScNotify)
	// 区域通知
	g.SyncRogueMapRoomScNotify()
	g.Send(cmd.SyncServerSceneChangeNotify, nil)
	// 队伍更新通知
	g.SyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_ROGUE), true)
	var buffList []uint32
	g.SyncEntityBuffChangeListScNotify(buffList)
	// g.CommonRogueUpdateScNotify()
	dbRogue := g.GetDbRogue()
	rogue := g.GetRogue()

	scene, _, _ := g.GetRogueScene(dbRogue.CurRogue.RogueSceneMap[dbRogue.CurRogue.CurSiteId].RoomId)
	if scene == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	rsp := &proto.StartRogueScRsp{
		Scene:     scene,
		Lineup:    g.GetLineUpPb(g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ROGUE))),
		RogueInfo: g.GetRogueInfo(),
	}
	rsp.RogueInfo.RogueCurrentInfo = &proto.RogueCurrentInfo{
		PendingAction: &proto.RogueCommonPendingAction{QueuePosition: rogue.BuffNum},
		RogueAeonInfo: &proto.GameAeonInfo{
			IsUnlocked: true,
			AeonId:     req.AeonId, // 解锁的命途
		},
		Status: proto.RogueStatus_ROGUE_STATUS_DOING,
	}

	g.GetBattleState().BattleType = spb.BattleType_Battle_ROGUE
	g.Send(cmd.StartRogueScRsp, rsp)
}

// 获取当前模拟宇宙信息
func (g *GamePlayer) GetRogueInfo() *proto.RogueInfo {
	cfRogueManager := gdconf.GetRogueManager()
	rogueInfo := &proto.RogueInfo{
		RogueGetInfo: &proto.RogueGetInfo{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: cfRogueManager.BeginTime,
				EndTime:   cfRogueManager.EndTime,
				Season:    cfRogueManager.RogueSeason,
			},
			RogueScoreRewardInfo: &proto.RogueScoreRewardInfo{
				BeginTime:              cfRogueManager.BeginTime,
				EndTime:                cfRogueManager.EndTime,
				HasTakenInitialScore:   true,
				PoolId:                 22,
				PoolRefreshed:          true,
				TakenNormalFreeRowList: nil,  // 已领取奖励
				ExploreScore:           5000, // 本期得分
			},
			RogueAreaInfo: &proto.RogueAreaInfo{RogueAreaList: g.GetRogueArea()},
			RogueAeonInfo: &proto.RogueAeonInfo{
				// TODO
				AeonIdList:      []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9},
				IsUnlocked:      true,
				UnlockedAeonNum: 9,
			},
		},
	}

	return rogueInfo
}

// 获取模拟宇宙关卡信息
func (g *GamePlayer) GetRogueArea() []*proto.RogueArea {
	rogueAreaList := make([]*proto.RogueArea, 0)
	cfRogueManager := gdconf.GetRogueManager()
	for _, rogueArea := range cfRogueManager.RogueAreaIDList {
		dbRogueArea := g.GetDbRogueArea(rogueArea)
		RogueArea := &proto.RogueArea{
			AreaId:     dbRogueArea.AreaId,
			AreaStatus: proto.RogueAreaStatus(dbRogueArea.RogueAreaStatus),
		}
		rogueAreaList = append(rogueAreaList, RogueArea)
	}

	return rogueAreaList
}

// 新建模拟宇宙
func (g *GamePlayer) NewRogue(avatarIdList []uint32, areaId uint32) {
	// 更新队伍
	db := g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	db.LeaderSlot = 0
	db.AvatarIdList = make(map[uint32]*spb.LineAvatarList)
	if avatarIdList != nil {
		for id, avatarId := range avatarIdList {
			db.AvatarIdList[uint32(id)] = &spb.LineAvatarList{AvatarId: avatarId, Slot: uint32(id)}
		}
	}

	// 获取地图
	rogueAreaConfig := gdconf.GetRogueAreaConfigById(areaId)
	if rogueAreaConfig == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}
	rogueMapID := (rogueAreaConfig.AreaProgress * 100) + rogueAreaConfig.Difficulty
	rogueMap := gdconf.GetRogueMapById(rogueMapID)
	if rogueMap == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	rogueSceneMap := make(map[uint32]*spb.RogueRoom)
	for id, rogue := range rogueMap.SiteList {
		rogueSceneMap[id] = &spb.RogueRoom{
			RoomId:         gdconf.GetRogueRoomIDBySiteID(id),
			RoomStatus:     spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_NONE,
			NextSiteIdList: rogue.NextSiteIDList,
		}
	}

	rogueSceneMap[rogueMap.StartId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY

	g.GetDbRogue().CurRogue = &spb.CurRogue{
		CurAreaId:      areaId,
		CurSiteId:      rogueMap.StartId,
		CosmicFragment: 50,
		RogueSceneMap:  rogueSceneMap,
		RogueMapID:     rogueMapID,
	}
}

func (g *GamePlayer) GetRogueMap() *proto.RogueMapInfo {
	rogue := g.GetDbRogue()
	roomMap := &proto.RogueMapInfo{
		MapId:     rogue.CurRogue.RogueMapID,
		AreaId:    rogue.CurRogue.CurAreaId,
		CurSiteId: rogue.CurRogue.CurSiteId, // 当前id
		CurRoomId: rogue.CurRogue.RogueSceneMap[rogue.CurRogue.CurSiteId].RoomId,
		RoomList:  make([]*proto.RogueRoom, 0),
	}

	for id, rogueScene := range rogue.CurRogue.RogueSceneMap {
		roomList := &proto.RogueRoom{
			SiteId:    id,
			RoomId:    rogueScene.RoomId,
			CurStatus: proto.RogueRoomStatus(rogueScene.RoomStatus),
		}

		roomMap.RoomList = append(roomMap.RoomList, roomList)
	}

	return roomMap
}

func (g *GamePlayer) GetRogueScene(roomId uint32) (*proto.SceneInfo, map[uint32]*AvatarEntity, map[uint32]*MonsterEntity) {
	rogueRoom := gdconf.GetRogueRoomById(roomId)
	if rogueRoom == nil {
		return nil, nil, nil
	}
	mapEntrance := gdconf.GetMapEntranceById(rogueRoom.MapEntrance)
	if mapEntrance == nil {
		return nil, nil, nil
	}

	leaderEntityId := g.GetNextGameObjectGuid()
	scene := &proto.SceneInfo{
		ClientPosVersion:   5,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		EntryId:            rogueRoom.MapEntrance,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        nil,
		LightenSectionList: nil,
		EntityList:         nil,
		GroupStateList:     nil,
	}

	// 获取场景实体
	monsterEntity := make(map[uint32]*MonsterEntity, 0)
	// 添加角色信息
	avatarEntity := make(map[uint32]*AvatarEntity, 0)
	entityGroupList := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	startGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, rogueRoom.GroupID)
	anchor := startGroup.AnchorList[0]
	baseAvatarIdList := g.GetLineUpById(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	for id, avatarList := range baseAvatarIdList.AvatarIdList {
		if avatarList == nil || avatarList.AvatarId == 0 {
			continue
		}
		entityId := g.GetNextGameObjectGuid()
		entityList := &proto.SceneEntityInfo{
			Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarList.AvatarId,
			},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(anchor.PosX * 1000),
					Y: int32(anchor.PosY * 1000),
					Z: int32(anchor.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: int32(anchor.RotX * 1000),
					Y: int32(anchor.RotY * 1000),
					Z: int32(anchor.RotZ * 1000),
				},
			},
		}
		// 为进入场景的角色设置与上面相同的实体id
		if id == 0 {
			entityList.EntityId = leaderEntityId
			avatarEntity[leaderEntityId] = &AvatarEntity{
				AvatarId: avatarList.AvatarId,
			}
		} else {
			entityList.EntityId = entityId
			avatarEntity[entityId] = &AvatarEntity{
				AvatarId: avatarList.AvatarId,
			}
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroupList)
	for groupID, ida := range rogueRoom.GroupWithContent {
		sceneGroup := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, stou32(groupID))
		if sceneGroup == nil {
			continue
		}
		scene.GroupIdList = append(scene.GroupIdList, stou32(groupID))

		sceneGroupState := &proto.SceneGroupState{
			GroupId:   stou32(groupID),
			IsDefault: true,
		}

		scene.GroupStateList = append(scene.GroupStateList, sceneGroupState)

		// 添加物品实体
		entityGroupLists := g.GetRoguePropByID(sceneGroup, stou32(groupID))
		// 添加怪物实体
		entityGroupLists, x := g.GetRogueNPCMonsterByID(entityGroupLists, sceneGroup, stou32(groupID), monsterEntity, ida)
		monsterEntity = x
		// 添加NPC实体
		entityGroupLists = g.GetNPCByID(entityGroupLists, sceneGroup)
		if len(entityGroupLists.EntityList) != 0 {
			scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
		}
	}

	return scene, avatarEntity, monsterEntity
}

func (g *GamePlayer) GetRoguePropByID(sceneGroup *gdconf.GoppLevelGroup, groupID uint32) *proto.SceneEntityGroupInfo {
	entityGroupLists := &proto.SceneEntityGroupInfo{
		GroupId:    groupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, propList := range sceneGroup.PropList {
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,     // 文件名后那个G
			InstId:   propList.ID, // ID
			EntityId: g.GetNextGameObjectGuid(),
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(propList.PosX * 1000),
					Y: int32(propList.PosY * 1000),
					Z: int32(propList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(propList.RotY * 1000),
					Z: 0,
				},
			},
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: 0,               // gdconf.GetPropState(strconv.Itoa(int(propList.PropID))),
			},
		}
		if propList.State != "CheckPointDisable" && propList.State != "CheckPointEnable" {
			entityList.Prop.PropState = 8 // 解锁
		}
		if propList.PropID == 1000 || propList.PropID == 1021 || propList.PropID == 1022 || propList.PropID == 1023 {
			index := 0
			if propList.Name == "Door2" {
				index = 1
			}
			room := g.GetCurDbRoom()
			if propList.Name == "Door1" && len(room.NextSiteIdList) == 1 {
				continue
			}
			if len(room.NextSiteIdList) == 1 {
				index = 0
			}
			if len(room.NextSiteIdList) > 0 {
				siteId := room.NextSiteIdList[index]
				nextRoom := g.GetDbRoomBySiteId(siteId)
				exceRoom := gdconf.GetRogueRoomById(nextRoom.RoomId)

				switch exceRoom.RogueRoomType {
				case 3, 8:
					entityList.Prop.PropId = 1022
				case 5:
					entityList.Prop.PropId = 1023
				default:
					entityList.Prop.PropId = 1021
				}
				entityList.Prop.ExtraInfo = &proto.PropExtraInfo{
					InfoOneofCase: &proto.PropExtraInfo_RogueInfo{
						RogueInfo: &proto.PropRogueInfo{
							RoomId: nextRoom.RoomId,
							SiteId: siteId,
						},
					},
				}
			} else {
				entityList.Prop.PropId = 1000
			}
			entityList.Prop.PropState = 1
		}
		entityGroupLists.EntityList = append(entityGroupLists.EntityList, entityList)
	}
	return entityGroupLists
}

func (g *GamePlayer) GetRogueNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, groupID uint32, entityMap map[uint32]*MonsterEntity, ida uint32) (*proto.SceneEntityGroupInfo, map[uint32]*MonsterEntity) {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := g.GetNextGameObjectGuid()
		rogueMonsterID := gdconf.GetRogueMonsterGroupByGroupID(ida)
		rogueMonster := gdconf.GetRogueMonsterByRogueMonsterID(rogueMonsterID)
		entityList := &proto.SceneEntityInfo{
			GroupId:  groupID,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(monsterList.PosX * 1000),
					Y: int32(monsterList.PosY * 1000),
					Z: int32(monsterList.PosZ * 1000),
				},
				Rot: &proto.Vector{
					X: 0,
					Y: int32(monsterList.RotY * 1000),
					Z: 0,
				},
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.BasicBin.WorldLevel,
				MonsterId:  rogueMonster.NpcMonsterID,
				EventId:    rogueMonster.EventID,
			},
		}
		// 添加实体
		entityMap[entityId] = &MonsterEntity{
			// MonsterEId: rogueMonster.EventID,
			// GroupId:    groupID,
			// Pos: &Vector{
			// 	X: int32(monsterList.PosX * 1000),
			// 	Y: int32(monsterList.PosY * 1000),
			// 	Z: int32(monsterList.PosZ * 1000),
			// },
			// Rot: &Vector{
			// 	X: 0,
			// 	Y: int32(monsterList.RotY * 1000),
			// 	Z: 0,
			// },
		}
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList, entityMap
}

// 模拟宇宙攻击事件
func (g *GamePlayer) RogueSceneCastSkillCsReq(rsp *proto.SceneCastSkillScRsp) {
	rsp.BattleInfo.BattleTargetInfo = make(map[uint32]*proto.BattleTargetList)
	rsp.BattleInfo.BattleTargetInfo[2] = &proto.BattleTargetList{
		BattleTargetList: make([]*proto.BattleTarget, 0),
	}
	battleTargetList := make([]*proto.BattleTarget, 0)
	bTL := []uint32{50008, 50007, 50006, 50005, 50004, 50003, 50001, 30072, 30071, 30070, 30069, 30068, 30067, 30066, 30065, 30064, 30063, 30062, 30061, 30060, 30059, 30058, 30057, 30056, 30055, 30054, 30053, 30046, 30045, 30044, 30043, 30042, 30041, 30040, 30039, 30037, 30034, 30024, 30021, 30016, 30012, 30008, 30006, 30004, 30002, 20074, 20069, 20068, 20067, 20066, 20064, 20060, 20058, 20052, 20050, 20049, 20047, 20046, 20045, 20044, 20043, 20040, 20039, 20034, 20033, 20032, 20031, 20024, 20021, 20019, 20018, 20017, 20016, 20015, 20013, 20010, 20009, 20008, 20002}
	for _, id := range bTL {
		battleTarget := &proto.BattleTarget{
			Id: id,
		}
		battleTargetList = append(battleTargetList, battleTarget)
	}
	rsp.BattleInfo.BattleTargetInfo[2].BattleTargetList = battleTargetList
	// 添加角色
	// rsp.BattleInfo.BattleAvatarList = g.GetBattleAvatarList(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	// 添加buff
	rsp.BattleInfo.BuffList = make([]*proto.BattleBuff, 0)
	for id, buff := range g.GetRogueBuff() {
		battleBuff := &proto.BattleBuff{
			Id:         id,
			Level:      buff.Level,
			OwnerIndex: 4294967295,
			WaveFlag:   4294967295,
		}
		rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, battleBuff)
	}
	// 路途buff！！！
	rsp.BattleInfo.BuffList = append(rsp.BattleInfo.BuffList, &proto.BattleBuff{
		Id:       1000111,
		Level:    1,
		WaveFlag: 4294967295,
		DynamicValues: map[string]float32{
			"SkillIndex": 1,
		},
		TargetIndexList: []uint32{1},
	})
	g.Send(cmd.SceneCastSkillScRsp, rsp)
}

// 模拟宇宙攻击事件结算
func (g *GamePlayer) RoguePVEBattleResultCsReq(req *proto.PVEBattleResultCsReq, rsp *proto.PVEBattleResultScRsp) {
	// battle := g.GetRogueBattle()[req.BattleId]
	// 队伍状态通知
	// g.ChallengeSyncLineupNotify(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	// rsp.BattleAvatarList = g.GetBattleAvatarList(uint32(proto.ExtraLineupType_LINEUP_ROGUE))
	// buff同步
	var buffList []uint32
	g.SyncEntityBuffChangeListScNotify(buffList)
	// 积分同步
	g.SyncRogueVirtualItemInfoScNotify()
	// 物品增加通知
	var pileItem []*Material
	pileItem = append(pileItem, &Material{
		Tid: 31,
		Num: 21,
	})
	g.AddMaterial(pileItem)

	// 祝福选择页通知 SyncRogueCommonPendingActionScNotify
	buffIdList := gdconf.GetBuffListByNum(3)
	g.SyncRogueCommonPendingActionScNotify(buffIdList)
	// 场景实体刷新通知 SceneGroupRefreshScNotify （门和删除刚刚战斗的实体）
	// 删除实体
	nitify := &proto.SceneGroupRefreshScNotify{
		GroupRefreshList: make([]*proto.GroupRefreshInfo, 0),
	}
	// for _, eventId := range battle.monsterEntityMap {
	// 	entity := g.GetSceneEntity().MonsterEntity[eventId]
	// 	if entity != nil {
	// 		groupRefreshInfo := &proto.SceneGroupRefreshInfo{
	// 			GroupId: entity.GroupId,
	// 			RefreshEntity: []*proto.SceneEntityRefreshInfo{
	// 				{
	// 					DelEntity: eventId,
	// 				},
	// 			},
	// 		}
	// 		nitify.GroupRefreshInfo = append(nitify.GroupRefreshInfo, groupRefreshInfo)
	// 		delete(g.GetSceneEntity().MonsterEntity, eventId)
	// 	}
	// }
	// 刷新门
	g.Send(cmd.SceneGroupRefreshScNotify, nitify)

	g.Send(cmd.PVEBattleResultScRsp, rsp)
}

func (g *GamePlayer) QuitRogueCsReq(payloadMsg []byte) {

	g.Send(cmd.QuitRogueScRsp, nil)
}

func (g *GamePlayer) LeaveRogueCsReq(payloadMsg []byte) {
	curLine := g.GetCurLineUp()
	rsp := &proto.LeaveRogueScRsp{
		RogueInfo: g.GetRogueInfo(),
		Lineup:    g.GetLineUpPb(curLine),
		Scene:     g.GetSceneInfo(g.GetScene().EntryId, g.GetPosPb(), g.GetRotPb(), curLine),
	}

	g.Send(cmd.LeaveRogueScRsp, rsp)
}

func (g *GamePlayer) HandleRogueCommonPendingActionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.HandleRogueCommonPendingActionCsReq, payloadMsg)
	req := msg.(*proto.HandleRogueCommonPendingActionCsReq)
	buffSelectResult := req.GetBuffSelectResult()
	if buffSelectResult == nil {
		return
	}
	rogue := g.GetRogue()
	var buffIdList []uint32
	// 祝福通知
	if rogue.BuffList[buffSelectResult.BuffId] == nil {
		buffIdList = append(buffIdList, buffSelectResult.BuffId)
	} else {

	}

	g.RogueAddBuff(buffSelectResult.BuffId)

	g.SyncEntityBuffChangeListScNotify(buffIdList)
	// 添加后通知启动
	g.SyncRogueCommonActionResultScNotify(buffSelectResult.BuffId)
	// 模拟宇宙图鉴更新通知？ SyncRogueHandbookDataUpdateScNotify
	// 模拟宇宙常见操作结果通知 SyncRogueCommonActionResultScNotify // add buff, buff状态
	rsp := &proto.HandleRogueCommonPendingActionScRsp{
		// Times: rogue.BuffNum,
	}

	g.Send(cmd.HandleRogueCommonPendingActionScRsp, rsp)
}

func (g *GamePlayer) EnterRogueMapRoomCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.EnterRogueMapRoomCsReq, payloadMsg)
	req := msg.(*proto.EnterRogueMapRoomCsReq)
	curRogue := g.GetCurDbRogue()
	if curRogue.RogueSceneMap[req.SiteId] == nil && curRogue.RogueSceneMap[req.SiteId].RoomId != req.RoomId {
		return
	}
	curRogue.RogueSceneMap[curRogue.CurSiteId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_FINISH
	curRogue.RogueSceneMap[req.SiteId].RoomStatus = spb.RoomStatus_RogueRoomStatus_ROGUE_ROOM_STATUS_PLAY
	curRogue.CurSiteId = req.SiteId
	// 更新通知
	g.SyncRogueMapRoomScNotify()
	scene, _, _ := g.GetRogueScene(req.RoomId)
	if scene == nil {
		rsp := &proto.StartRogueScRsp{
			Retcode: uint32(proto.Retcode_RET_FIGHT_ACTIVITY_STAGE_NOT_OPEN),
		}
		g.Send(cmd.StartRogueScRsp, rsp)
		return
	}

	rsp := &proto.EnterRogueMapRoomScRsp{
		Lineup:    g.GetLineUpPb(g.GetBattleLineUpById(uint32(proto.ExtraLineupType_LINEUP_ROGUE))),
		CurSiteId: req.SiteId,
		Retcode:   0,
		Scene:     scene,
	}

	g.Send(cmd.EnterRogueMapRoomScRsp, rsp)
}
