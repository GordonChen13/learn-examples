<?php

use NoahBuscher\Macaw\Macaw;
use App\Controllers;

Macaw::get('api/register', 'Auth\RegisterController@register');
Macaw::get('api/hello', function () {
    echo 'hello world! blog';
});
Macaw::get('test', function () {
    echo "php test";
});

Macaw::dispatch();