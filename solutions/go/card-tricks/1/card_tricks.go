package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	
	// Jika index valid, kembalikan elemen pada posisi tersebut
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		// Jika di luar jangkauan, tambahkan nilai di akhir slice
		return append(slice, value)
	}
	
	// Jika indeks valid, timpa nilai lama dengan nilai baru
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// Memeriksa validasi indeks agar tidak terjadi panic (out of bounds)
	if index < 0 || index >= len(slice) {
		return slice // Kembalikan slice utuh jika indeks tidak valid
	}
	
	// Menggabungkan slice sebelum indeks dan slice setelah indeks
	return append(slice[:index], slice[index+1:]...)
}
