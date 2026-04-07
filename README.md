# Stability Team Technical Test : Christian Alexander

## Issues Found 
1. HTTP status code yang salah ketika akses GetTasks dengan ID yang tidak ada
2. tidak ada validasi inputan di create task
3. fix return dari task_store untuk get task by id, sebelumnya pakai clone an dari tasks hasil iterate

## How to fix 
1. Ketika ada hit dengan ID yang tidak ditemukan, harusnya return 404 (not found) bukan 200 (OK)
2. diberi validasi dengan menggunakan package validator, untuk title harus diisi (required)
3. returnan untuk tasks, jangan pakai copy an dari tasks, return pointer tasknya. return copyan dari tasks hanya cocok untuk data return yang read only. kalau sewaktu waktu porgram ini berkembang lebih kompleks dan get task dipakai di module lain. dan ada yang mutate hasil GetTaskById, maka instance Tasks nya tidak berubah karena yang direturn copyan

## improvement
1. 