package pembayaran

type PembayaranFormatter struct {
	IDPembayaran string  `json:"id_pembayaran"`
	IDTagihan string `json:"id_tagihan"`
	IDPelanggan string `json:"id_pelanggan"`
	TglBayar string `json:"tgl_bayar"`
	BiayaAdmin int `json:"biaya_admin"`
	TotalBayar int `json:"total_bayar"`
	IDUser string `json:"id_user"`
}

func FormatterPembayaran(pembayaran Pembayaran) PembayaranFormatter{
	formatter := PembayaranFormatter {
		IDPembayaran: pembayaran.IDPembayaran,
		IDTagihan: pembayaran.IDTagihan,
		IDPelanggan: pembayaran.IDPelanggan,
		TglBayar: pembayaran.TglBayar,
		BiayaAdmin: pembayaran.BiayaAdmin,
		TotalBayar: pembayaran.TotalBayar,
		IDUser: pembayaran.IDUser,
	}
	return formatter
}