package main

import "fmt"

const nMAX int = 1000

type user struct {
	username    string
	password    string
	persetujuan bool
	role        string // Change to bool
}

type email struct {
	pengirim string
	penerima string
	subject  string
	pesan    string
}

type users [nMAX]user

var tab_password users
var tab_users users
var emails [nMAX]email
var jumlah_user int
var jumlah_email int

func main() {
	var pilih int
	for pilih != 4 {
		menu()
		fmt.Print("\nPilih (1,2,3,4,5)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			registrasi()
		case 2:
			PersetujuanAdmin()

		case 3:
			var username, password string
			var q bool
			fmt.Print("masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("masukkan password: ")
			fmt.Scan(&password)
			q = Login(tab_users, username, password, jumlah_user)
			for q == true {

				var pilih_2 int
				for pilih_2 != 6 {
					fmt.Println("-------------------------")
					fmt.Println("=                       =")
					fmt.Println("=         Home          =")
					fmt.Println("=                       =")
					fmt.Println("-------------------------")
					fmt.Println("1. Kirim pesan")
					fmt.Println("2. Baca pesan")
					fmt.Println("3. Balas Pesan")
					fmt.Println("4. Hapus Pesan")
					fmt.Println("5. Cetak Pesan")
					fmt.Println("6. Logout")
					fmt.Println("-------------------------")
					// menu()
					fmt.Print("\nPilih  (1,2,3,4,5,6)? ")
					fmt.Scan(&pilih_2)
					switch pilih_2 {
					case 1:
						kirimPesan()
					case 2:
						bacaPesan()
					case 3:
						balasPesan()
					case 4:
						hapusPesan()
					case 5:
						cetakPesan()
					case 6:
						q = false
					}
				}

			}
		case 4:
			fmt.Println("Keluar dari program.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menu() {
	fmt.Println("-------------------------")
	fmt.Println("=                       =")
	fmt.Println("=    Aplikasi  Email    =")
	fmt.Println("=                       =")
	fmt.Println("-------------------------")
	fmt.Println("1. Registrasi")
	fmt.Println("2. PersetujuanAdmin")
	fmt.Println("3. Login")
	fmt.Println("4. Exit")
	fmt.Println("-------------------------")
}

func registrasi() {
	if jumlah_user >= nMAX {
		fmt.Println("Tidak dapat menambah pengguna lagi.")
		return
	}
	var username, password string
	fmt.Print("\nMasukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	addUser(username, password)
	fmt.Println("Registrasi berhasil, menunggu persetujuan admin.")
}

func addUser(username, password string) {
	newUser := createUser(username, password, false)
	tab_users[jumlah_user] = newUser
	jumlah_user++
}
func createUser(username, password string, persetujuan bool) user {
	return user{
		username:    username,
		password:    password,
		persetujuan: persetujuan,
	}
}

func PersetujuanAdmin() {
	var username string
	fmt.Print("\nMasukkan username yang ingin disetujui/ditolak: ")
	fmt.Scan(&username)
	index := sequentialSearchUser(username)
	if index == -1 {
		fmt.Println("User tidak ditemukan.")
		return
	}
	var persetujuan int
	fmt.Print("Setujui pengguna? (1 untuk Ya, 0 untuk Tidak): ")
	fmt.Scan(&persetujuan)
	if persetujuan == 1 {
		tab_users[index].persetujuan = true
		fmt.Println("Pengguna disetujui.")
	} else {
		fmt.Println("Pengguna ditolak.")
	}
}
func Login(x users, username string, password string, n_User int) bool {

	index := sequentialSearchUser(username)
	if index == -1 {
		fmt.Println("Username salah.")
		return false
	}

	if x[index].password != password {
		fmt.Println("Password salah.")
		return false
	}

	if x[index].persetujuan != true {
		fmt.Println("Akun belum disetujui oleh admin.")
		return false
	}

	return true

}

func kirimPesan() {
	if jumlah_email >= nMAX {
		fmt.Println("Tidak dapat mengirim pesan lagi.")
		return
	}
	var sender, receiver, subject, message string
	fmt.Print("Pengirim: ")
	fmt.Scan(&sender)
	senderIndex := sequentialSearchUser(sender)
	if senderIndex == -1 || !tab_users[senderIndex].persetujuan {
		fmt.Println("Pengirim tidak valid atau belum disetujui.")
		return
	}

	fmt.Print("Penerima: ")
	fmt.Scan(&receiver)
	receiverIndex := sequentialSearchUser(receiver)
	if receiverIndex == -1 || !tab_users[receiverIndex].persetujuan {
		fmt.Println("Penerima tidak valid atau belum disetujui.")
		return
	}

	fmt.Print("Subjek: ")
	fmt.Scan(&subject)
	fmt.Print("Pesan: ")
	fmt.Scan(&message)

	emails[jumlah_email] = email{pengirim: sender, penerima: receiver, subject: subject, pesan: message}
	jumlah_email++
	fmt.Println("Pesan berhasil dikirim.")
}

// Fungsi untuk membaca pesan
func bacaPesan() {
	var username string
	fmt.Print("Masukkan username Anda: ")
	fmt.Scan(&username)
	userIndex := sequentialSearchUser(username)
	if userIndex == -1 || !tab_users[userIndex].persetujuan {
		fmt.Println("Username tidak valid atau belum disetujui.")
		return
	}

	fmt.Println("Daftar Pesan Masuk:")
	for i := 0; i < jumlah_email; i++ {
		if emails[i].penerima == username {
			fmt.Printf("Dari: %s\nSubjek: %s\nPesan: %s\n\n", emails[i].pengirim, emails[i].subject, emails[i].pesan)
		}
	}
}

// Fungsi untuk membalas pesan
func balasPesan() {
	var username, subjek string
	fmt.Print("Masukkan username Anda: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan subjek pesan yang ingin dibalas: ")
	fmt.Scan(&subjek)

	userIndex := sequentialSearchUser(username)
	if userIndex == -1 || !tab_users[userIndex].persetujuan {
		fmt.Println("Username tidak valid atau belum disetujui.")
		return
	}
	found := false
	for i := 0; i < jumlah_email; i++ {
		if emails[i].penerima == username && emails[i].subject == subjek {
			kirimPesan()
			found = true
		}
	}
	if !found {
		fmt.Println("Pesan dengan subjek tersebut tidak ditemukan.")
	}
}

// Fungsi untuk menghapus pesan
func hapusPesan() {
	var username, subjek string
	fmt.Print("Masukkan username Anda: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan subjek pesan yang ingin dihapus: ")
	fmt.Scan(&subjek)

	userIndex := sequentialSearchUser(username)
	if userIndex == -1 || !tab_users[userIndex].persetujuan {
		fmt.Println("Username tidak valid atau belum disetujui.")
		return
	}

	for i := 0; i < jumlah_email; i++ {
		if emails[i].penerima == username && emails[i].subject == subjek {
			// Menggeser semua email setelah index i ke kiri
			for j := i; j < jumlah_email-1; j++ {
				emails[j] = emails[j+1]
			}
			jumlah_email--
			fmt.Println("Pesan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pesan tidak ditemukan.")
}

// Fungsi untuk mencetak pesan
func cetakPesan() {
	var username string
	fmt.Print("Masukkan username Anda: ")
	fmt.Scan(&username)
	userIndex := sequentialSearchUser(username)
	if userIndex == -1 || !tab_users[userIndex].persetujuan {
		fmt.Println("Username tidak valid atau belum disetujui.")
		return
	}

	fmt.Println("Daftar Pesan Masuk:")
	for i := 0; i < jumlah_email; i++ {
		if emails[i].penerima == username {
			fmt.Printf("Dari: %s\nSubjek: %s\nPesan: %s\n\n", emails[i].pengirim, emails[i].subject, emails[i].pesan)
		}
	}
}

// Implement sequentialSearchUser
func sequentialSearchUser(username string) int {
	for i := 0; i < jumlah_user; i++ {
		if tab_users[i].username == username {
			return i
		}
	}
	return -1
}
func sequentialSearchpassword(password string) int {
	for i := 0; i < jumlah_user; i++ {
		if tab_password[i].password == password {
			return i
		}
	}
	return -1
}
