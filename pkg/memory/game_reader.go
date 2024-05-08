package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

type GameReader struct {
	offset       Offset
	previousRead data.Data
	Process
}

func NewGameReader(process Process) *GameReader {
	return &GameReader{
		offset:  calculateOffsets(process),
		Process: process,
	}
}

func (gd *GameReader) GetData() data.Data {
	if gd.offset.UnitTable == 0 {
		gd.offset = calculateOffsets(gd.Process)
	}

	roster := gd.getRoster()
	playerUnitPtr, corpse := gd.GetPlayerUnitPtr(roster)

	pu := gd.GetPlayerUnit(playerUnitPtr)
	hover := gd.hoveredData()

	// Quests
	q1 := uintptr(gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x230E9A8, Uint64))
	q2 := uintptr(gd.Process.ReadUInt(q1, Uint64))
	gameQuestsBytes := gd.Process.ReadBytesFromMemory(q2, 85)

	gameQuestsBytes = gameQuestsBytes[3:]

	d := data.Data{
		Corpse:      corpse,
		Monsters:    gd.Monsters(pu.Position, hover),
		PlayerUnit:  pu,
		Items:       gd.Items(pu, hover),
		Objects:     gd.Objects(pu.Position, hover),
		OpenMenus:   gd.openMenus(),
		Roster:      roster,
		HoverData:   hover,
		TerrorZones: gd.TerrorZones(),
		Quests:      gd.getQuests(gameQuestsBytes),
		KeyBindings: gd.GetKeyBindings(),
	}

	if playerUnitPtr == 0 {
		return gd.previousRead
	}

	gd.previousRead = d

	return d
}

func (gd *GameReader) InGame() bool {
	pu, _ := gd.GetPlayerUnitPtr([]data.RosterMember{})

	return pu > 0
}

func (gd *GameReader) openMenus() data.OpenMenus {
	uiBase := gd.Process.moduleBaseAddressPtr + gd.offset.UI - 0xA

	buffer := gd.Process.ReadBytesFromMemory(uiBase, 0x16D)

	isMapShown := gd.Process.ReadUInt(gd.Process.moduleBaseAddressPtr+gd.offset.UI, Uint8)

	return data.OpenMenus{
		Inventory:     buffer[0x01] != 0,
		LoadingScreen: buffer[0x16C] != 0,
		NPCInteract:   buffer[0x08] != 0,
		NPCShop:       buffer[0x0B] != 0,
		Stash:         buffer[0x18] != 0,
		Waypoint:      buffer[0x13] != 0,
		MapShown:      isMapShown != 0,
		SkillTree:     buffer[0x04] != 0,
		Character:     buffer[0x02] != 0,
		QuitMenu:      buffer[0x09] != 0,
		Cube:          buffer[0x19] != 0,
		SkillSelect:   buffer[0x03] != 0,
		Anvil:         buffer[0x0D] != 0,
	}
}

func (gd *GameReader) hoveredData() data.HoverData {
	hoverAddressPtr := gd.Process.moduleBaseAddressPtr + gd.offset.Hover
	hoverBuffer := gd.Process.ReadBytesFromMemory(hoverAddressPtr, 12)
	isUnitHovered := ReadUIntFromBuffer(hoverBuffer, 0, Uint16)
	if isUnitHovered > 0 {
		hoveredType := ReadUIntFromBuffer(hoverBuffer, 0x04, Uint32)
		hoveredUnitID := ReadUIntFromBuffer(hoverBuffer, 0x08, Uint32)

		return data.HoverData{
			IsHovered: true,
			UnitID:    data.UnitID(hoveredUnitID),
			UnitType:  int(hoveredType),
		}
	}

	return data.HoverData{}
}

func (gd *GameReader) getStatsList(statListPtr uintptr) stat.Stats {
	statList := gd.Process.ReadUInt(statListPtr, Uint64)
	statCount := gd.Process.ReadUInt(statListPtr+0x8, Uint64)
	if statCount == 0 {
		return []stat.Data{}
	}

	var stats = make([]stat.Data, 0)
	statBuffer := gd.Process.ReadBytesFromMemory(uintptr(statList), statCount*10)
	for i := 0; i < int(statCount); i++ {
		offset := uint(i * 8)
		statLayer := ReadUIntFromBuffer(statBuffer, offset, Uint16)
		statEnum := ReadUIntFromBuffer(statBuffer, offset+0x2, Uint16)
		statValue := ReadUIntFromBuffer(statBuffer, offset+0x4, Uint32)

		value := int(statValue)
		switch stat.ID(statEnum) {
		case stat.Life,
			stat.MaxLife,
			stat.Mana,
			stat.MaxMana,
			stat.Stamina,
			stat.LifePerLevel,
			stat.ManaPerLevel:
			value = int(statValue >> 8)
		case stat.ColdLength,
			stat.PoisonLength:
			value = int(statValue / 25)
		}

		stats = append(stats, stat.Data{
			ID:    stat.ID(statEnum),
			Value: value,
			Layer: int(statLayer),
		})
	}

	return stats
}

// TODO: Take a look to better ways to get this data, now it's very flakky, is just a random memory position + not in game
func (gd *GameReader) InCharacterSelectionScreen() bool {
	uiBase := gd.Process.moduleBaseAddressPtr + gd.offset.UI - 0xA

	return gd.Process.ReadUInt(uiBase, Uint8) != 1 && gd.Process.ReadUInt(gd.moduleBaseAddressPtr+0x1EC5AA8, Uint64) == 0
}

func (gd *GameReader) GetSelectedCharacterName() string {
	return gd.Process.ReadStringFromMemory(gd.Process.moduleBaseAddressPtr+0x21A2374, 0)
}
