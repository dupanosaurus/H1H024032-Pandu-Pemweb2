@extends('layouts.app')

@section('judul', 'Detail Mata Kuliah')

@section('konten')

<h1 class="h3 mb-4">Detail Mata Kuliah</h1>

<div class="card">
    <div class="card-body">

        <p class="mb-0">
            Kode mata kuliah yang diminta:
            <strong>{{ $kode }}</strong>
        </p>

    </div>
</div>

<a
    href="{{ route('matakuliah.index') }}"
    class="btn btn-secondary mt-3"
>
    Kembali
</a>

@endsection