<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $data = [
            'id'         => $this->id,
            'nim'        => $this->nim,
            'nama'       => $this->nama,
            'email'      => $this->email,
            'angkatan'   => (int) $this->angkatan,
            'ipk'        => (float) $this->ipk,
            'aktif'      => (bool) $this->aktif,
        ];

        if ($this->relationLoaded('programStudi')) {
            $data['program_studi'] = $this->whenLoaded('programStudi', function () {
                return [
                    'id'      => $this->programStudi->id,
                    'kode'    => $this->programStudi->kode,
                    'nama'    => $this->programStudi->nama,
                    'jenjang' => $this->programStudi->jenjang,
                ];
            });
        }

        $fields = $request->attributes->get('fields');

        if (!empty($fields) && is_array($fields)) {
            $data = array_intersect_key($data, array_flip($fields));
        }

        return $data;
    }
}