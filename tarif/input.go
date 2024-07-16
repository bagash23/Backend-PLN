package tarif

type TarifInput struct {
	IDTarif string `json:"id_tarif" binding:"required"`
	Daya string `json:"daya" binding:"required"`
	TarifPerkwh int `json:"tarif_perkwh" binding:"required"`
}

type ReqUpdateTarif struct {	
	Daya string `json:"daya"`
	TarifPerkwh int `json:"tarif_perkwh"`
}