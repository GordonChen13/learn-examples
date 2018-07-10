<?php

namespace App\Controllers\Auth;

use App\Models\User;
use App\Controllers\ApiController;

class RegisterController extends ApiController
{
    public function register()
    {
        $model = new User();

        $this->validate();

        $data = array(
            'name' => $_POST['name'],
            'phone' => $_POST['phone'],
            'password' => $_POST['password']
        );

        $result = $model->create($data);

        echo json_encode($result->getInsertedId());
    }

    private function validate()
    {
        $name = $_POST['name'];
        $phone = $_POST['phone'];
        $password = $_POST['password'];
        $password_confirmed = $_POST['password_confirmed'];

        if (is_null($name)) {
            $this->badRequestError('name could not be empty');
        }

        if (is_null($phone)) {
            $this->badRequestError('phone could not be empty');
        }

        if (is_null($password)) {
            $this->badRequestError('password could not be empty');
        }

        if (is_null($password_confirmed)) {
            $this->badRequestError('password_confirmed could not be empty');
        }

        if ($password_confirmed !== $password) {
            $this->badRequestError('password was not same');
        }
    }
}
