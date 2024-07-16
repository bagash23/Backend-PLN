package pembayaran

type PembayaranInput struct {
	IDPembayaran string  `json:"id_pembayaran" binding:"required"`
	IDTagihan string `json:"id_tagihan" binding:"required"`
	IDPelanggan string `json:"id_pelanggan" binding:"required"`
	TglBayar string `json:"tgl_bayar" binding:"required"`
	BiayaAdmin int `json:"biaya_admin" binding:"required"`
	TotalBayar int `json:"total_bayar" binding:"required"`
	IDUser string `json:"id_user" binding:"required"`
}