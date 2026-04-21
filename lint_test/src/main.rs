fn main() {
    let numbers = vec![1, 2, 3, 4, 5, 6, 7, 8, 9];
    let mut sum = 0;

    for i in 0..numbers.len() {
        sum += numbers[i];
    }

    println!("sum = {}", sum);

    let avg = sum / numbers.len() as i32;
    println!("avg = {}", avg);

    let name = String::from("Rust");
    greet(name);

    let result = divide(10, 2);
    println!("result = {}", result.unwrap());

    let maybe_value = get_value(true);
    match maybe_value {
        Some(v) => println!("Got value: {}", v),
        None => println!("No value"),
    }

    let data = vec![10, 20, 30];
    print_data(&data);

    let mut counter = Counter::new();

    counter.increment();
    counter.increment();
    counter.increment();

    println!("counter = {}", counter.value);

    let squared: Vec<i32> = data.iter().map(|x| x * x).collect();
    println!("squared = {:?}", squared);

    let filtered: Vec<i32> = numbers.into_iter().filter(|x| x % 2 == 0).collect();
    println!("filtered = {:?}", filtered);

    let message = format_message("hello", 42);
    println!("{}", message);

    let config = Config {
        debug: true,
        max_connections: 10,
    };

    println!("config: debug={}, max={}", config.debug, config.max_connections);
}

fn greet(name: String) {
    println!("Hello, {}", name);
}

fn divide(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        return Err(String::from("division by zero"));
    }
    Ok(a / b)
}

fn get_value(flag: bool) -> Option<i32> {
    if flag {
        Some(100)
    } else {
        None
    }
}

fn print_data(data: &Vec<i32>) {
    for i in 0..data.len() {
        println!("data[{}] = {}", i, data[i]);
    }
}

fn format_message(msg: &str, code: i32) -> String {
    format!("message: {} with code {}", msg, code)
}

struct Counter {
    value: i32,
}

impl Counter {
    fn new() -> Self {
        Counter { value: 0 }
    }

    fn increment(&mut self) {
        self.value += 1;
    }
}

struct Config {
    debug: bool,
    max_connections: i32,
}