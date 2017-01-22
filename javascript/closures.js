// simple example
function sayHello(name) {
    var text = 'Hello' + name; //局部变量text
    var say = function () {
        console.log(text);
    }
    return say;
}
var say2 = sayHello('Bob');
say2(); //打印日志： "hello Bob"