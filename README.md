# Golang
Berikut ada documentasi untuk menggunakan API ini 

## Nama Database
<b>pln</b>

## Start Server Golang
```bash
go run main.go
```
## List API metode [ POST ]
<li>{{server}}/user</li>
<li>{{server}}/login</li>
<li>{{serverprivate}}/create-penggunaan</li>
<li>{{serverprivate}}/create-tarif</li>
<li>{{serverprivate}}/create-tagihan</li>
<li>{{serverprivate}}/create-pelanggan</li>
<li>{{serverprivate}}/create-pembayaran</li>
<li>{{serverprivate}}/create-level</li>

## Request API Reguster
```bash
--header 'Content-Type: application/json' \
--data '{
    "id_user": "ADM0000",
    "username": "superAdmin",
    "password": "12345",
    "nama_admin": "superAdmin",
    "id_level": "LVLADMIN"
}'
```

## Request API Login
```bash
--header 'Content-Type: application/json' \
--data '{
    "username": "BAGAS",
    "password": "12345"
}'
```

## List API metode [ GET ]
<li>{{serverprivate}}/penggunaan</li>
<li>{{serverprivate}}/tarif</li>
<li>{{serverprivate}}/tagihan</li>
<li>{{serverprivate}}/pelanggan</li>
<li>{{serverprivate}}/pembayaran</li>
<li>{{serverprivate}}/level</li>

## GET API PENGGUNAAN
```bash
--header 'Authorization: ••••••' \
--data ''
```

## GET API TARIF
```bash
--header 'Authorization: ••••••'
```

## GET API PELANGGAN
```bash
--header 'Authorization: ••••••'
```

## GET API PEMBAYARAN
```bash
--header 'Authorization: ••••••'
```

## GET API TAGIHAN
```bash
--header 'Authorization: ••••••'
```

## GET API LEVEL
```bash
--header 'Authorization: ••••••'
```

## List API metode [ PATCH ]
<li>{{serverprivate}}/edit-penggunaan/:id_penggunaan</li>
<li>{{serverprivate}}/edit-tarif/:id_tarif</li>
<li>{{serverprivate}}/edit-tagihan/:id_tagihan</li>

## PATCH API PENGGUNAAN
```bash
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{        
    "bulan": 5,
    "tahun": 3000,
    "meter_awal": 100,
    "meter_akhir": 640
}'
```

## PATCH API TARIF
```bash
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{    
    "daya": "500VA",
    "tarif_perkwh": 25000
}'
```

## PATCH API TAGIHAN
```bash
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
    "status": "PAID"
}'
```

## List API metode [ DELETE ]
<li>{{serverprivate}}/delete-penggunaan/:id_penggunaan</li>
<li>{{serverprivate}}/delete-tari/:id_tarif</li>
<li>{{serverprivate}}/delete-tagihan/:id_tagihan</li>

## DELETE API PENGGUNAAN
```bash
--header 'Authorization: ••••••'
```

## PATCH API TARIF
```bash
--header 'Authorization: ••••••'
```

## PATCH API TAGIHAN
```bash
--header 'Authorization: ••••••'
```
