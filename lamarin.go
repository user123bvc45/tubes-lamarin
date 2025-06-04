package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type UserData struct {
	Nama         string
	Email        string
	Pendidikan   []Pendidikan
	Pengalaman   []Pengalaman
	Keterampilan []string
}

type Pendidikan struct {
	Institusi string
	Gelar     string
	Tahun     int
}

type Pengalaman struct {
	Posisi     string
	Perusahaan string
	Durasi     string
	Deskripsi  string
}

type Pekerjaan struct {
	Judul        string
	Perusahaan   string
	Lokasi       string
	Gaji         int
	Deskripsi    string
	Keterampilan []string
}

func main() {
	// Inisialisasi data pengguna
	user := UserData{}
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama dan email di awal
	fmt.Println("=== Selamat Datang di Aplikasi Pembuat Resume ===")
	fmt.Print("Masukkan nama Anda: ")
	if scanner.Scan() {
		user.Nama = scanner.Text()
	}

	fmt.Print("Masukkan alamat email: ")
	if scanner.Scan() {
		user.Email = scanner.Text()
	}

	// ==== Data dummy pekerjaan ====
	jobs := []Pekerjaan{
		{"Software Engineer", "Tech Corp", "Jakarta", 15000000, "Mengembangkan aplikasi berbasis web", []string{"Go", "JavaScript", "Docker"}},
		{"Data Scientist", "Data Analytics", "Bandung", 12000000, "Menganalisis data dan membuat model prediktif", []string{"Python", "Machine Learning", "SQL"}},
		{"Product Manager", "Digital Solutions", "Surabaya", 18000000, "Memimpin pengembangan produk digital", []string{"Leadership", "Agile", "Product Strategy"}},
		{"UI/UX Designer", "Creative Studio", "Yogyakarta", 10000000, "Mendesain tampilan dan pengalaman pengguna aplikasi mobile", []string{"Figma", "Adobe XD", "Creativity"}},
		{"Network Engineer", "Net Solutions", "Jakarta", 13000000, "Mengelola dan memelihara jaringan perusahaan", []string{"Networking", "Cisco", "Linux"}},
		{"Mobile Developer", "AppDev", "Bandung", 14000000, "Membuat aplikasi mobile Android/iOS", []string{"Kotlin", "Swift", "UI Design"}},
		{"System Analyst", "IT Consult", "Semarang", 12500000, "Menganalisis kebutuhan sistem dan membuat dokumentasi", []string{"Analysis", "Documentation", "SQL"}},
		{"Web Developer", "Webku id", "Yogyakarta", 11500000, "Membuat dan mengelola website perusahaan", []string{"HTML", "CSS", "JavaScript"}},
		{"Mobile App Tester", "Aplikasi Kita", "Yogyakarta", 9000000, "Menguji aplikasi Android/iOS sebelum rilis", []string{"Android", "iOS", "Testing"}},
		{"Helpdesk IT", "Universitas Maju", "Semarang", 8000000, "Menjawab pertanyaan dan membantu masalah IT kampus", []string{"Komunikasi", "Troubleshooting", "Sabar"}},
	}

	// Menu utama
	for {
		fmt.Println("\n=== Aplikasi Pembuat Resume dan Surat Lamaran ===")
		fmt.Println("1. Kelola Profil")
		fmt.Println("2. Cari Pekerjaan")
		fmt.Println("3. Buat Resume")
		fmt.Println("4. Evaluasi Resume")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			kelolaProfil(&user)
		case 2:
			cariPekerjaan(jobs)
		case 3:
			buatResume(user)
		case 4:
			evaluasiResume(user, jobs)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini! ")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid")
		}
	}
}

func kelolaProfil(user *UserData) {
	for {
		fmt.Println("\n=== Kelola Profil ===")
		fmt.Println("1. Tambah Pendidikan")
		fmt.Println("2. Tambah Pengalaman")
		fmt.Println("3. Tambah Keterampilan")
		fmt.Println("4. Ubah Data")
		fmt.Println("5. Hapus Data")
		fmt.Println("6. Lihat Profil")
		fmt.Println("7. Kembali ke menu awal")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPendidikan(user)
		case 2:
			tambahPengalaman(user)
		case 3:
			tambahKeterampilan(user)
		case 4:
			ubahData(user)
		case 5:
			hapusData(user)
		case 6:
			lihatProfil(*user)
		case 7:
			return
		default:
			fmt.Println("❌ Pilihan tidak valid")
		}
	}
}

func ubahData(user *UserData) {
	fmt.Println("\n=== Ubah Data ===")
	fmt.Println("1. Pendidikan")
	fmt.Println("2. Pengalaman")
	fmt.Println("3. Keterampilan")
	fmt.Print("Pilih jenis data yang ingin diubah: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		if len(user.Pendidikan) == 0 {
			fmt.Println("❌ Belum ada data pendidikan")
			return
		}
		for i, p := range user.Pendidikan {
			fmt.Printf("%d. %s, %s (%d)\n", i+1, p.Institusi, p.Gelar, p.Tahun)
		}
		fmt.Print("Pilih data yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Pendidikan) {
			fmt.Println("❌ Indeks tidak valid")
			return
		}
		fmt.Print("Institusi: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			user.Pendidikan[index-1].Institusi = scanner.Text()
		}
		fmt.Print("Gelar: ")
		if scanner.Scan() {
			user.Pendidikan[index-1].Gelar = scanner.Text()
		}
		fmt.Print("Tahun: ")
		fmt.Scanln(&user.Pendidikan[index-1].Tahun)
		fmt.Println("✅ Data pendidikan berhasil diubah!")
	case 2:
		if len(user.Pengalaman) == 0 {
			fmt.Println("❌ Belum ada data pengalaman")
			return
		}
		for i, exp := range user.Pengalaman {
			fmt.Printf("%d. %s di %s (%s)\n", i+1, exp.Posisi, exp.Perusahaan, exp.Durasi)
		}
		fmt.Print("Pilih data yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Pengalaman) {
			fmt.Println("❌ Indeks tidak valid")
			return
		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Posisi: ")
		if scanner.Scan() {
			user.Pengalaman[index-1].Posisi = scanner.Text()
		}
		fmt.Print("Perusahaan: ")
		if scanner.Scan() {
			user.Pengalaman[index-1].Perusahaan = scanner.Text()
		}
		fmt.Print("Durasi: ")
		if scanner.Scan() {
			user.Pengalaman[index-1].Durasi = scanner.Text()
		}
		fmt.Print("Deskripsi: ")
		if scanner.Scan() {
			user.Pengalaman[index-1].Deskripsi = scanner.Text()
		}
		fmt.Println("✅ Data pengalaman berhasil diubah!")
	case 3:
		if len(user.Keterampilan) == 0 {
			fmt.Println("❌ Belum ada keterampilan")
			return
		}
		for i, skill := range user.Keterampilan {
			fmt.Printf("%d. %s\n", i+1, skill)
		}
		fmt.Print("Pilih keterampilan yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Keterampilan) {
			fmt.Println("❌ Indeks tidak valid")
			return
		}
		fmt.Print("Masukkan keterampilan baru: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			user.Keterampilan[index-1] = scanner.Text()
		}
		fmt.Println("✅ Data keterampilan berhasil diubah!")
	default:
		fmt.Println("❌ Pilihan tidak valid")
	}
}

func hapusData(user *UserData) {
	fmt.Println("\n=== Hapus Data ===")
	fmt.Println("1. Pendidikan")
	fmt.Println("2. Pengalaman")
	fmt.Println("3. Keterampilan")
	fmt.Print("Pilih jenis data yang ingin dihapus: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		if len(user.Pendidikan) == 0 {
			fmt.Println("❌ Tidak ada data pendidikan untuk dihapus")
			return
		}
		for i, p := range user.Pendidikan {
			fmt.Printf("%d. %s, %s (%d)\n", i+1, p.Institusi, p.Gelar, p.Tahun)
		}
		fmt.Print("Pilih nomor yang akan dihapus: ")
		var index int
		fmt.Scanln(&index)
		if index >= 1 && index <= len(user.Pendidikan) {
			user.Pendidikan = append(user.Pendidikan[:index-1], user.Pendidikan[index:]...)
			fmt.Println("✅ Data pendidikan berhasil dihapus")
		} else {
			fmt.Println("❌ Indeks tidak valid")
		}

	} else if pilihan == 2 {
		if len(user.Pengalaman) == 0 {
			fmt.Println("❌ Tidak ada data pengalaman untuk dihapus")
			return
		}
		for i, p := range user.Pengalaman {
			fmt.Printf("%d. %s di %s (%s)\n", i+1, p.Posisi, p.Perusahaan, p.Durasi)
		}
		fmt.Print("Pilih nomor yang akan dihapus: ")
		var index int
		fmt.Scanln(&index)
		if index >= 1 && index <= len(user.Pengalaman) {
			user.Pengalaman = append(user.Pengalaman[:index-1], user.Pengalaman[index:]...)
			fmt.Println("✅ Data pengalaman berhasil dihapus")
		} else {
			fmt.Println("❌ Indeks tidak valid")
		}

	} else if pilihan == 3 {
		if len(user.Keterampilan) == 0 {
			fmt.Println("❌ Tidak ada keterampilan untuk dihapus")
			return
		}
		for i, k := range user.Keterampilan {
			fmt.Printf("%d. %s\n", i+1, k)
		}
		fmt.Print("Pilih nomor yang akan dihapus: ")
		var index int
		fmt.Scanln(&index)
		if index >= 1 && index <= len(user.Keterampilan) {
			user.Keterampilan = append(user.Keterampilan[:index-1], user.Keterampilan[index:]...)
			fmt.Println("✅ Keterampilan berhasil dihapus")
		} else {
			fmt.Println("❌ Indeks tidak valid")
		}
	} else {
		fmt.Println("❌ Indeks tidak valid")
	}
}

func tambahPendidikan(user *UserData) {
	fmt.Println("\n=== Tambah Pendidikan ===")
	var pendidikan Pendidikan

	fmt.Print("Institusi: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		pendidikan.Institusi = scanner.Text()
	}
	fmt.Print("Gelar: ")
	if scanner.Scan() {
		pendidikan.Gelar = scanner.Text()
	}
	fmt.Print("Tahun: ")
	fmt.Scanln(&pendidikan.Tahun)

	user.Pendidikan = append(user.Pendidikan, pendidikan)
	fmt.Println("✅ Pendidikan berhasil ditambahkan!")
}

func tambahPengalaman(user *UserData) {
	fmt.Println("\n=== Tambah Pengalaman ===")

	fmt.Print("Apakah Anda ingin menambahkan pengalaman kerja? (y/n): ")
	var jawaban string
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) != "y" {
		fmt.Println("Tidak menambahkan pengalaman kerja.")
		return
	}

	var pengalaman Pengalaman
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Posisi: ")
	if scanner.Scan() {
		pengalaman.Posisi = scanner.Text()
	}
	fmt.Print("Perusahaan: ")
	if scanner.Scan() {
		pengalaman.Perusahaan = scanner.Text()
	}
	fmt.Print("Durasi (contoh: 2018-2020): ")
	if scanner.Scan() {
		pengalaman.Durasi = scanner.Text()
	}
	fmt.Print("Deskripsi: ")
	if scanner.Scan() {
		pengalaman.Deskripsi = scanner.Text()
	}
	user.Pengalaman = append(user.Pengalaman, pengalaman)
	fmt.Println("✅ Pengalaman berhasil ditambahkan!")
}

func tambahKeterampilan(user *UserData) {
	fmt.Println("\n=== Tambah Keterampilan ===")
	fmt.Print("Masukkan keterampilan (pisahkan dengan koma jika lebih dari satu): ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		skills := strings.Split(input, ",")
		for i := range skills {
			skills[i] = strings.TrimSpace(skills[i])
		}
		user.Keterampilan = append(user.Keterampilan, skills...)
	}

	fmt.Println("✅ Keterampilan berhasil ditambahkan!")
}

func lihatProfil(user UserData) {
	fmt.Println("\n=== Profil Pengguna ===")
	fmt.Println("Nama:", user.Nama)
	fmt.Println("Email:", user.Email)

	fmt.Println("Pendidikan:")
	for i := 0; i < len(user.Pendidikan); i++ {
		fmt.Println(user.Pendidikan[i].Institusi, user.Pendidikan[i].Gelar, user.Pendidikan[i].Tahun)
	}
	fmt.Println("Pengalaman:")
	for i := 0; i < len(user.Pengalaman); i++ {
		fmt.Println(user.Pengalaman[i].Posisi, user.Pengalaman[i].Perusahaan, user.Pengalaman[i].Durasi, user.Pengalaman[i].Deskripsi)
	}
	fmt.Println("Keterampilan:")
	for i := 0; i < len(user.Keterampilan); i++ {
		fmt.Println(user.Keterampilan[i])
	}
}

func cariPekerjaan(jobs []Pekerjaan) {
	for {
		fmt.Println("\n=== Cari Pekerjaan ===")
		fmt.Println("1. Cari berdasarkan kata kunci (Sequential Search)")
		fmt.Println("2. Cari berdasarkan gaji (Binary Search)")
		fmt.Println("3. Urutkan berdasarkan relevansi (Selection Sort)")
		fmt.Println("4. Urutkan berdasarkan gaji (Insertion Sort)")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			sequentialSearch(jobs)
		case 2:
			binarySearch(jobs)
		case 3:
			selectionSortByRelevansi(jobs)
		case 4:
			insertionSortByGaji(jobs)
		case 5:
			return
		default:
			fmt.Println("❌ Pilihan tidak valid")
		}
	}
}

func sequentialSearch(jobs []Pekerjaan) {
	fmt.Print("Masukkan kata kunci pencarian: ")
	scanner := bufio.NewScanner(os.Stdin)
	var keyword string
	if scanner.Scan() {
		keyword = strings.ToLower(scanner.Text())
	}

	fmt.Println("\nHasil Pencarian:")
	found := false
	for _, job := range jobs {
		if strings.Contains(strings.ToLower(job.Judul), keyword) ||
			strings.Contains(strings.ToLower(job.Perusahaan), keyword) ||
			strings.Contains(strings.ToLower(job.Deskripsi), keyword) {
			tampilkanPekerjaan(job)
			found = true
		}
	}

	if !found {
		fmt.Println("❌ Tidak ditemukan pekerjaan dengan kata kunci tersebut")
	}
}

func binarySearch(jobs []Pekerjaan) {
	// Urutkan jobs berdasarkan gaji ASCENDING
	sortedJobs := make([]Pekerjaan, len(jobs))
	copy(sortedJobs, jobs)
	sort.Slice(sortedJobs, func(i, j int) bool {
		return sortedJobs[i].Gaji < sortedJobs[j].Gaji
	})

	fmt.Print("Masukkan gaji minimal yang diinginkan: ")
	var minGaji int
	fmt.Scanln(&minGaji)

	left := 0
	right := len(sortedJobs) - 1
	found := -1

	for left <= right && found == -1 {
		mid := (left + right) / 2
		if sortedJobs[mid].Gaji == minGaji {
			found = mid
		} else if sortedJobs[mid].Gaji > minGaji {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found != -1 {
		fmt.Println("\nPekerjaan ditemukan pada indeks:", found)
		tampilkanPekerjaan(sortedJobs[found])
	} else {
		fmt.Println("❌ Tidak ditemukan pekerjaan dengan gaji sesuai")
	}
}

func selectionSortByRelevansi(jobs []Pekerjaan) {
	sorted := make([]Pekerjaan, len(jobs))
	copy(sorted, jobs)

	for i := 0; i < len(sorted)-1; i++ {
		max := i
		for j := i + 1; j < len(sorted); j++ {
			if len(sorted[j].Keterampilan) > len(sorted[max].Keterampilan) {
				max = j
			}
		}
		temp := sorted[i]
		sorted[i] = sorted[max]
		sorted[max] = temp
	}

	fmt.Println("\nPekerjaan diurutkan berdasarkan relevansi:")
	for i := 0; i < len(sorted); i++ {
		tampilkanPekerjaan(sorted[i])
	}
}

func insertionSortByGaji(jobs []Pekerjaan) {
	sorted := make([]Pekerjaan, len(jobs))
	copy(sorted, jobs)

	for i := 1; i < len(sorted); i++ {
		temp := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j].Gaji > temp.Gaji {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = temp
	}

	fmt.Println("\nPekerjaan diurutkan berdasarkan gaji (terendah ke tertinggi):")
	for i := 0; i < len(sorted); i++ {
		tampilkanPekerjaan(sorted[i])
	}
}

func tampilkanPekerjaan(job Pekerjaan) {
	fmt.Printf("\n%s di %s (%s)\n", job.Judul, job.Perusahaan, job.Lokasi)
	fmt.Printf("Gaji: Rp%d\n", job.Gaji)
	fmt.Printf("Deskripsi: %s\n", job.Deskripsi)
	fmt.Printf("Keterampilan yang dibutuhkan: %s\n", strings.Join(job.Keterampilan, ", "))
}

func buatResume(user UserData) {
	fmt.Println("\n=== Resume Anda ===")
	fmt.Println("Nama:", user.Nama)
	fmt.Println("Email:", user.Email)

	fmt.Println("PENDIDIKAN")
	for i := 0; i < len(user.Pendidikan); i++ {
		fmt.Println(user.Pendidikan[i].Institusi, user.Pendidikan[i].Gelar, user.Pendidikan[i].Tahun)
	}
	fmt.Println("\nPENGALAMAN KERJA")
	for i := 0; i < len(user.Pengalaman); i++ {
		fmt.Println(user.Pengalaman[i].Posisi)
		fmt.Println(user.Pengalaman[i].Perusahaan, user.Pengalaman[i].Durasi)
		fmt.Println(user.Pengalaman[i].Deskripsi)
		fmt.Println()
	}
	fmt.Println("KEAHLIAN")
	for i := 0; i < len(user.Keterampilan); i++ {
		fmt.Println(user.Keterampilan[i])
	}
	fmt.Println("\n✅ Resume telah dibuat!")
}

func evaluasiResume(user UserData, jobs []Pekerjaan) {
	if len(jobs) == 0 {
		fmt.Println("❌ Tidak ada pekerjaan tersedia untuk evaluasi")
		return
	}

	fmt.Println("\n=== Evaluasi Resume ===")
	fmt.Println("Pilih pekerjaan untuk evaluasi:")
	for i, job := range jobs {
		fmt.Printf("%d. %s di %s\n", i+1, job.Judul, job.Perusahaan)
	}

	fmt.Print("Pilih nomor pekerjaan: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	if pilihan < 1 || pilihan > len(jobs) {
		fmt.Println("❌ Pilihan tidak valid")
		return
	}

	selectedJob := jobs[pilihan-1]
	skor := hitungKesesuaian(user, selectedJob)

	fmt.Printf("\nHasil Evaluasi untuk %s di %s:\n", selectedJob.Judul, selectedJob.Perusahaan)
	fmt.Printf("Skor Kesesuaian: %.1f%%\n", skor*100)

	// Analisis keterampilan
	fmt.Println("\nAnalisis Keterampilan:")
	skillMatch := 0
	for _, jobSkill := range selectedJob.Keterampilan {
		found := false
		for _, userSkill := range user.Keterampilan {
			if strings.EqualFold(jobSkill, userSkill) {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("- [✅] %s\n", jobSkill)
			skillMatch++
		} else {
			fmt.Printf("- [ ] %s (disarankan untuk ditambahkan)\n", jobSkill)
		}
	}

	// Saran perbaikan resume (simulasi AI sederhana)
	fmt.Println("\nSaran AI untuk meningkatkan resume:")
	if skor < 0.3 {
		fmt.Println("- Tambahkan lebih banyak pengalaman yang relevan")
		fmt.Println("- Pelajari keterampilan utama yang dibutuhkan untuk posisi ini")
		fmt.Println("- Ikuti pelatihan atau kursus online untuk meningkatkan skill")
		fmt.Println("- Perbaiki format dan tata letak resume agar lebih menarik")
	} else if skor < 0.7 {
		fmt.Println("- Soroti pengalaman yang paling relevan di bagian atas resume")
		fmt.Println("- Gunakan kata kunci dari deskripsi pekerjaan dalam resume Anda")
		fmt.Println("- Tampilkan pencapaian atau proyek yang pernah dikerjakan")
		fmt.Println("- Perbarui informasi kontak dan portofolio")
	} else {
		fmt.Println("- Resume Anda sudah cukup kuat untuk posisi ini")
		fmt.Println("- Fokus pada penulisan surat lamaran yang menarik")
		fmt.Println("- Siapkan diri untuk wawancara kerja")
	}
}

func hitungKesesuaian(user UserData, job Pekerjaan) float64 {
	matchedSkills := 0
	for i := 0; i < len(job.Keterampilan); i++ {
		for j := 0; j < len(user.Keterampilan); j++ {
			if strings.EqualFold(job.Keterampilan[i], user.Keterampilan[j]) {
				matchedSkills++
				break
			}
		}
	}
	var skillScore float64 = 0
	if len(job.Keterampilan) > 0 {
		skillScore = float64(matchedSkills) / float64(len(job.Keterampilan))
	}
	var expScore float64 = 0
	if len(user.Pengalaman) > 0 {
		expScore = 1
	}
	return (skillScore + expScore) / 2
}
