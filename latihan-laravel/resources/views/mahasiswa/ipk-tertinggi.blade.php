@extends('layouts.app')

@section('judul', '10 Mahasiswa dengan IPK Tertinggi')

@section('konten')
    <h1 class="h3 mb-4">
        10 Mahasiswa dengan IPK Tertinggi
    </h1>

    <div class="alert alert-info">
        Daftar mahasiswa Program Studi Teknik Komputer
        berdasarkan urutan IPK tertinggi.
    </div>

    <table class="table table-bordered table-striped bg-white">
        <thead>
            <tr>
                <th>Ranking</th>
                <th>NIM</th>
                <th>Nama</th>
                <th>Program Studi</th>
                <th>Angkatan</th>
                <th>IPK</th>
            </tr>
        </thead>

        <tbody>
            @forelse ($daftarMahasiswa as $mahasiswa)
                <tr>
                    <td>{{ $loop->iteration }}</td>
                    <td>{{ $mahasiswa->nim }}</td>
                    <td>{{ $mahasiswa->nama }}</td>
                    <td>{{ $mahasiswa->programStudi->nama }}</td>
                    <td>{{ $mahasiswa->angkatan }}</td>
                    <td>
                        <strong>{{ $mahasiswa->ipk }}</strong>
                    </td>
                </tr>
            @empty
                <tr>
                    <td colspan="6">
                        Data mahasiswa belum tersedia.
                    </td>
                </tr>
            @endforelse
        </tbody>
    </table>

    <a href="{{ route('mahasiswa.data') }}"
       class="btn btn-secondary">
        Kembali
    </a>
@endsection