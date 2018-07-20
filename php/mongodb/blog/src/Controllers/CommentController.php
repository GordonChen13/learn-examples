<?php
namespace App\Controllers;

use App\Models\Comment;

class CommentController extends ApiController
{
    private $model;

    public function __construct()
    {
        $post_id = $_GET['post_id'];
        $this->model = new Comment($post_id);
    }

    public function create()
    {
        $data = $this->validateCreateRequest();

        $this->model->create($data);
    }

    public function update()
    {

    }

    public function index()
    {

    }

    public function delete()
    {

    }
}