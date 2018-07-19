<?php
namespace App\Controllers;

use App\Models\Posts;

class PostController extends ApiController
{
    private $model;

    public function __construct()
    {
        $this->model = new Posts();
    }

    public function create()
    {
        $data = $this->validateCreateRequest();

        $result = $this->model->create($data);

        if (!$result) {
            $this->serverHandleError('error.... wait a minute and try again');
        }

        echo json_encode($result);
    }
}