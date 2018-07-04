<?php

namespace Store\Toys;

use Phalcon\Mvc\Model;
use Phalcon\Mvc\Model\Message;
use Phalcon\Mvc\Model\Validator\Uniqueness;
use Phalcon\Mvc\Model\Validator\Inclusionin;

class Robots extends Model
{
    public function validate(\Phalcon\ValidationInterface $validator)
    {
        $this->validate(
            new Inclusionin(
                [
                    'field' => 'type',
                    'domain' => [
                        'droid',
                        'mechanical',
                        'virtual'
                    ],
                ]
            )
        );

        $this->validate(
            new Uniqueness(
                [
                    'field' => 'name',
                    'message' => 'the robot name must be unique',
                ]
            )
        );

        if ($this->year < 0) {
            $this->appendMessage(
                new Message('The year cannot be less than zero')
            );
        }

        if ($this->validationHasFailed() === true) {
            return false;
        }
    }
}