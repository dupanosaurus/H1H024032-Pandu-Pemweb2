<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use Illuminate\Database\Seeder;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    public function run(): void
    {
        $this->call([
            ProgramStudiSeeder::class,
            MatakuliahSeeder::class,
        ]);

        Mahasiswa::factory()->count(30)->create();
    }
}