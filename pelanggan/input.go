package pelanggan

type PelangganInput struct {
	IDPelanggan string `json:"id_pelanggan" binding:"required"`
	NamaPelanggan string `json:"nama_pelanggan" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NomorKWH string `json:"nomor_kwh" binding:"required"`
	Alamat string `json:"alamat" binding:"required"`
	IDTarif string `json:"id_tarif" binding:"required"`
}