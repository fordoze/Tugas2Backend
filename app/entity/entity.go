package entity

type SuperAdmin struct {
	IdSuperAdmin uint   `gorm:"primaryKey;autoIncrement" json:"id_super_admin"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	CreatedAt    string `json:"created_at"`
}

type Admin struct {
	IdAdmin          uint   `gorm:"primaryKey;autoIncrement" json:"id_admin"`
	IdPerpustakaan   uint   `json:"id_perpustakaan"`
	Nama             string `json:"nama"`
	NamaPerpustakaan string `json:"nama_perpustakaan"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	CreatedAt        string `json:"created_at"`
}

type Perpustakaan struct {
	IdPerpustakaan uint    `gorm:"primaryKey;autoIncrement" json:"id_perpustakaan"`
	Nama           string  `json:"nama"`
	Lokasi         string  `json:"lokasi"`
	Kota           string  `json:"kota"`
	Provinsi       string  `json:"provinsi"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Admin          []Admin `gorm:"foreignKey:IdPerpustakaan" json:"admin"`
	Buku           []Buku  `gorm:"foreignKey:IdPerpustakaan" json:"buku"`
}

type Buku struct {
	IdBuku         uint   `gorm:"primaryKey;autoIncrement" json:"id_buku"`
	IdPerpustakaan uint   `json:"id_perpustakaan"`
	Judul          string `json:"judul"`
	Halaman        string `json:"halaman"`
	Penerbit       string `json:"penerbit"`
	Pengarang      string `json:"pengarang"`
	Tahun          string `json:"tahun"`
	Kategori       string `json:"kategori"`
	CreatedAt      string `json:"created_at"`
}

type Users struct {
	IdUsers   uint        `gorm:"primaryKey;autoIncrement" json:"id_users"`
	Name      string      `json:"name"`
	Jurusan   string      `json:"jurusan"`
	Fakultas  string      `json:"fakultas"`
	Kampus    string      `json:"kampus"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	UsersBook []UsersBook `gorm:"foreignKey:IdUsers" json:"users_book"`
}

type UsersBook struct {
	IdUsersBook    uint   `gorm:"primaryKey;autoIncrement" json:"id_users_book"`
	IdBuku         uint   `json:"id_buku"`
	IdUsers        uint   `json:"id_users"`
	Durasi         string `json:"durasi"`
	TangggalPinjam string `json:"tanggal_pinjam"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
