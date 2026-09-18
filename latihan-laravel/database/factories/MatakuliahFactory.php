<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Matakuliah>
 */
class MatakuliahFactory extends Factory
{
    public function definition(): array
    {
        return [
            'kode' => fake()->unique()->bothify('MK###'),
            'nama' => fake()->randomElement([
                'Pemrograman Web II',
                'Basis Data',
                'Jaringan Komputer',
                'Sistem Kendali',
                'Internet of Things',
                'Sistem Informasi Manajemen',
            ]),
            'sks' => fake()->randomElement([2, 3, 4]),
            'semester' => fake()->numberBetween(1, 8),
        ];
    }
}