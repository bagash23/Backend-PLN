package penggunaan

type PenggunaanFormatter struct {
	IDPenggunaan string `json:"id_penggunaan"`
	IDPelanggan string `json:"id_pelanggan"`
	Bulan int `json:"bulan"`
	Tahun int `json:"tahun"`
	MeterAwal int `json:"meter_awal"`
	MeterAkhir int `json:"meter_akhir"`
}

func FormatterPenggunaan(penggunaan Penggunaan) PenggunaanFormatter{
	formatter := PenggunaanFormatter {
		IDPenggunaan:  penggunaan.IDPenggunaan,
		IDPelanggan: penggunaan.IDPelanggan,
		Bulan: penggunaan.Bulan,
		Tahun: penggunaan.Tahun,
		MeterAwal: penggunaan.MeterAwal,
		MeterAkhir: penggunaan.MeterAkhir,
	}
	return formatter
}