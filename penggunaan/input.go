package penggunaan

type PenggunaanInput struct {
	IDPenggunaan string `json:"id_penggunaan" binding:"required"`
	IDPelanggan string `json:"id_pelanggan" binding:"required"`
	Bulan int `json:"bulan" binding:"required"`
	Tahun int `json:"tahun" binding:"required"`
	MeterAwal int `json:"meter_awal" binding:"required"`
	MeterAkhir int `json:"meter_akhir" binding:"required"`
}

type ReqUpdatePenggunaan struct {
	Bulan int `json:"bulan"`
	Tahun int `json:"tahun"`
	MeterAwal int `json:"meter_awal"`
	MeterAkhir int `json:"meter_akhir"`
}


type DeleteInput struct {
	IDPenggunaan string `json:"id_penggunaan" binding:"required"`
}