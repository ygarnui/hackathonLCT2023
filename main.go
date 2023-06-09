package main

import u "PDFGenerator/stringGenerator"

const (
	StepBody         = 10
	MoscowTaxFrom    = "MoscowTaxFrom"
	MoscowTaxTo      = "MoscowTaxTo"
	PropertyTaxFrom  = "PropertyTaxFrom"
	PropertyTaxTo    = "PropertyTaxTo"
	ProfitTaxFrom    = "ProfitTaxFrom"
	ProfitTaxTo      = "ProfitTaxTo"
	TransportTaxFrom = "TransportTaxFrom"
	TransportTaxTo   = "TransportTaxTo"
	OtherTaxFrom     = "OtherTaxFrom"
	OtherTaxTo       = "OtherTaxTo"
	GovReg           = "GovReg"
	PatentPrice      = "PatentPrice"
)

func main() {

	val := make(map[string]float64)
	val[MoscowTaxFrom] = 1
	val[MoscowTaxTo] = 1
	val[PropertyTaxFrom] = 1
	val[PropertyTaxTo] = 1
	val[ProfitTaxFrom] = 1
	val[ProfitTaxTo] = 1
	val[TransportTaxFrom] = 1
	val[TransportTaxTo] = 1
	val[OtherTaxFrom] = 1
	val[OtherTaxTo] = 1
	val[GovReg] = 1
	val[PatentPrice] = 1
	val["PersonalCount"] = 30
	val["PersonalSalaryTo"] = 3210
	val["PersonalSocialTo"] = 4321000
	val["PersonalPensionTo"] = 543210000000
	val["PersonalNDFLTo"] = 1.25
	val["EstateRentTo"] = 1
	val["EstateTaxTo"] = 1
	val["MoscowTaxTo"] = 123456789000000
	val["PropertyTaxTo"] = 123456789000000000
	val["ProfitTaxTo"] = 1
	val["TransportTaxTo"] = 1
	val["PatentPriceTo"] = 1
	val["OtherTaxTo"] = 1
	val["CapBuildTo"] = 1
	val["CapRebuildTo"] = 1
	val["FinancialTo"] = 1
	val["EquipmentPrice"] = 1
	val["EstatePriceTo"] = 1

	val["PersonalFrom"] = 100000
	val["PersonalTo"] = 1000000
	val["EstateFrom"] = 2000000
	val["EstateTo"] = 20000000
	val["TaxFrom"] = 3000000
	val["TaxTo"] = 30000000
	val["ServiceFrom"] = 4000000
	val["ServiceTo"] = 40000000

	valStr := make(map[string]string)

	valStr["UserName"] = "ygarnui"
	valStr["Industry"] = "Общее машиностроение (в т.ч. оборудование пищевой переработки, дорожностроительная и сельскохозяйственная техника)"
	valStr["Industry"] = "Производство оружия, боеприпасов, спецхимии, военных машин"

	valStr["Organization"] = "тип организации"
	valStr["WorkersCount"] = "30"
	valStr["District"] = "ЗАО"

	u.GenerateReport(valStr, val)
}
