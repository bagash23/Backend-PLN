package tagihan

type TagihanFormatter struct {
	IDTagihan string `json:"id_tagihan"`
	IDPenggunaan string `json:"id_penggunaan"`
	IDPelanggan string `json:"id_pelanggan"`
	Bulan int `json:"bulan"`
	Tahun int `json:"tahun"`
	JumlahMeter int `json:"jumlah_meter"`
	Status string `json:"status"`
}

func FormatterTagihan(tagihan Tagihan) TagihanFormatter {
	formatter := TagihanFormatter {
		IDTagihan: tagihan.IDTagihan,
		IDPenggunaan: tagihan.IDPenggunaan,
		IDPelanggan: tagihan.IDPelanggan,
		Bulan: tagihan.Bulan,
		Tahun: tagihan.Tahun,
		JumlahMeter: tagihan.JumlahMeter,
		Status: tagihan.Status,
	}
	return formatter
}