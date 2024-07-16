package tagihan

type TagihanInput struct {
	IDTagihan string `json:"id_tagihan" binding:"required"`
	IDPenggunaan string `json:"id_penggunaan" binding:"required"`
	IDPelanggan string `json:"id_pelanggan" binding:"required"`
	Bulan int `json:"bulan" binding:"required"`
	Tahun int `json:"tahun" binding:"required"`
	JumlahMeter int `json:"jumlah_meter" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type ReqUpdateTagihan struct {
	Status string `json:"status"`
}