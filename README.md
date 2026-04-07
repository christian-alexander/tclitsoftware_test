# Stability Team Technical Test : Christian Alexander

## Issues Found 
1. HTTP status code yang salah ketika akses GetTasks dengan ID yang tidak ada
2. tidak ada validasi inputan di create task

## How to fix 
1. Ketika ada hit dengan ID yang tidak ditemukan, harusnya return 404 (not found) bukan 200 (OK)
2. diberi validasi dengan menggunakan package validator, untuk title harus diisi (required)

## improvement
1. 