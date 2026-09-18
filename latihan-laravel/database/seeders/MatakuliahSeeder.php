<?php

namespace Database\Seeders;

use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        $daftarMatakuliah = [
            [
                'kode' => 'PWEB2',
                'nama' => 'Pemrograman Web II',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'BD',
                'nama' => 'Basis Data',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'JARKOM',
                'nama' => 'Jaringan Komputer',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'SISKEN',
                'nama' => 'Sistem Kendali',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'IOT',
                'nama' => 'Internet of Things',
                'sks' => 3,
                'semester' => 5,
            ],
            [
                'kode' => 'SIM',
                'nama' => 'Sistem Informasi Manajemen',
                'sks' => 2,
                'semester' => 4,
            ],
        ];

        foreach ($daftarMatakuliah as $matakuliah) {
            Matakuliah::create($matakuliah);
        }
    }
}