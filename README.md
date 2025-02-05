# German Vocabulary Trainer (CLI) 📚🇩🇪

A **Command-Line Application** built in **Go** to help users learn and practice German vocabulary using an interactive menu-driven approach. The app supports adding new words, reviewing learned vocabulary, and testing knowledge through quizzes.

## Features
✅ **Learn New Words** – Add words with translations and examples.  
✅ **Review Mode** – Go through previously learned words.  
✅ **Quiz Mode** – Test your vocabulary knowledge interactively.  
✅ **TDD-Driven Development** – Developed with a test-first approach.  

## Installation

```sh
git clone https://github.com/yourusername/german-vocab-trainer.git
cd german-vocab-trainer
go build -o vocab-trainer
./vocab-trainer
```

## Usage

Run the program and choose an option:

```
Let me know what do you want to do now:
1. Learn a new word
2. Review learnt words
3. Quiz
4. Exit
Enter your choice: 
```

## Running Tests

```sh
go test ./...
```

## Roadmap 🛤️
🔹 Integrate AI-generated example sentences  
🔹 Implement Spaced Repetition for better retention  
🔹 Expand to GUI/iOS app  

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
MIT License
