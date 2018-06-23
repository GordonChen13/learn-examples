<?php
namespace Tests\Sorting;

use PHPUnit\Framework\TestCase;
use App\Sorting\ArraySorter;

class ArraySorterTest extends TestCase
{
    /*
     * 生产随机数组
     */
    public function randomArray() :array
    {
        for($i = 0; $i < rand(1, 20); $i++) {
            $data[] = rand(-100, 100);
        }

        return $data;
    }

    /*
     * 判断两个数组是否相等
     */
    public function isArrayEqual(array $a, array $b)
    {
        if (count(array_diff_assoc($a, $b))) {
            return false;
        }

        foreach($a as $k => $v) {
            if ($v !== $b[$k]) {
                return false;
            }
        }

        return true;
    }

    /**
     * 数组相等测试
     */
    public function testArrayEqual()
    {
        $arr = $this->randomArray();

        $bubbleSorted = ArraySorter::bubble($arr);

        asort($arr);
        $phpSorted = array_values($arr);

        $this->assertTrue($this->isArrayEqual($phpSorted, $bubbleSorted));
    }
}
