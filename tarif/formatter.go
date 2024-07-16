package tarif

type TarifFormatter struct {
	IDTarif string `json:"id_tarif"`
	Daya string `json:"daya"`
	TarifPerkwh int `json:"tarif_perkwh"`
}

func FormatterTarif(tarif Tarif) TarifFormatter {
	formatter := TarifFormatter {
		IDTarif: tarif.IDTarif,
		Daya: tarif.Daya,
		TarifPerkwh: tarif.TarifPerkwh,
	}
	return formatter
}