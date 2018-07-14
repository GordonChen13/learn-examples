<?php
namespace App\Models;

use MongoDB\Client;
use App\Controllers\ApiController;
use Carbon\Carbon;

class Posts extends ApiController
{
    private $collection;

    public function __construct()
    {
        $this->collection = (new Client())->blog->posts;
    }

    public function create(array $post)
    {
        $this->collection->insertOne(
            [
                'title' => $post['title'],
                ''
            ]
        );
    }
}