use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn fuel_calc(weight:u32) -> u32 {
    let mut total = (weight/3)-2;
    let mut current = total;
    while (current as f32 /3.0)-2.0 > 0.0 {
        //println!("{}", current);
        current = (current/3)-2;
        total += current;
    }
    return total as u32;
}

fn main() {
    //let data = &[1969, 100756]; //TESTING
    let mut fuel_total = 0;
    if let Ok(lines) = read_lines("../day_1_input.txt") {
        for line in lines {
            if let Ok(module) = line {
                let fuel = fuel_calc(module.parse::<u32>().unwrap());
                fuel_total += fuel;
            }
        }
    }
    /* TESTING
    for i in data {
        println!("number {}", i);
        let fuel = fuel_calc(*i);
        println!("Fuel weight: {}", fuel);
        fuel_total += fuel;
    }
    */
    println!("\nTotal: {}", fuel_total);
}
