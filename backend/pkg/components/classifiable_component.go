package components

type PropertyType string
type PropertySubtype string

const (
	Residential PropertyType = "Residential"
	Commercial  PropertyType = "Commercial"
)

// Subtypes for residential properties
const (
	Apartment    PropertySubtype = "Apartment"
	Condo        PropertySubtype = "Condo"
	Duplex       PropertySubtype = "Duplex"
	Multifamily  PropertySubtype = "Multifamily"
	Penthouse    PropertySubtype = "Penthouse"
	SingleFamily PropertySubtype = "SingleFamily"
	Townhome     PropertySubtype = "Townhome"
)

const (
	// Retail
	Bookstore        PropertySubtype = "Bookstore"
	ClothingStore    PropertySubtype = "ClothingStore"
	ConvenienceStore PropertySubtype = "ConvenienceStore"
	ElectronicsStore PropertySubtype = "ElectronicsStore"
	Florist          PropertySubtype = "Florist"
	FurnitureStore   PropertySubtype = "FurnitureStore"
	JewelryStore     PropertySubtype = "JewelryStore"
	LiquorStore      PropertySubtype = "LiquorStore"
	Mall             PropertySubtype = "Mall"
	PetStore         PropertySubtype = "PetStore"
	Pharmacy         PropertySubtype = "Pharmacy"
	ShoeStore        PropertySubtype = "ShoeStore"
	Supermarket      PropertySubtype = "Supermarket"
)

const (
	// Food & Beverage
	Bakery            PropertySubtype = "Bakery"
	Bar               PropertySubtype = "Bar"
	Brewery           PropertySubtype = "Brewery"
	Cafe              PropertySubtype = "Cafe"
	IceCreamShop      PropertySubtype = "IceCreamShop"
	Microbrewery      PropertySubtype = "Microbrewery"
	NightClub         PropertySubtype = "NightClub"
	Restaurant        PropertySubtype = "Restaurant"
	Winery            PropertySubtype = "Winery"
	WineryTastingRoom PropertySubtype = "WineryTastingRoom"
)

const (
	// Services
	AutoRepairShop   PropertySubtype = "AutoRepairShop"
	CarWash          PropertySubtype = "CarWash"
	Clinic           PropertySubtype = "Clinic"
	DaycareCenter    PropertySubtype = "DaycareCenter"
	DryCleaners      PropertySubtype = "DryCleaners"
	GasStation       PropertySubtype = "GasStation"
	Gym              PropertySubtype = "Gym"
	Hotel            PropertySubtype = "Hotel"
	MedicalOffice    PropertySubtype = "MedicalOffice"
	Salon            PropertySubtype = "Salon"
	Spa              PropertySubtype = "Spa"
	TattooParlor     PropertySubtype = "TattooParlor"
	VeterinaryClinic PropertySubtype = "VeterinaryClinic"
)

const (
	// Entertainment
	AmusementPark PropertySubtype = "AmusementPark"
	Arcade        PropertySubtype = "Arcade"
	ArtGallery    PropertySubtype = "ArtGallery"
	BowlingAlley  PropertySubtype = "BowlingAlley"
	ConcertHall   PropertySubtype = "ConcertHall"
	EventVenue    PropertySubtype = "EventVenue"
	GamingCenter  PropertySubtype = "GamingCenter"
	MovieTheater  PropertySubtype = "MovieTheater"
	Museum        PropertySubtype = "Museum"
	SportsArena   PropertySubtype = "SportsArena"
	Theater       PropertySubtype = "Theater"
)

const (
	// Office & Business
	AccountingFirm    PropertySubtype = "AccountingFirm"
	ArchitecturalFirm PropertySubtype = "ArchitecturalFirm"
	ConsultingFirm    PropertySubtype = "ConsultingFirm"
	CoWorkingSpace    PropertySubtype = "CoWorkingSpace"
	CreditUnion       PropertySubtype = "CreditUnion"
	EngineeringFirm   PropertySubtype = "EngineeringFirm"
	InsuranceFirm     PropertySubtype = "InsuranceFirm"
	InvestmentBank    PropertySubtype = "InvestmentBank"
	LawOffice         PropertySubtype = "LawOffice"
	MortgageBank      PropertySubtype = "MortgageBank"
	RealEstateOffice  PropertySubtype = "RealEstateOffice"
	ResearchLab       PropertySubtype = "ResearchLab"
	TechHub           PropertySubtype = "TechHub"
)

const (
	// Other
	DataCenter         PropertySubtype = "DataCenter"
	DistributionCenter PropertySubtype = "DistributionCenter"
	EducationalCenter  PropertySubtype = "EducationalCenter"
	Factory            PropertySubtype = "Factory"
	FitnessCenter      PropertySubtype = "FitnessCenter"
	GreenBuilding      PropertySubtype = "GreenBuilding"
	ParkingGarage      PropertySubtype = "ParkingGarage"
	RecyclingCenter    PropertySubtype = "RecyclingCenter"
	StorageFacility    PropertySubtype = "StorageFacility"
)

type Classifiable struct {
	Type    PropertyType
	Subtype PropertySubtype
}
