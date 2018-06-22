<?php
namespace App\Sorting;

final class ArraySorter
{
    public static function bubble(array $array) :array
    {
        for ($i = 0; $i < count($array); $i++) {
            for ($j = $i; $j < (count($array) - 1); $j++) {
                if ($array[$j] > $array[$j + 1]) {
                    $temp = $array[$j + 1];
                    $array[$j + 1] = $array[$j];
                    $array[$j] = $temp;
                }
            }
        }
        return $array;
    }
}
