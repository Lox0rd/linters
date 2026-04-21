fn main() {
    let mut x = 10;
    let y = 20;

    if x == 10 {
        println!("x is ten");
    }

    // лишняя мутация (clippy может ругаться)
    x = x + y;

    let v = vec![1, 2, 3, 4, 5];

    for i in 0..v.len() {
        println!("value: {}", v[i]);
    }

    let _unused_variable = 123;

    let maybe = Some(5);
    match maybe {
        Some(val) => {
            println!("value: {}", val);
        }
        None => {}
    }

    println!("done");
}