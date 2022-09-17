# assignment2go
#ini adalah tugas kedua di kelas golang, yaitu membuat rest API sederhana
rest API-nya menggunakan gorm, gin, dan juga mysql.
Untuk menggunakannya, bisa dengan aplikasi postman atau insomnia.
Lalu menulisakan alamat atau url http://localhost:8080/

Perintah yang bisa dilakukan, yaitu POST, PUT, GET, dan DELETE.
- untuk melakukan POST atau menginsert data baru, bisa dengan url http://localhost:8080/orders
lalu memasukkan data dengan syntax json pada bagian body, misalnya :
{
	"id":3,
	"name":"Mayoo Io",
	"ordered_at":"2020-09-17T18:38:45+00:00",
	"item":[{
		"order_id":3,
		"item_id":7,		
		"item_code": 130,
				"description": "iPhone 14",
				"quantity": 2
	}
	]
}
- untuk melakukan PUT atau mengupdate data, bisa dengan menambahkan id dibelakang url, misalnya untuk mengubah data dengan id 3, maka urlnya http://localhost:8080/orders/3
- untuk melakukan DELETE atau menghapus data bisa dengan menambahkan id dibelakang url, misalnya jika ingin menghapus data dengan id 4, maka urlnya http://localhost:8080/orders/4
- untuk melakukan GET atau melihat data, bisa dengan url http://localhost:8080/orders

