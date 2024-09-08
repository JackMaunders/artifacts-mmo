package schema

import "time"

type Character struct {
	Name                    string           `json:"name"`
	Skin                    string           `json:"skin"`
	Level                   int              `json:"level"`
	Xp                      int              `json:"xp"`
	Max_xp                  int              `json:"max_xp"`
	AchievementPoints       int              `json:"achievement_points"`
	Gold                    int              `json:"gold"`
	Speed                   int              `json:"speed"`
	MiningLevel             int              `json:"mining_level"`
	MiningXp                int              `json:"mining_xp"`
	MiningMaxXp             int              `json:"mining_max_xp"`
	WoodcuttingLevel        int              `json:"woodcutting_level"`
	WoodcuttingXp           int              `json:"woodcutting_xp"`
	WoodcuttingMaxXp        int              `json:"woodcutting_max_xp"`
	FishingLevel            int              `json:"fishing_level"`
	FishingXp               int              `json:"fishing_xp"`
	FishingMaxXp            int              `json:"fishing_max_xp"`
	WeaponCraftingLevel     int              `json:"weapon_crafting_level"`
	WeaponCraftingXp        int              `json:"weapon_crafting_xp"`
	WeaponCraftingMaxXp     int              `json:"weapon_crafting_max_xp"`
	GearCraftingLevel       int              `json:"gear_crafting_level"`
	GearCraftingXp          int              `json:"gear_crafting_xp"`
	GearCraftingMaxXp       int              `json:"gear_crafting_max_xp"`
	JewelryCraftingLevel    int              `json:"jewelry_crafting_level"`
	JewelryCraftingXp       int              `json:"jewelry_crafting_xp"`
	JewelryCraftingMaxXp    int              `json:"jewelry_crafting_max_xp"`
	CookingLevel            int              `json:"cooking_level"`
	CookingXp               int              `json:"cooking_xp"`
	CookingMaxXp            int              `json:"cooking_max_xp"`
	Hp                      int              `json:"hp"`
	Haste                   int              `json:"haste"`
	CriticalStrike          int              `json:"critical_strike"`
	Stamina                 int              `json:"stamina"`
	AttackFire              int              `json:"attack_fire"`
	AttachEarth             int              `json:"attack_earth"`
	AttackWater             int              `json:"attack_water"`
	AttackAir               int              `json:"attack_air"`
	DamageFire              int              `json:"damage_fire"`
	DamageEarth             int              `json:"damage_earth"`
	DamageWater             int              `json:"damage_water"`
	DamageAir               int              `json:"damage_air"`
	ResFire                 int              `json:"res_fire"`
	ResEarth                int              `json:"res_earth"`
	ResWater                int              `json:"res_water"`
	ResAir                  int              `json:"res_air"`
	X                       int              `json:"x"`
	Y                       int              `json:"y"`
	Cooldown                Cooldown         `json:"cooldown"`
	CooldownExpiration      time.Time        `json:"cooldown_expiration"`
	WeaponSlot              string           `json:"weapon_slot"`
	ShieldSlot              string           `json:"shield_slot"`
	HelmetSlot              string           `json:"helmet_slot"`
	BodyArmorSlot           string           `json:"body_armor_slot"`
	LegArmorSlot            string           `json:"leg_armor_slot"`
	BootsSlot               string           `json:"boots_slot"`
	Ring1Slot               string           `json:"ring1_slot"`
	Ring2Slot               string           `json:"ring2_slot"`
	AmuletSlot              string           `json:"amulet_slot"`
	Artifact1Slot           string           `json:"artifact1_slot"`
	Artifact2Slot           string           `json:"artifact2_slot"`
	Artifact3Slot           string           `json:"artifact3_slot"`
	Consumable1Slot         string           `json:"consumable1_slot"`
	Consumable1SlotQuantity int              `json:"consumable1_slot_quantity"`
	Consumable2Slot         string           `json:"consumable2_slot"`
	Consumable2SlotQuantity int              `json:"consumable2_slot_quantity"`
	Task                    string           `json:"task"`
	TaskType                string           `json:"task_type"`
	TaskProgress            int              `json:"task_progress"`
	TaskTotal               int              `json:"task_total"`
	InventoryMaxItems       int              `json:"inventory_max_items"`
	Inventory               *[]InventorySlot `json:"inventory,omitempty"`
}

type InventorySlot struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
