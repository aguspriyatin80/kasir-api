package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Semua yang tergolong makanan berat dan ringan"},
	{ID: 2, Name: "Minuman", Description: "Semua yang tergolong minuman"},
}

var produk = []Produk{
	{ID: 1, Nama: "Fanta", Harga: 5000, Stok: 100},
	{ID: 2, Nama: "Sprite", Harga: 5000, Stok: 200},
	{ID: 3, Nama: "Coca-cola", Harga: 5000, Stok: 100}, //meskipun data terakhir, harus tetap diakhiri tanda koma (,)
}

// DELETE localhost:8080/api/kategori/{id}
func deleteCategoryById(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	// ubah id ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// loop kategori cari ID, dapat index yang mau dihapus
	for i, c := range categories {
		if c.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}
	http.Error(w, "Invalid Request", http.StatusNotFound)
}

// DELETE localhost:8080/api/produk/{id}
func deleteProductById(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ubah id ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// loop produk cari ID, dapat index yang mau dihapus
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}
	http.Error(w, "Invalid Request", http.StatusNotFound)
}

// PUT localhost:8080/api/kategori/{id}
func updateCategoryById(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	// ganti id ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateKategori Category
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	for i := range categories {
		if categories[i].ID == id {
			updateKategori.ID = id
			categories[i] = updateKategori
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}
	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// PUT localhost:8080/api/produk/{id}
func updateProductById(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

// GET localhost:8080/api/kategori/{id}
func getCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest) //404
		return
	}
	for _, c := range categories {
		if c.ID == id {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// GET localhost:8080/api/produk/{id}
func getProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk ID", http.StatusBadRequest) //404
		return
	}
	for _, p := range produk {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func main() {
	//GET localhost:8080/api/kategori/{id} => API untuk menampilkan data kategori berdasarkan ID
	//PUT localhost:8080/api/kategori/{id} => API untuk mengubah  data kategori berdasarkan ID
	//DELETE localhost:8080/api/kategori/{id} => API untuk menghapus data kategori berdasarkan ID
	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w, r)
		} else if r.Method == "PUT" {
			updateCategoryById(w, r)
		} else if r.Method == "DELETE" {
			deleteCategoryById(w, r)
		}
	})

	//GET localhost:8080/api/produk/{id} => API untuk menampilkan data produk  berdasarkan ID
	//PUT localhost:8080/api/produk/{id} => API untuk mengubah  data produk berdasarkan ID
	//DELETE localhost:8080/api/produk/{id} => API untuk menghapus data produk berdasarkan ID
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductById(w, r)
		} else if r.Method == "PUT" {
			updateProductById(w, r)
		} else if r.Method == "DELETE" {
			deleteProductById(w, r)
		}
	})

	//GET localhost:8080/api/kategori => API untuk menampilkan semua data kategori
	//POST localhost:8080/api/kategori => API untuk membuat kategori baru
	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var kategoriBaru Category
			err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest) //404
			}
			kategoriBaru.ID = len(categories) + 1
			categories = append(categories, kategoriBaru)
			w.WriteHeader(http.StatusCreated) //201
			json.NewEncoder(w).Encode(kategoriBaru)
		}
	})

	//GET localhost:8080/api/produk => API untuk menampilkan semua data produk
	//POST localhost:8080/api/produk => API untuk membuat produk baru
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest) //404
			}
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)
			w.WriteHeader(http.StatusCreated) //201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8080/health => API untuk mengecek apakah server sedang berjalan
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Server Running",
		})
	})
	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
