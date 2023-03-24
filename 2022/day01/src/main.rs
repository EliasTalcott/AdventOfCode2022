use std::fs;

fn main() {
    // Part A
    let elves = fs::read_to_string("calories.txt").expect("unable to read file");
    let elves = elves.split("\n\n");
    let calorie_counts= elves.map(|elf| elf.lines().map(|calories| calories.parse::<u32>().unwrap()).sum::<u32>());
    let max_calories = calorie_counts.clone().max().unwrap();
    println!("part a: {}", max_calories);

    // Part B
    let mut calorie_counts: Vec<u32> = calorie_counts.collect();
    calorie_counts.sort_unstable();
    calorie_counts.reverse();
    let max_three_calories = calorie_counts.into_iter().take(3).sum::<u32>();
    println!("part b: {}", max_three_calories);
}
