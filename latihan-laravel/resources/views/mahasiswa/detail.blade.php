@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
    <h1 class="h3 mb-4">Detail Mahasiswa</h1>

    <div class="card mb-4">
        <div class="card-header">
            Informasi Mahasiswa
        </div>

        <div class="card-body">
            <p>
                <strong>NIM:</strong>
                {{ $mahasiswa->nim }}
            </p>

            <p>
                <strong>Nama:</strong>
                {{ $mahasiswa->nama }}
            </p>

            <p>
                <strong>Program Studi:</strong>
                {{ $mahasiswa->programStudi->nama }}
            </p>

            <p>
                <strong>Angkatan:</strong>
                {{ $mahasiswa->angkatan }}
            </p>

            <p class="mb-0">
                <strong>IPK:</strong>
                {{ $mahasiswa->ipk }}
            </p>
        </div>
    </div>

    <div class="card">
        <div class="card-header">
            Daftar Mata Kuliah yang Diambil
        </div>

        <div class="card-body">
            @if ($mahasiswa->matakuliah->isEmpty())
                <div class="alert alert-warning">
                    Mahasiswa ini belum mengambil mata kuliah.
                </div>
            @else
                <table class="table table-bordered">
                    <thead>
                        <tr>
                            <th>No.</th>
                            <th>Kode</th>
                            <th>Nama Mata Kuliah</th>
                            <th>SKS</th>
                            <th>Semester</th>
                            <th>Nilai</th>
                        </tr>
                    </thead>

                    <tbody>
                        @foreach ($mahasiswa->matakuliah as $matakuliah)
                            <tr>
                                <td>{{ $loop->iteration }}</td>
                                <td>{{ $matakuliah->kode }}</td>
                                <td>{{ $matakuliah->nama }}</td>
                                <td>{{ $matakuliah->sks }}</td>
                                <td>{{ $matakuliah->semester }}</td>
                                <td>{{ $matakuliah->pivot->nilai }}</td>
                            </tr>
                        @endforeach
                    </tbody>
                </table>
            @endif
        </div>
    </div>

    <a href="{{ route('mahasiswa.data') }}"
       class="btn btn-secondary mt-3">
        Kembali
    </a>
@endsection