<?php
namespace App\Controllers;
header('Content-Type: application/json');

class ApiController
{
    public function __construct()
    {

    }

    public function badRequestError(string $message)
    {
        http_response_code(422);

        echo json_encode([
            'status' => 'bad request error',
            'message' => $message
        ]);
        die();
    }
}