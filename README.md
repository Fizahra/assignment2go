# assignment2go
#ini adalah tugas kedua di kelas golang, yaitu membuat rest API sederhana
rest API-nya menggunakan gorm, gin, dan juga mysql.

##Sebelum menjalankan rest API-nya, file env.example harus direname dulu menjadi .env

###Untuk menggunakannya, bisa dengan aplikasi postman atau insomnia.
Lalu menulisakan alamat atau url http://localhost:8080/

#Sebelum melakukan perintah CRUD, diharuskan untuk melakukan POST register dan login pada url http://localhost:8080/auth/register atau http://localhost:8080/auth/login di postman atau insomnia
-Saat melakukan register, diperlukan nama, email dan password yang harus ditulis dengan syntax json pada bagian body. Misalnya :
{
	"name" : "Fii",
	"email" : "fiimashima3@gmail.com",
	"password" : "mashima"
}
- Untuk melakukan login cukup dengan menuliskan email dan password. Lalu token pada bagian hasilnya harus di-copy untuk melakukan perintah CRUD

#Untuk melakukan perintah pada user, seperti melihat dan mengedit user bisa dengan melakukan GET dan PUT pada url http://localhost:8080/user

#Perintah yang bisa dilakukan, yaitu POST, PUT, GET, dan DELETE. 
Tetapi sebelum itu, pada bagian Header wajib ditambahkan "Authorization" dan memasukkan token yang telah didapat dari login tadi agar perintahnya berjalan.
- untuk melakukan POST atau menginsert data baru, bisa dengan url http://localhost:8080/orders
lalu memasukkan data dengan syntax json pada bagian body, misalnya :
{
	
	"name":"Luthfia H",
	"ordered_at":"2020-09-24T15:02:45+00:00",
	"item":[{	
		"item_code": 131,
				"description": "Poco X",
				"quantity": 2
	}
	],
	"user_id" : 8
}
- untuk melakukan PUT atau mengupdate data, bisa dengan menambahkan id dibelakang url, misalnya untuk mengubah data dengan id 3, maka urlnya http://localhost:8080/orders/3
- untuk melakukan DELETE atau menghapus data bisa dengan menambahkan id dibelakang url, misalnya jika ingin menghapus data dengan id 4, maka urlnya http://localhost:8080/orders/4
- untuk melakukan GET atau melihat data, bisa dengan url http://localhost:8080/orders

