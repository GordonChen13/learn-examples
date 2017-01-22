// question:什么是闭包？  simple example 1.
function sayHello(name) {
    var text = 'Hello' + name; //局部变量text
    var say = function () {
        console.log(text);
    }
    return say;
}
var say2 = sayHello('Bob');
say2(); //打印日志： "hello Bob"
     /*闭包的定义：js闭包是绑定了定义时环境中变量的函数。
     say2函数其实就是绑定了在定义sayHello()函数中的变量text的say()函数。*/

// question:闭包是怎么产生的？ simple example 2
function makeAdder(a) {
    function add(b) {
        return a + b;
    }
    return add;
}
var add1 = makeAdder(2);
console.log(add1(3)); //5   思考：实现两个数相加function add(a,b){ return a + b ;}不是更好吗？因为这并不是闭包的应用场景。
    /* js闭包产生的条件：
         1.内嵌函数使用了外部函数的局部变量。
         2.外部变量引用内嵌函数。*/
    /*js闭包的原理：
        JS闭包绑定定义时环境中的变量是通过JS的函数作用域链实现的
        在JS中，函数的作用域遵循一下规则：
            1.函数的作用域等于定义函数时的执行环境。
            2.函数的执行环境等于函数执行时的变量环境(局部变量、参数、...)链上函数的作用域。*/

// question:闭包有哪些应用场景？
    //  1.封装。 simple example 3


    //  2.保存现场。 simple example 4

// question:使用闭包会产生什么问题？ simple example 4


