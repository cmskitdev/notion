package types

// NumberConfig represents configuration for a number property.
type NumberConfig struct {
	Format NumberFormat `json:"format"`
}

// NumberFormat represents the display format for number properties.
type NumberFormat string

// NumberFormat represents the display format for number properties.
//
// See: https://developers.notion.com/reference/property-object#number
const (
	NumberFormatNumber           NumberFormat = "number"
	NumberFormatNumberWithCommas NumberFormat = "number_with_commas"
	NumberFormatPercent          NumberFormat = "percent"
	NumberFormatDollar           NumberFormat = "dollar"
	NumberFormatCanadianDollar   NumberFormat = "canadian_dollar"
	NumberFormatEuro             NumberFormat = "euro"
	NumberFormatPound            NumberFormat = "pound"
	NumberFormatYen              NumberFormat = "yen"
	NumberFormatRuble            NumberFormat = "ruble"
	NumberFormatRupee            NumberFormat = "rupee"
	NumberFormatWon              NumberFormat = "won"
	NumberFormatYuan             NumberFormat = "yuan"
	NumberFormatReal             NumberFormat = "real"
	NumberFormatLira             NumberFormat = "lira"
	NumberFormatRupiah           NumberFormat = "rupiah"
	NumberFormatFrank            NumberFormat = "frank"
	NumberFormatHongKongDollar   NumberFormat = "hong_kong_dollar"
	NumberFormatNewZealandDollar NumberFormat = "new_zealand_dollar"
	NumberFormatKrona            NumberFormat = "krona"
	NumberFormatNorwegianKrone   NumberFormat = "norwegian_krone"
	NumberFormatMexicanPeso      NumberFormat = "mexican_peso"
	NumberFormatRand             NumberFormat = "rand"
	NumberFormatNewTaiwanDollar  NumberFormat = "new_taiwan_dollar"
	NumberFormatDanishKrone      NumberFormat = "danish_krone"
	NumberFormatZloty            NumberFormat = "zloty"
	NumberFormatBaht             NumberFormat = "baht"
	NumberFormatForint           NumberFormat = "forint"
	NumberFormatKoruna           NumberFormat = "koruna"
	NumberFormatShekel           NumberFormat = "shekel"
	NumberFormatChileanPeso      NumberFormat = "chilean_peso"
	NumberFormatPhilippinePeso   NumberFormat = "philippine_peso"
	NumberFormatDirham           NumberFormat = "dirham"
	NumberFormatColombianPeso    NumberFormat = "colombian_peso"
	NumberFormatRiyal            NumberFormat = "riyal"
	NumberFormatRinggit          NumberFormat = "ringgit"
	NumberFormatLeu              NumberFormat = "leu"
	NumberFormatArgentinePeso    NumberFormat = "argentine_peso"
	NumberFormatUruguayanPeso    NumberFormat = "uruguayan_peso"
	NumberFormatSingaporeDollar  NumberFormat = "singapore_dollar"
)
