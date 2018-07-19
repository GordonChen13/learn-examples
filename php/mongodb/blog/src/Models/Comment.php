<?php
namespace App\Models;

use MongoDB\Client;

class Comment
{
    private $post;

    public function __construct(string $post_id)
    {
        $this->post = (new Client())->blog->posts->find(['id' => $post_id]);
    }

    public function create(array $comment)
    {

    }

    public function update(array $comment)
    {

    }

    public function find(array $condition)
    {

    }

    public function delete(array $condition)
    {

    }
}