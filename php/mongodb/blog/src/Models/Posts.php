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
                'content' => $post['content'],
                'created_at' => Carbon::now(),
                'update_at' => Carbon::now(),
                'owner' => $post['owner'],
                'comments' => $post['comments'],
                'likes_count' => $post['likes_count'],
                'views_count' => $post['views_count'],
            ]
        );
    }

    public function update(aray $post , array $condition)
    {
        $this->collection->updateOne($condition, $post);
    }

    public function delete(array $condition)
    {
        $this->collection->deleteOne($condition);
    }
}