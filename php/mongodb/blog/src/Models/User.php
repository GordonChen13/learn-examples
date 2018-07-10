<?php
namespace App\Models;

use MongoDB\Client;

class User
{
    public $collection;

    public function __construct()
    {
        $this->collection = (new Client())->blog->users;
    }

    /**
     *  新增用户
     * @param array $user
     * @return \MongoDB\InsertOneResult
     */
    public function create(array $user)
    {
         $result = $this->collection->insertOne([
             'name' => $user['name'],
             'phone' => $user['phone'],
             'password' => md5($user['password'])
         ]);

         return $result;
    }

    /**
     *  获取所有记录
     * @return array
     */
    public function all()
    {
        $result = $this->collection->find([

        ]);

        return $result->toArray();
    }
}