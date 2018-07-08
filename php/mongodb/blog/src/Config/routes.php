<?php

use NoahBuscher\Macaw\Macaw;
use App\Controllers;

Macaw::get('api/register', 'Auth\RegisterController@register');