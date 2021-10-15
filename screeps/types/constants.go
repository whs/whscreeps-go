package types

import "fmt"

type Error int8

const (
	Ok                     Error = 0
	ErrNotOwner            Error = -1
	ErrNoPath              Error = -2
	ErrNameExists          Error = -3
	ErrBusy                Error = -4
	ErrNotFound            Error = -5
	ErrNotEnoughEnergy     Error = -6
	ErrNotEnoughResources  Error = -6
	ErrInvalidTarget       Error = -7
	ErrFull                Error = -8
	ErrNotInRange          Error = -9
	ErrInvalidArgs         Error = -10
	ErrTired               Error = -11
	ErrNoBodypart          Error = -12
	ErrNotEnoughExtensions Error = -6
	ErrRclNotEnough        Error = -14
	ErrGclNotEnough        Error = -15
)

func (e Error) Error() string {
	return fmt.Sprintf("errno %d", e)
}

type FindType = uint8

const (
	FindExitTop                  FindType = 1
	FindExitRight                FindType = 3
	FindExitBottom               FindType = 5
	FindExitLeft                 FindType = 7
	FindExit                     FindType = 10
	FindCreeps                   FindType = 101
	FindMyCreeps                 FindType = 102
	FindHostileCreeps            FindType = 103
	FindSourcesActive            FindType = 104
	FindSources                  FindType = 105
	FindDroppedResources         FindType = 106
	FindStructures               FindType = 107
	FindMyStructures             FindType = 108
	FindHostileStructures        FindType = 109
	FindFlags                    FindType = 110
	FindConstructionSites        FindType = 111
	FindMySpawns                 FindType = 112
	FindHostileSpawns            FindType = 113
	FindMyConstructionSites      FindType = 114
	FindHostileConstructionSites FindType = 115
	FindMinerals                 FindType = 116
	FindNukes                    FindType = 117
	FindTombstones               FindType = 118
	FindPowerCreeps              FindType = 119
	FindMyPowerCreeps            FindType = 120
	FindHostilePowerCreeps       FindType = 121
	FindDeposits                 FindType = 122
	FindRuins                    FindType = 123
)

type BodyPart = string

const (
	Move         BodyPart = "move"
	Work         BodyPart = "work"
	Carry        BodyPart = "carry"
	Attack       BodyPart = "attack"
	RangedAttack BodyPart = "ranged_attack"
	Tough        BodyPart = "tough"
	Heal         BodyPart = "heal"
	Claim        BodyPart = "claim"
)

type Direction = uint8

const (
	Top         Direction = 1
	TopRight    Direction = 2
	Right       Direction = 3
	BottomRight Direction = 4
	Bottom      Direction = 5
	BottomLeft  Direction = 6
	Left        Direction = 7
	TopLeft     Direction = 8
)

type Resource = string

const (
	ResourceEnergy Resource = "energy"
	ResourcePower  Resource = "power"

	ResourceHydrogen  Resource = "H"
	ResourceOxygen    Resource = "O"
	ResourceUtrium    Resource = "U"
	ResourceLemergium Resource = "L"
	ResourceKeanium   Resource = "K"
	ResourceZynthium  Resource = "Z"
	ResourceCatalyst  Resource = "X"
	ResourceGhodium   Resource = "G"

	ResourceSilicon Resource = "silicon"
	ResourceMetal   Resource = "metal"
	ResourceBiomass Resource = "biomass"
	ResourceMist    Resource = "mist"

	ResourceHydroxide       Resource = "OH"
	ResourceZynthiumKeanite Resource = "ZK"
	ResourceUtriumLemergite Resource = "UL"

	ResourceUtriumHydride    Resource = "UH"
	ResourceUtriumOxide      Resource = "UO"
	ResourceKeaniumHydride   Resource = "KH"
	ResourceKeaniumOxide     Resource = "KO"
	ResourceLemergiumHydride Resource = "LH"
	ResourceLemergiumOxide   Resource = "LO"
	ResourceZynthiumHydride  Resource = "ZH"
	ResourceZynthiumOxide    Resource = "ZO"
	ResourceGhodiumHydride   Resource = "GH"
	ResourceGhodiumOxide     Resource = "GO"

	ResourceUtriumAcid        Resource = "UH2O"
	ResourceUtriumAlkalide    Resource = "UHO2"
	ResourceKeaniumAcid       Resource = "KH2O"
	ResourceKeaniumAlkalide   Resource = "KHO2"
	ResourceLemergiumAcid     Resource = "LH2O"
	ResourceLemergiumAlkalide Resource = "LHO2"
	ResourceZynthiumAcid      Resource = "ZH2O"
	ResourceZynthiumAlkalide  Resource = "ZHO2"
	ResourceGhodiumAcid       Resource = "GH2O"
	ResourceGhodiumAlkalide   Resource = "GHO2"

	ResourceCatalyzedUtriumAcid        Resource = "XUH2O"
	ResourceCatalyzedUtriumAlkalide    Resource = "XUHO2"
	ResourceCatalyzedKeaniumAcid       Resource = "XKH2O"
	ResourceCatalyzedKeaniumAlkalide   Resource = "XKHO2"
	ResourceCatalyzedLemergiumAcid     Resource = "XLH2O"
	ResourceCatalyzedLemergiumAlkalide Resource = "XLHO2"
	ResourceCatalyzedZynthiumAcid      Resource = "XZH2O"
	ResourceCatalyzedZynthiumAlkalide  Resource = "XZHO2"
	ResourceCatalyzedGhodiumAcid       Resource = "XGH2O"
	ResourceCatalyzedGhodiumAlkalide   Resource = "XGHO2"

	ResourceOps Resource = "ops"

	ResourceUtriumBar    Resource = "utrium_bar"
	ResourceLemergiumBar Resource = "lemergium_bar"
	ResourceZynthiumBar  Resource = "zynthium_bar"
	ResourceKeaniumBar   Resource = "keanium_bar"
	ResourceGhodiumMelt  Resource = "ghodium_melt"
	ResourceOxidant      Resource = "oxidant"
	ResourceReductant    Resource = "reductant"
	ResourcePurifier     Resource = "purifier"
	ResourceBattery      Resource = "battery"

	ResourceComposite Resource = "composite"
	ResourceCrystal   Resource = "crystal"
	ResourceLiquid    Resource = "liquid"

	ResourceWire       Resource = "wire"
	ResourceSwitch     Resource = "switch"
	ResourceTransistor Resource = "transistor"
	ResourceMicrochip  Resource = "microchip"
	ResourceCircuit    Resource = "circuit"
	ResourceDevice     Resource = "device"

	ResourceCell     Resource = "cell"
	ResourcePhlegm   Resource = "phlegm"
	ResourceTissue   Resource = "tissue"
	ResourceMuscle   Resource = "muscle"
	ResourceOrganoid Resource = "organoid"
	ResourceOrganism Resource = "organism"

	ResourceAlloy      Resource = "alloy"
	ResourceTube       Resource = "tube"
	ResourceFixtures   Resource = "fixtures"
	ResourceFrame      Resource = "frame"
	ResourceHydraulics Resource = "hydraulics"
	ResourceMachine    Resource = "machine"

	ResourceCondensate  Resource = "condensate"
	ResourceConcentrate Resource = "concentrate"
	ResourceExtract     Resource = "extract"
	ResourceSpirit      Resource = "spirit"
	ResourceEmanation   Resource = "emanation"
	ResourceEssence     Resource = "essence"
)

type Structure = string

const (
	StructureSpawn       Structure = "spawn"
	StructureExtension   Structure = "extension"
	StructureRoad        Structure = "road"
	StructureWall        Structure = "constructedWall"
	StructureRampart     Structure = "rampart"
	StructureKeeperLair  Structure = "keeperLair"
	StructurePortal      Structure = "portal"
	StructureController  Structure = "controller"
	StructureLink        Structure = "link"
	StructureStorage     Structure = "storage"
	StructureTower       Structure = "tower"
	StructureObserver    Structure = "observer"
	StructurePowerBank   Structure = "powerBank"
	StructurePowerSpawn  Structure = "powerSpawn"
	StructureExtractor   Structure = "extractor"
	StructureLab         Structure = "lab"
	StructureTerminal    Structure = "terminal"
	StructureContainer   Structure = "container"
	StructureNuker       Structure = "nuker"
	StructureFactory     Structure = "factory"
	StructureInvaderCore Structure = "invaderCore"
)
