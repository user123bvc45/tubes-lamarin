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

	// Data pekerjaan
	jobs := []Pekerjaan{
		{"Software k", "Tech Corp", "Jakarta", 15000000, "Mengembangkan aplikasi berbasis web", []string{"Go", "JavaScript", "Docker"}},
		{"Data Scientist", "Data Analytics", "Bandung", 12000000, "Menganalisis data dan membuat model prediktif", []string{"Python", "Machine Learning", "SQL"}},
		{"Product Manager", "Digital Solutions", "Surabaya", 18000000, "Memimpin pengembangan produk digital", []string{"Leadership", "Agile", "Product Strategy"}},
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
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid")
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
		fmt.Println("7. Kembali")
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
			fmt.Println("Pilihan tidak valid")
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
			fmt.Println("Belum ada data pendidikan")
			return
		}
		for i, p := range user.Pendidikan {
			fmt.Printf("%d. %s, %s (%d)\n", i+1, p.Institusi, p.Gelar, p.Tahun)
		}
		fmt.Print("Pilih data yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Pendidikan) {
			fmt.Println("Indeks tidak valid")
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
	case 2:
		if len(user.Pengalaman) == 0 {
			fmt.Println("Belum ada data pengalaman")
			return
		}
		for i, exp := range user.Pengalaman {
			fmt.Printf("%d. %s di %s (%s)\n", i+1, exp.Posisi, exp.Perusahaan, exp.Durasi)
		}
		fmt.Print("Pilih data yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Pengalaman) {
			fmt.Println("Indeks tidak valid")
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
	case 3:
		if len(user.Keterampilan) == 0 {
			fmt.Println("Belum ada keterampilan")
			return
		}
		for i, skill := range user.Keterampilan {
			fmt.Printf("%d. %s\n", i+1, skill)
		}
		fmt.Print("Pilih keterampilan yang ingin diubah (nomor): ")
		var index int
		fmt.Scanln(&index)
		if index < 1 || index > len(user.Keterampilan) {
			fmt.Println("Indeks tidak valid")
			return
		}
		fmt.Print("Masukkan keterampilan baru: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			user.Keterampilan[index-1] = scanner.Text()
		}
	default:
		fmt.Println("Pilihan tidak valid")
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

	switch pilihan {
	case 1:
		if len(user.Pendidikan) == 0 {
			fmt.Println("Tidak ada data pendidikan untuk dihapus")
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
			fmt.Println("Data pendidikan berhasil dihapus")
		} else {
			fmt.Println("Indeks tidak valid")
		}
	case 2:
		if len(user.Pengalaman) == 0 {
			fmt.Println("Tidak ada data pengalaman untuk dihapus")
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
			fmt.Println("Data pengalaman berhasil dihapus")
		} else {
			fmt.Println("Indeks tidak valid")
		}
	case 3:
		if len(user.Keterampilan) == 0 {
			fmt.Println("Tidak ada keterampilan untuk dihapus")
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
			fmt.Println("Keterampilan berhasil dihapus")
		} else {
			fmt.Println("Indeks tidak valid")
		}
	default:
		fmt.Println("Pilihan tidak valid")
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
	fmt.Println("Pendidikan berhasil ditambahkan!")
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
	fmt.Println("Pengalaman berhasil ditambahkan!")
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

	fmt.Println("Keterampilan berhasil ditambahkan!")
}

func lihatProfil(user UserData) {
	fmt.Println("\n=== Profil Pengguna ===")
	fmt.Printf("Nama: %s\n", user.Nama)
	fmt.Printf("Email: %s\n", user.Email)

	fmt.Println("\nPendidikan:")
	for _, p := range user.Pendidikan {
		fmt.Printf("- %s, %s (%d)\n", p.Institusi, p.Gelar, p.Tahun)
	}

	fmt.Println("\nPengalaman:")
	for _, exp := range user.Pengalaman {
		fmt.Printf("- %s di %s (%s)\n", exp.Posisi, exp.Perusahaan, exp.Durasi)
		fmt.Printf("  Deskripsi: %s\n", exp.Deskripsi)
	}

	fmt.Println("\nKeterampilan:")
	for _, skill := range user.Keterampilan {
		fmt.Printf("- %s\n", skill)
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
			fmt.Println("Pilihan tidak valid")
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
		fmt.Println("Tidak ditemukan pekerjaan dengan kata kunci tersebut")
	}
}

func binarySearch(jobs []Pekerjaan) {
	// Untuk binary search, kita perlu mengurutkan terlebih dahulu
	sortedJobs := make([]Pekerjaan, len(jobs))
	copy(sortedJobs, jobs)
	sort.Slice(sortedJobs, func(i, j int) bool {
		return sortedJobs[i].Gaji < sortedJobs[j].Gaji
	})

	fmt.Print("Masukkan gaji minimal yang diinginkan: ")
	var minGaji int
	fmt.Scanln(&minGaji)

	// Implementasi binary search
	low := 0
	high := len(sortedJobs) - 1
	var result []Pekerjaan

	for low <= high {
		mid := (low + high) / 2
		if sortedJobs[mid].Gaji >= minGaji {
			// Tambahkan semua pekerjaan dengan gaji >= minGaji
			result = append(result, sortedJobs[mid:]...)
			break
		} else {
			low = mid + 1
		}
	}

	if len(result) > 0 {
		fmt.Println("\nPekerjaan dengan gaji sesuai atau lebih tinggi:")
		for _, job := range result {
			tampilkanPekerjaan(job)
		}
	} else {
		fmt.Println("Tidak ditemukan pekerjaan dengan gaji sesuai")
	}
}

func selectionSortByRelevansi(jobs []Pekerjaan) {
	// Untuk demo, kita anggap relevansi adalah jumlah keterampilan yang sesuai
	sortedJobs := make([]Pekerjaan, len(jobs))
	copy(sortedJobs, jobs)

	n := len(sortedJobs)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			// Bandingkan berdasarkan jumlah keterampilan (semakin banyak semakin relevan)
			if len(sortedJobs[j].Keterampilan) > len(sortedJobs[minIdx].Keterampilan) {
				minIdx = j
			}
		}
		sortedJobs[i], sortedJobs[minIdx] = sortedJobs[minIdx], sortedJobs[i]
	}

	fmt.Println("\nPekerjaan diurutkan berdasarkan relevansi:")
	for _, job := range sortedJobs {
		tampilkanPekerjaan(job)
	}
}

func insertionSortByGaji(jobs []Pekerjaan) {
	sortedJobs := make([]Pekerjaan, len(jobs))
	copy(sortedJobs, jobs)

	for i := 1; i < len(sortedJobs); i++ {
		key := sortedJobs[i]
		j := i - 1

		for j >= 0 && sortedJobs[j].Gaji > key.Gaji {
			sortedJobs[j+1] = sortedJobs[j]
			j = j - 1
		}
		sortedJobs[j+1] = key
	}

	fmt.Println("\nPekerjaan diurutkan berdasarkan gaji (terendah ke tertinggi):")
	for _, job := range sortedJobs {
		tampilkanPekerjaan(job)
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
	fmt.Printf("Nama: %s\n", user.Nama)
	fmt.Printf("Email: %s\n\n", user.Email)

	fmt.Println("PENDIDIKAN")
	for _, p := range user.Pendidikan {
		fmt.Printf("%s - %s (%d)\n", p.Institusi, p.Gelar, p.Tahun)
	}

	fmt.Println("\nPENGALAMAN KERJA")
	for _, exp := range user.Pengalaman {
		fmt.Printf("%s\n%s, %s\n%s\n\n", exp.Posisi, exp.Perusahaan, exp.Durasi, exp.Deskripsi)
	}

	fmt.Println("KEAHLIAN")
	fmt.Println(strings.Join(user.Keterampilan, ", "))

	fmt.Println("\nResume telah dibuat! (Dalam aplikasi nyata, ini akan diekspor ke PDF atau format lain)")
}

func evaluasiResume(user UserData, jobs []Pekerjaan) {
	if len(jobs) == 0 {
		fmt.Println("Tidak ada pekerjaan tersedia untuk evaluasi")
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
		fmt.Println("Pilihan tidak valid")
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
			fmt.Printf("- [✓] %s\n", jobSkill)
			skillMatch++
		} else {
			fmt.Printf("- [ ] %s (disarankan untuk ditambahkan)\n", jobSkill)
		}
	}

	// Saran AI (simulasi)
	fmt.Println("\nSaran AI untuk meningkatkan resume:")
	if skor < 0.3 {
		fmt.Println("- Tambahkan lebih banyak pengalaman yang relevan")
		fmt.Println("- Pelajari keterampilan utama yang dibutuhkan untuk posisi ini")
	} else if skor < 0.7 {
		fmt.Println("- Soroti pengalaman yang paling relevan di bagian atas resume")
		fmt.Println("- Gunakan kata kunci dari deskripsi pekerjaan dalam resume Anda")
	} else {
		fmt.Println("- Resume Anda sudah cukup kuat untuk posisi ini")
		fmt.Println("- Fokus pada penulisan surat lamaran yang menarik")
	}
}

func hitungKesesuaian(user UserData, job Pekerjaan) float64 {
	// Hitung kesesuaian keterampilan
	skillScore := 0.0
	if len(job.Keterampilan) > 0 {
		matchedSkills := 0
		for _, jobSkill := range job.Keterampilan {
			for _, userSkill := range user.Keterampilan {
				if strings.EqualFold(jobSkill, userSkill) {
					matchedSkills++
					break
				}
			}
		}
		skillScore = float64(matchedSkills) / float64(len(job.Keterampilan))
	}

	// Hitung kesesuaian pengalaman (sederhana)
	expScore := 0.0
	if len(user.Pengalaman) > 0 {
		// Asumsikan beberapa relevansi jika ada pengalaman
		expScore = 0.3
		// Periksa jika ada kata kunci yang cocok di pengalaman
		for _, exp := range user.Pengalaman {
			if strings.Contains(strings.ToLower(exp.Posisi), strings.ToLower(job.Judul)) {
				expScore = 0.7
				break
			}
		}
	}

	// Gabungkan skor (bobot: 60% keterampilan, 40% pengalaman)
	return 0.6*skillScore + 0.4*expScore
}
