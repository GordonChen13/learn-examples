<?php

namespace App\Controllers\Auth;

use App\Models\User;

class RegisterController
{
    public function register()
    {
        $model = new User();

        $users = $model->all();

//        var_dump($users);
        header('Content-Type: application/json');
        echo json_encode($users);
    }
}
