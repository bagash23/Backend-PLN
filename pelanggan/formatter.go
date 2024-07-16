package pelanggan


type PelangganFormatter struct {
	IDPelanggan string `json:"id_pelanggan"`
	NamaPelanggan string `json:"nama_pelanggan"`
	Username string `json:"username"`
	Password string `json:"password"`
	NomorKWH string `json:"nomor_kwh"`
	Alamat string `json:"alamat"`
	IDTarif string `json:"id_tarif"`
}

func FormatterPelanggan(pelanggan Pelanggan) PelangganFormatter {
	formatter := PelangganFormatter {
		IDPelanggan: pelanggan.IDPelanggan,
		NamaPelanggan: pelanggan.NamaPelanggan,
		Username: pelanggan.Username,
		Password: pelanggan.Password,
		NomorKWH: pelanggan.NomorKWH,
		Alamat: pelanggan.Alamat,
		IDTarif: pelanggan.IDTarif,
	}
	return formatter
}