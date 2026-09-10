<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $daftarMatakuliah = [
            [
                'kode' => 'IF201',
                'nama' => 'Pemrograman Web II',
                'sks' => 3
            ],
            [
                'kode' => 'IF202',
                'nama' => 'Basis Data',
                'sks' => 3
            ],
            [
                'kode' => 'IF203',
                'nama' => 'Sistem Operasi',
                'sks' => 3
            ],
            [
                'kode' => 'IF204',
                'nama' => 'Jaringan Komputer',
                'sks' => 2
            ],
            [
                'kode' => 'IF205',
                'nama' => 'Sistem Kendali',
                'sks' => 2
            ],
        ];

        $kataKunci = $request->query('q', '');

        if ($kataKunci !== '') {
            $daftarMatakuliah = array_filter(
                $daftarMatakuliah,
                function ($matakuliah) use ($kataKunci) {
                    return stripos($matakuliah['kode'], $kataKunci) !== false
                        || stripos($matakuliah['nama'], $kataKunci) !== false;
                }
            );
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'kataKunci' => $kataKunci
        ]);
    }

    public function show(string $kode)
    {
        return view('matakuliah.show', [
            'kode' => $kode
        ]);
    }
}