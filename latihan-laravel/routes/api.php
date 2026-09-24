<?php

use App\Http\Controllers\Api\MahasiswaController;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\MatakuliahController;

Route::get('/status', function () {
    return response()->json([
        'sukses' => true,
        'pesan' => 'API Pemweb II aktif',
        'waktu' => now()->toIso8601String(),
    ]);
});

Route::get('program-studi/{programStudi}/mahasiswa', [MahasiswaController::class, 'byProgramStudi'])
    ->name('program-studi.mahasiswa');

Route::apiResource('mahasiswa', MahasiswaController::class);
Route::apiResource('matakuliah', MatakuliahController::class);