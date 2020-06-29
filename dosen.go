package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type MDosen struct {
	ID             int    `json:"ID"`
	Login          string `json:"Login"`
	KodeID         string `json:"KodeID"`
	IDFINGER       string `json:"IDFInger"`
	NIDN           string `json:"NIDN"`
	HomebaseInduk  string `json:"HomebaseInduk"`
	NIPPNS         string `json:"NIPPNS"`
	NPWP           string `json:"NPWP"`
	Nama           string `json:"Nama"`
	Struktural     string `json:"Struktural"`
	Jenis_Kelamin  string `json:"Jenis_Kelamin"`
	TempatLahir    string `json:"TempatLahir"`
	TanggalLahir   string `json:"TanggalLahir"`
	LevelID        int    `json:"LevelID"`
	KTP            string `json:"KTP"`
	Telephone      string `json:"Telephone"`
	Password       string `json:"Password"`
	Handphone      string `json:"Handphone"`
	Email          string `json:"Email"`
	Alamat         string `json:"Alamat"`
	Kota           string `json:"Kota"`
	KodePos        string `json:"KodePos"`
	Propinsi       string `json:"Propinsi"`
	Negara         string `json:"Negara"`
	NA             string `json:"NA"`
	Homebase       string `json:"Homebase"`
	ProdiID        string `json:"ProdiID"`
	Gelar          string `json:"Gelar"`
	JenjangID      string `json:"JenjangID"`
	Keilmuan       string `json:"Keilmuan"`
	LulusanPT      string `json:"LulusanPT"`
	AgamaID        string `json:"AgamaID"`
	KelaminID      string `json:"KelaminID"`
	GolonganID     int    `json:"GolonganID"`
	FotoKTP        string `json:"FotoKTP"`
	Foto           string `json:"foto"`
	KategoriID     string `json:"KategoriID"`
	IkatanID       string `json:"IkatanID"`
	JabatanID      string `json:"JabatanID"`
	JabatanDiktiID string `json:"JabatanDiktiID"`
	InstitusiInduk string `json:"InstitusiInduk"`
	StatusDosenID  string `json:"StatusDosenID"`
	StatusKerjaID  string `json:"StatusKerjaID"`
	TglBekerja     string `json:"TglBekerja"`
	NamaBank       string `json:"NamaBank"`
	NamaAkun       string `json:"NamaAkun"`
	NomerAkun      string `json:"NomerAkun"`
	LoginBuat      string `json:"LoginBuat"`
	TanggalBuat    string `json:"TanggalBuat"`
	LoginEdit      string `json:"LoginEdit"`
	TanggalEdit    string `json:"TanggalEdit"`
	TransReg       int    `json:"TransReg"`
	TransKhus      int    `json:"TransKhus"`
	SKSKhus        int    `json:"SksKhus"`
	SKSReg         int    `json:"SksReg"`
}

// Get all orders

func getMDosen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m_dosen []MDosen

	sql := `SELECT
				ID, 		
				IFNULL(Login,''), 		
				IFNULL(KodeID,'')KodeID, 		
				IFNULL(IDFINGER,'')IDFINGER, 	
				IFNULL(NIDN,'')NIDN, 		
				IFNULL(HomebaseInduk,'')HomebaseInduk, 
				IFNULL(NIPPNS,'')NIPPNS, 		
				IFNULL(NPWP,'')NPWP, 		
				IFNULL(Nama,'')Nama, 		
				IFNULL(Struktural,'')Struktural, 	
				IFNULL(Jenis_Kelamin,'')Jenis_Kelamin, 
				IFNULL(TempatLahir,'')TempatLahir, 
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(LevelID,'')LevelID, 	
				IFNULL(KTP,'')KTP, 		
				IFNULL(Telephone,'')Telephone, 	
				IFNULL(Password,'')Password, 	
				IFNULL(Handphone,'')Handphone,	 
				IFNULL(Email,'')Email, 		
				IFNULL(Alamat,'')Alamat, 		
				IFNULL(Kota,'')Kota, 		
				IFNULL(KodePos,'')KodePos, 	
				IFNULL(Propinsi,'')Propinsi,	
				IFNULL(Negara,'')Negara, 		
				IFNULL(NA,'')NA, 			
				IFNULL(Homebase,'')Homebase, 	
				IFNULL(ProdiID,'')ProdiID, 	
				IFNULL(Gelar,'')Gelar, 		
				IFNULL(JenjangID,'')JenjangID,	
				IFNULL(Keilmuan,'')Keilmuan, 	
				IFNULL(LulusanPT,'')LulusanPT, 	
				IFNULL(AgamaID,'')AgamaID, 	
				IFNULL(KelaminID,'')KelaminID, 	
				IFNULL(GolonganID,'')GolonganID, 	
				IFNULL(FotoKTP,'')FotoKTP, 	
				IFNULL(Foto,'')Foto ,		
				IFNULL(KategoriID,'')KategoriID, 	
				IFNULL(IkatanID,'')IkatanID, 	
				IFNULL(JabatanID,'')JabatanID, 	
				IFNULL(JabatanDiktiID,'')JabatanDiktiID,
				IFNULL(InstitusiInduk,'')InstitusiInduk,
				IFNULL(StatusDosenID,'')StatusDosenID,
				IFNULL(StatusKerjaID,'')StatusKerjaID,
				IFNULL(TglBekerja,'')TglBekerja,  
				IFNULL(NamaBank,'')NamaBank,     
				IFNULL(NamaAkun,'')NamaAkun,     
				IFNULL(NomerAkun,'')NomerAkun,    
				IFNULL(LoginBuat,'')LoginBuat,    
				IFNULL(TanggalBuat,'')TanggalBuat,  
				IFNULL(LoginEdit,'')LoginEdit,    
				IFNULL(TanggalEdit,'')TanggalEdit,  
				IFNULL(TransReg,'')TransReg,     
				IFNULL(TransKhus,'')TransKhus,    
				IFNULL(SKSKhus,'')SKSKhus,      
				IFNULL(SKSReg,'')SKSReg       
			FROM m_dosen`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var mDosen MDosen
		err := result.Scan(&mDosen.ID, &mDosen.Login, &mDosen.KodeID, &mDosen.IDFINGER, &mDosen.NIDN,
			&mDosen.HomebaseInduk, &mDosen.NIPPNS, &mDosen.NPWP,
			&mDosen.Nama, &mDosen.Struktural, &mDosen.Jenis_Kelamin, &mDosen.TempatLahir,
			&mDosen.TanggalLahir, &mDosen.LevelID, &mDosen.KTP, &mDosen.Telephone, &mDosen.Password, &mDosen.Handphone,
			&mDosen.Email, &mDosen.Alamat, &mDosen.Kota, &mDosen.KodePos, &mDosen.Propinsi, &mDosen.Negara,
			&mDosen.NA, &mDosen.Homebase, &mDosen.ProdiID, &mDosen.Gelar, &mDosen.JenjangID, &mDosen.Keilmuan,
			&mDosen.LulusanPT, &mDosen.AgamaID, &mDosen.KelaminID, &mDosen.GolonganID, &mDosen.FotoKTP, &mDosen.Foto,
			&mDosen.KategoriID, &mDosen.IkatanID, &mDosen.JabatanID, &mDosen.JabatanDiktiID,
			&mDosen.InstitusiInduk, &mDosen.StatusDosenID, &mDosen.StatusKerjaID, &mDosen.TglBekerja, &mDosen.NamaBank,
			&mDosen.NamaAkun, &mDosen.NomerAkun, &mDosen.LoginBuat, &mDosen.TanggalBuat, &mDosen.LoginEdit,
			&mDosen.TanggalEdit, &mDosen.TransReg, &mDosen.TransKhus, &mDosen.SKSKhus, &mDosen.SKSReg)

		if err != nil {
			panic(err.Error())
		}
		m_dosen = append(m_dosen, mDosen)
	}

	json.NewEncoder(w).Encode(m_dosen)
}

func createMDosen(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		ID := r.FormValue("id")
		Login := r.FormValue("login")
		KodeID := r.FormValue("kodeID")
		IDFINGER := r.FormValue("idFinger")
		NIDN := r.FormValue("nidn")
		HomebaseInduk := r.FormValue("homebaseInduk")
		NIPPNS := r.FormValue("nippns")
		NPWP := r.FormValue("npwp")
		Nama := r.FormValue("nama")
		Struktural := r.FormValue("struktural")
		Jenis_Kelamin := r.FormValue("jenisKelamin")
		TempatLahir := r.FormValue("tempatLahir")
		TanggalLahir := r.FormValue("tanggalLahir")
		LevelID := r.FormValue("levelID")
		KTP := r.FormValue("ktp")
		Telephone := r.FormValue("telephone")
		Password := r.FormValue("password")
		Handphone := r.FormValue("handphone")
		Email := r.FormValue("email")
		Alamat := r.FormValue("alamat")
		Kota := r.FormValue("kota")
		KodePos := r.FormValue("kodePos")
		Propinsi := r.FormValue("propinsi")
		Negara := r.FormValue("negara")
		NA := r.FormValue("na")
		Homebase := r.FormValue("homebase")
		ProdiID := r.FormValue("prodiID")
		Gelar := r.FormValue("gelar")
		JenjangID := r.FormValue("jenjangID")
		Keilmuan := r.FormValue("keilmuan")
		LulusanPT := r.FormValue("lulusanPT")
		AgamaID := r.FormValue("agamaID")
		KelaminID := r.FormValue("kelaminID")
		GolonganID := r.FormValue("golonganID")
		FotoKTP := r.FormValue("fotoKTP")
		Foto := r.FormValue("foto")
		KategoriID := r.FormValue("kategoriID")
		IkatanID := r.FormValue("ikatanID")
		JabatanID := r.FormValue("jabatanID")
		JabatanDiktiID := r.FormValue("jabatanDiktiID")
		InstitusiInduk := r.FormValue("institusiInduk")
		StatusDosenID := r.FormValue("statusDosenID")
		StatusKerjaID := r.FormValue("statusKerjaID")
		TglBekerja := r.FormValue("tglBekerja")
		NamaBank := r.FormValue("namaBank")
		NamaAkun := r.FormValue("namaAkun")
		NomerAkun := r.FormValue("nomerAkun")
		LoginBuat := r.FormValue("loginBuat")
		TanggalBuat := r.FormValue("tanggalBuat")
		LoginEdit := r.FormValue("loginEdit")
		TanggalEdit := r.FormValue("tanggalEdit")
		TransReg := r.FormValue("transReg")
		TransKhus := r.FormValue("transKhus")
		SKSKhus := r.FormValue("sksKhus")
		SKSReg := r.FormValue("sksReg")

		stmt, err := db.Prepare("INSERT INTO m_dosen (ID, Login, KodeID, IDFINGER, NIDN, HomebaseInduk, NIPPNS, NPWP, Nama, Struktural, Jenis_Kelamin, TempatLahir, TanggalLahir, LevelID, KTP, Telephone,Password, Handphone, Email, Alamat, Kota, KodePos, Propinsi, Negara, NA, Homebase, ProdiID, Gelar, JenjangID,Keilmuan, LulusanPT, AgamaID, KelaminID, GolonganID, FotoKTP, Foto, KategoriID, IkatanID, JabatanID, JabatanDiktiID, InstitusiInduk, StatusDosenID, StatusKerjaID, TglBekerja, NamaBank, NamaAkun, NomerAkun, LoginBuat, TanggalBuat, LoginEdit, TanggalEdit,TransReg, TransKhus, SKSKhus, SKSReg) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(ID, Login, KodeID, IDFINGER, NIDN, HomebaseInduk, NIPPNS, NPWP, Nama, Struktural, Jenis_Kelamin, TempatLahir, TanggalLahir, LevelID, KTP, Telephone, Password, Handphone, Email, Alamat, Kota, KodePos, Propinsi, Negara, NA, Homebase, ProdiID, Gelar, JenjangID, Keilmuan, LulusanPT, AgamaID, KelaminID, GolonganID, FotoKTP, Foto, KategoriID, IkatanID, JabatanID, JabatanDiktiID, InstitusiInduk, StatusDosenID, StatusKerjaID, TglBekerja, NamaBank, NamaAkun, NomerAkun, LoginBuat, TanggalBuat, LoginEdit, TanggalEdit, TransReg, TransKhus, SKSKhus, SKSReg)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {

			fmt.Fprintf(w, "Date Created")
			//http.Redirect(w, r, "/", 301)
		}

	}
}

func getMDosenByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m_dosen []MDosen
	params := mux.Vars(r)

	sql := `SELECT
				ID, 		
				IFNULL(Login,''), 		
				IFNULL(KodeID,'')KodeID, 		
				IFNULL(IDFINGER,'')IDFINGER, 	
				IFNULL(NIDN,'')NIDN, 		
				IFNULL(HomebaseInduk,'')HomebaseInduk, 
				IFNULL(NIPPNS,'')NIPPNS, 		
				IFNULL(NPWP,'')NPWP, 		
				IFNULL(Nama,'')Nama, 		
				IFNULL(Struktural,'')Struktural, 	
				IFNULL(Jenis_Kelamin,'')Jenis_Kelamin, 
				IFNULL(TempatLahir,'')TempatLahir, 
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(LevelID,'')LevelID, 	
				IFNULL(KTP,'')KTP, 		
				IFNULL(Telephone,'')Telephone, 	
				IFNULL(Password,'')Password, 	
				IFNULL(Handphone,'')Handphone,	 
				IFNULL(Email,'')Email, 		
				IFNULL(Alamat,'')Alamat, 		
				IFNULL(Kota,'')Kota, 		
				IFNULL(KodePos,'')KodePos, 	
				IFNULL(Propinsi,'')Propinsi,	
				IFNULL(Negara,'')Negara, 		
				IFNULL(NA,'')NA, 			
				IFNULL(Homebase,'')Homebase, 	
				IFNULL(ProdiID,'')ProdiID, 	
				IFNULL(Gelar,'')Gelar, 		
				IFNULL(JenjangID,'')JenjangID,	
				IFNULL(Keilmuan,'')Keilmuan, 	
				IFNULL(LulusanPT,'')LulusanPT, 	
				IFNULL(AgamaID,'')AgamaID, 	
				IFNULL(KelaminID,'')KelaminID, 	
				IFNULL(GolonganID,'')GolonganID, 	
				IFNULL(FotoKTP,'')FotoKTP, 	
				IFNULL(Foto,'')Foto ,		
				IFNULL(KategoriID,'')KategoriID, 	
				IFNULL(IkatanID,'')IkatanID, 	
				IFNULL(JabatanID,'')JabatanID, 	
				IFNULL(JabatanDiktiID,'')JabatanDiktiID,
				IFNULL(InstitusiInduk,'')InstitusiInduk,
				IFNULL(StatusDosenID,'')StatusDosenID,
				IFNULL(TglBekerja,'')TglBekerja,  
				IFNULL(NamaBank,'')NamaBank,     
				IFNULL(NamaAkun,'')NamaAkun,     
				IFNULL(NomerAkun,'')NomerAkun,    
				IFNULL(LoginBuat,'')LoginBuat,    
				IFNULL(TanggalBuat,'')TanggalBuat,  
				IFNULL(LoginEdit,'')LoginEdit,    
				IFNULL(TanggalEdit,'')TanggalEdit,  
				IFNULL(TransReg,'')TransReg,     
				IFNULL(TransKhus,'')TransKhus,    
				IFNULL(SKSKhus,'')SKSKhus,      
				IFNULL(SKSReg,'')SKSReg  
			FROM m_dosen WHERE ID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mDosen MDosen

	for result.Next() {

		err := result.Scan(&mDosen.ID, &mDosen.Login, &mDosen.KodeID, &mDosen.IDFINGER, &mDosen.NIDN,
			&mDosen.HomebaseInduk, &mDosen.NIPPNS, &mDosen.NPWP,
			&mDosen.Nama, &mDosen.Struktural, &mDosen.Jenis_Kelamin, &mDosen.TempatLahir,
			&mDosen.TanggalLahir, &mDosen.LevelID, &mDosen.KTP, &mDosen.Telephone, &mDosen.Password, &mDosen.Handphone,
			&mDosen.Email, &mDosen.Alamat, &mDosen.Kota, &mDosen.KodePos, &mDosen.Propinsi, &mDosen.Negara,
			&mDosen.NA, &mDosen.Homebase, &mDosen.ProdiID, &mDosen.Gelar, &mDosen.JenjangID, &mDosen.Keilmuan,
			&mDosen.LulusanPT, &mDosen.AgamaID, &mDosen.KelaminID, &mDosen.GolonganID, &mDosen.FotoKTP, &mDosen.Foto,
			&mDosen.KategoriID, &mDosen.IkatanID, &mDosen.JabatanID, &mDosen.JabatanDiktiID,
			&mDosen.InstitusiInduk, &mDosen.StatusDosenID, &mDosen.StatusKerjaID, &mDosen.TglBekerja, &mDosen.NamaBank,
			&mDosen.NamaAkun, &mDosen.NomerAkun, &mDosen.LoginBuat, &mDosen.TanggalBuat, &mDosen.LoginEdit,
			&mDosen.TanggalEdit, &mDosen.TransReg, &mDosen.TransKhus, &mDosen.SKSKhus, &mDosen.SKSReg)

		if err != nil {
			panic(err.Error())
		}

		m_dosen = append(m_dosen, mDosen)
	}

	json.NewEncoder(w).Encode(m_dosen)
}

func updateMDosen(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newHomebaseInduk := r.FormValue("HomebaseInduk")

		stmt, err := db.Prepare("UPDATE m_dosen SET HomebaseInduk = ? WHERE ID = ?")

		_, err = stmt.Exec(newHomebaseInduk, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Dosen with ID = %s was updated", params["id"])
	}
}

func deleteMDosen(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM m_dosen WHERE ID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Dosen with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var m_dosen []MDosen

	ID := r.FormValue("id")
	HomebaseInduk := r.FormValue("homebaseInduk")

	sql := `SELECT
				ID, 		
				IFNULL(Login,''), 		
				IFNULL(KodeID,'')KodeID, 		
				IFNULL(IDFINGER,'')IDFINGER, 	
				IFNULL(NIDN,'')NIDN, 		
				IFNULL(HomebaseInduk,'')HomebaseInduk, 
				IFNULL(NIPPNS,'')NIPPNS, 		
				IFNULL(NPWP,'')NPWP, 		
				IFNULL(Nama,'')Nama, 		
				IFNULL(Struktural,'')Struktural, 	
				IFNULL(Jenis_Kelamin,'')Jenis_Kelamin, 
				IFNULL(TempatLahir,'')TempatLahir, 
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(LevelID,'')LevelID, 	
				IFNULL(KTP,'')KTP, 		
				IFNULL(Telephone,'')Telephone, 	
				IFNULL(Password,'')Password, 	
				IFNULL(Handphone,'')Handphone,	 
				IFNULL(Email,'')Email, 		
				IFNULL(Alamat,'')Alamat, 		
				IFNULL(Kota,'')Kota, 		
				IFNULL(KodePos,'')KodePos, 	
				IFNULL(Propinsi,'')Propinsi,	
				IFNULL(Negara,'')Negara, 		
				IFNULL(NA,'')NA, 			
				IFNULL(Homebase,'')Homebase, 	
				IFNULL(ProdiID,'')ProdiID, 	
				IFNULL(Gelar,'')Gelar, 		
				IFNULL(JenjangID,'')JenjangID,	
				IFNULL(Keilmuan,'')Keilmuan, 	
				IFNULL(LulusanPT,'')LulusanPT, 	
				IFNULL(AgamaID,'')AgamaID, 	
				IFNULL(KelaminID,'')KelaminID, 	
				IFNULL(GolonganID,'')GolonganID, 	
				IFNULL(FotoKTP,'')FotoKTP, 	
				IFNULL(Foto,'')Foto ,		
				IFNULL(KategoriID,'')KategoriID, 	
				IFNULL(IkatanID,'')IkatanID, 	
				IFNULL(JabatanID,'')JabatanID, 	
				IFNULL(JabatanDiktiID,'')JabatanDiktiID,
				IFNULL(InstitusiInduk,'')InstitusiInduk,
				IFNULL(StatusDosenID,'')StatusDosenID,
				IFNULL(TglBekerja,'')TglBekerja,  
				IFNULL(NamaBank,'')NamaBank,     
				IFNULL(NamaAkun,'')NamaAkun,     
				IFNULL(NomerAkun,'')NomerAkun,    
				IFNULL(LoginBuat,'')LoginBuat,    
				IFNULL(TanggalBuat,'')TanggalBuat,  
				IFNULL(LoginEdit,'')LoginEdit,    
				IFNULL(TanggalEdit,'')TanggalEdit,  
				IFNULL(TransReg,'')TransReg,     
				IFNULL(TransKhus,'')TransKhus,    
				IFNULL(SKSKhus,'')SKSKhus,      
				IFNULL(SKSReg,'')SKSReg  
			FROM m_dosen WHERE ID = ? AND HomebaseInduk = ?`

	result, err := db.Query(sql, ID, HomebaseInduk)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mDosen MDosen

	for result.Next() {

		err := result.Scan(&mDosen.ID, &mDosen.Login, &mDosen.KodeID, &mDosen.IDFINGER, &mDosen.NIDN,
			&mDosen.HomebaseInduk, &mDosen.NIPPNS, &mDosen.NPWP,
			&mDosen.Nama, &mDosen.Struktural, &mDosen.Jenis_Kelamin, &mDosen.TempatLahir,
			&mDosen.TanggalLahir, &mDosen.LevelID, &mDosen.KTP, &mDosen.Telephone, &mDosen.Password, &mDosen.Handphone,
			&mDosen.Email, &mDosen.Alamat, &mDosen.Kota, &mDosen.KodePos, &mDosen.Propinsi, &mDosen.Negara,
			&mDosen.NA, &mDosen.Homebase, &mDosen.ProdiID, &mDosen.Gelar, &mDosen.JenjangID, &mDosen.Keilmuan,
			&mDosen.LulusanPT, &mDosen.AgamaID, &mDosen.KelaminID, &mDosen.GolonganID, &mDosen.FotoKTP, &mDosen.Foto,
			&mDosen.KategoriID, &mDosen.IkatanID, &mDosen.JabatanID, &mDosen.JabatanDiktiID,
			&mDosen.InstitusiInduk, &mDosen.StatusDosenID, &mDosen.StatusKerjaID, &mDosen.TglBekerja, &mDosen.NamaBank,
			&mDosen.NamaAkun, &mDosen.NomerAkun, &mDosen.LoginBuat, &mDosen.TanggalBuat, &mDosen.LoginEdit,
			&mDosen.TanggalEdit, &mDosen.TransReg, &mDosen.TransKhus, &mDosen.SKSKhus, &mDosen.SKSReg)

		if err != nil {
			panic(err.Error())
		}

		m_dosen = append(m_dosen, mDosen)
	}

	json.NewEncoder(w).Encode(m_dosen)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/m_dosen", getMDosen).Methods("GET")
	r.HandleFunc("/m_dosen/{id}", getMDosenByID).Methods("GET")
	r.HandleFunc("/m_dosen", createMDosen).Methods("POST")
	r.HandleFunc("/m_dosen/{id}", updateMDosen).Methods("PUT")
	r.HandleFunc("/m_dosen/{id}", deleteMDosen).Methods("DELETE")

	//new
	r.HandleFunc("/getmdosen", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))
}
