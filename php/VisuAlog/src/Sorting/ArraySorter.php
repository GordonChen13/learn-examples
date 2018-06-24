<?php
namespace App\Sorting;

/*
 * 数组排序
 * 从小到大
 */
final class ArraySorter
{
    /*
     * 冒泡排序
     */
    public static function bubble(array $array) :array
    {
        for ($i = 0; $i < count($array); $i++) {
            for ($j = 0; $j < (count($array) - $i - 1); $j++) {
                if ($array[$j] > $array[$j + 1]) {
                    $temp = $array[$j + 1];
                    $array[$j + 1] = $array[$j];
                    $array[$j] = $temp;
                }
            }
        }
        return $array;
    }

    /*
     * 选择排序
     */
    public static function select(array $arr) :array
    {
        for ($i = 0; $i < count($arr); $i++) {
            for($j = $i; $j < count($arr); $j++) {
                if($arr[$i] > $arr[$j]) {
                    $temp = $arr[$i];
                    $arr[$i] = $arr[$j];
                    $arr[$j] = $temp;
                }
            }
        }

        return $arr;
    }

    /*
     * 插入排序
     */
    public static function insert(array $arr) :array
    {
        if (count($arr) == 0) {
            return array();
        }

        $sorted = array($arr[0]);

        for ($i = 1; $i < count($arr); $i++) {
            for ($j = 0; $j < count($sorted); $j++) {
                if($sorted[$j] > $arr[$i]) {
                    array_splice($sorted, $j, 0, $arr[$i]);
                }
            }
        }
        return $sorted;
    }
}
