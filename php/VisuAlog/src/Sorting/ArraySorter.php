<?php
namespace App\Sorting;

/*
 * 数组排序
 * 从小到大
 */
final class ArraySorter
{
    /**
     * 冒泡排序 O(N^2)
     * @param array $array
     * @return array $arr
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

    /**
     * 选择排序 0(N^2)
     * @param array $arr
     * @return array $arr
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

    /**
     * 插入排序 O(N^2)
     * @param array $arr
     * @return array $arr
     */
    public static function insert(array $arr) :array
    {
        for ($i = 1; $i < count($arr); $i++) {
            for ($j = $i; $j > 0; $j--) {
                if ($arr[$j] < $arr[$j - 1]) {
                    $temp = $arr[$j - 1];
                    $arr[$j - 1] = $arr[$j];
                    $arr[$j] = $temp;
                }
            }
        }
        return $arr;
    }

    /**
     * merge sort O(N logN)
     * @param array $arr
     * @return array
     */
    public static function merge(array $arr) :array
    {
        $count = count($arr);
        if ($count == 1) {
            return $arr;
        }

        $mid = $count/2;
        $left = array_slice($arr, 0, $mid);
        $right = array_slice($arr, $mid);

        $leftSorted = self::merge($left);
        $rightSorted = self::merge($right);

        $n1 = $n2 = 0;

        // merge two sorted array
        for ($i = 0; $i < count($arr); $i++) {
            if ($n2 == count($rightSorted)) {
                $arr[$i] = $leftSorted[$n1];
            } else if ($n1 == count($leftSorted) || $leftSorted[$n1] > $rightSorted[$n2]) {
                $arr[$i] = $rightSorted[$n2];
                $n2++;

            } else {
                $arr[$i] = $leftSorted[$n1];
                $n1++;
            }
        }

        return $arr;

    }

    public static function quick(array $arr) :array
    {
        $left = $right = array();
        if (count($arr) < 2) {
            return $arr;
        }

        $pk = key($arr);
        $p = array_shift($arr);

        for ($i = 0; $i < count($arr); $i++) {
            if ($arr[$i] <= $p) {
                $left[] = $arr[$i];
            } else {
                $right[] = $arr[$i];
            }
        }

        return array_merge(self::quick($left), array($pk => $p), self::quick($right));
    }

}
