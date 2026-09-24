<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'kode'     => ['required', 'string', 'max:10', 'unique:matakuliahs,kode'],
            'nama'     => ['required', 'string', 'max:100'],
            'sks'      => ['required', 'integer', 'min:1', 'max:6'],
            'semester' => ['required', 'integer', 'min:1', 'max:14'],
        ];
    }

    public function messages(): array
    {
        return [
            'kode.unique' => 'Kode matakuliah sudah terdaftar',
            'sks.max'     => 'SKS maksimal 6',
        ];
    }
}